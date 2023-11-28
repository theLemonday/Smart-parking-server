package state_manager

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"strings"
)

func (s *StateManager) waitingToBeIdentified() {
	log.Info().Msg("Waiting the user to be identified")
	s.state.newUserIdentifyStatus = waitingToBeIdentified

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
		}
		return
	}

	log.Info().Msg("Car goes in detected")
	s.state.isGoIn = true

	if s.state.newUserIdentifyStatus == unknown {
		s.waitingToBeIdentified()
	}
}

func (s *StateManager) OnCarGoOutDetected(detected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.state.carGoOut = detected
	if !s.state.carGoOut {
		if !s.state.carGoIn {
			s.onNoCarDetection()
		}
		return
	}

	log.Info().Msg("Car goes out detected")
	s.state.isGoIn = false

	if s.state.newUserIdentifyStatus == unknown {
		s.waitingToBeIdentified()
	}
}

func (s *StateManager) onUserIdentified() {
	s.controllerImpl.OpenBarrier()
	if s.state.isGoIn {
		s.controllerImpl.DisplayShowWelcome()
	} else {
		s.controllerImpl.DisplayShowSeeYouAgain()
	}
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.TurnLEDOn(mqtt_client.GreenLEDPubTop)
}

func (s *StateManager) OnRFIDTagRead(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	log.Info().Msgf("New user identified by rfid: %s", id)

	s.state.newUserIdentifyStatus = identified

	if s.state.isGoIn {
		user := s.userUseCase.NewUserEnter(id)
		s.transfer2websocketService.OnNewUserEnter(user)
		s.onUserIdentified()
		return
	}

	bill := s.userUseCase.CalculateCarParkBill(id)
	s.transfer2websocketService.OnUserGoOutIdentified(bill)
}

func (s *StateManager) OnUserIdentifiedByRFIDDonePayment(id string) {
	s.userUseCase.UserLeave(id)
	s.transfer2websocketService.OnUserLeave(id)
	s.onUserIdentified()
}

func (s *StateManager) OnCarGoInSlotDetected(sensorTopic string, detected bool) {
	slotId := strings.TrimPrefix(sensorTopic, mqtt_client.IRSlotPrefix)
	s.state.slotsStatus[slotId] = detected
	s.transfer2websocketService.OnSlotStatusChanged(slotId, detected)
}
