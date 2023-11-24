package state_manager

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"strings"

	"github.com/rs/zerolog/log"
)

func (s *StateManager) carGoInOrOutActions() {
	log.Info().Msg("Publish car in or out actions")

	s.controllerImpl.TurnLEDOn(mqtt_client.RedLEDPubTop)
	s.controllerImpl.TurnLEDOff(mqtt_client.GreenLEDPubTop)
	s.controllerImpl.DisplayShowWelcome()
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
		s.carGoInOrOutActions()

		s.state.newUserIdentifyStatus = waitingToBeIdentified

		log.Info().Msg("Waiting the user to be identified")

		return
	}
}

func (s *StateManager) OnCarGoInSlotDetected(sensorTopic string, detected bool) {
	slotId := strings.TrimPrefix(sensorTopic, mqtt_client.IRSlotPrefix)
	s.state.slotsStatus[slotId] = detected
	s.transfer2websocketService.OnSlotStatusChanged(slotId, detected)
}

func (s *StateManager) onUserIdentified() {
	s.controllerImpl.OpenBarrier()
	s.controllerImpl.DisplayShowWelcome()
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.TurnLEDOn(mqtt_client.GreenLEDPubTop)
}

func (s *StateManager) OnRFIDTagRead(uid string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	log.Info().Msgf("New user identified by rfid: %s", uid)

	s.state.newUserIdentifyStatus = identified

	if s.state.isGoIn {
		user := s.carParkStatusDb.NewUserIdentifiedByRFID(uid)
		if user == nil {
			log.Info().Msg("Invalid RFID tags, do nothing")
			return
		}
		s.transfer2websocketService.OnNewUserEnter(user)
		s.onUserIdentified()
		return
	}

	bill := s.carParkStatusDb.CalculateCost(uid)
	s.transfer2websocketService.OnUserGoOutIdentified(bill)
}

func (s *StateManager) OnUserIdentifiedByRFIDDonePayment(uid string) {
	s.carParkStatusDb.OnUserLeave(uid)
	s.transfer2websocketService.OnUserLeave(uid)
	s.onUserIdentified()
}
