package state_manager

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

func generateID4Identify(isGoIn bool) string {
	uid := util.GenerateNewNanoID(10)
	if !isGoIn {
		uid = fmt.Sprintf("tt%s", uid)
	}

	return uid
}

func (s *StateManager) carGoInOrOutActions() {
	log.Info().Msg("Publish car in or out actions")

	s.controllerImpl.TurnLEDOn(mqtt_client.RedLEDPubTop)
	s.controllerImpl.TurnLEDOff(mqtt_client.GreenLEDPubTop)
	s.controllerImpl.DisplayShowQRCode(s.state.identificationID)
}

func (s *StateManager) OnCarGoInDetected(detected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.state.carGoIn = detected
	if !s.state.carGoIn {
		if !s.state.carGoOut {
			s.onNoCarDetection()
			return
		}
		return
	}

	log.Info().Msg("Car goes in detected")
	s.state.isGoIn = true

	if s.state.newUserIdentifyStatus == unknown {
		s.state.identificationID = generateID4Identify(s.state.isGoIn)

		s.carGoInOrOutActions()

		s.state.newUserIdentifyStatus = waitingToBeIdentified

		log.Info().Msg("Waiting the user to be identified")

		return
	}
}

func (s *StateManager) OnCarGoOutDetected(detected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.state.carGoIn = detected
	if !s.state.carGoIn {
		if !s.state.carGoOut {
			s.onNoCarDetection()
			return
		}
		return
	}

	log.Info().Msg("Car goes out detected")
	s.state.isGoIn = false

	if s.state.newUserIdentifyStatus == unknown {
		s.state.identificationID = generateID4Identify(s.state.isGoIn)

		s.carGoInOrOutActions()

		s.state.newUserIdentifyStatus = waitingToBeIdentified

		log.Info().Msg("Waiting the user to be identified")

		return
	}
}

func (s *StateManager) OnCarGoInSlotDetected(sensorTopic string, detected bool) {
	slotId := strings.TrimLeft(sensorTopic, mqtt_client.IRSlotPrefix)
	s.state.slotsStatus[slotId] = detected
	s.transfer2websocketService.OnSlotStatusChanged(slotId, detected)
}

func (s *StateManager) onUserIdentified() {
	s.controllerImpl.OpenBarrier()
	s.controllerImpl.DisplayShowText("Welcome to our smart parking system")
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.TurnLEDOn(mqtt_client.GreenLEDPubTop)
}

func (s *StateManager) OnRFIDTagRead(uid string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	if !s.carParkStatusDb.IsRFIDTagValid(uid) {
		log.Info().Msg("Invalid RFID tags, do nothing")
		return
	}

	log.Info().Msgf("New user identified by rfid: %s", uid)

	s.state.newUserIdentifyStatus = identified
	user := s.carParkStatusDb.NewUserIdentifiedByRFID(uid)

	if s.state.isGoIn {
		s.transfer2websocketService.OnNewUserEnter(user)
		s.onUserIdentified()
	} else {
		bill := s.carParkStatusDb.CalculateCost(uid)
		s.transfer2websocketService.OnUserGoOutIdentified(bill)
	}
}

func (s *StateManager) OnQRCodeScanned(QRCode, username string) {
	if s.state.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	if s.state.identificationID != QRCode {
		log.Info().Msg("QRCode does not match")
		return
	}

	log.Info().Msgf("New user identified by qr code: %s", QRCode)

	s.state.newUserIdentifyStatus = identified
	user := s.carParkStatusDb.NewUserIdentifiedByQRCode(username)

	if s.state.isGoIn {
		s.transfer2websocketService.OnNewUserEnter(user)
	}

	s.onUserIdentified()
}

func (s *StateManager) OnUserIdentifiedByRFIDDonePayment(uid string) {
	s.carParkStatusDb.DeleteUser(uid)
	s.onUserIdentified()
}
