package state_manager

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
)

func (s *StateManager) noCarActions() {
	s.controllerImpl.TurnLEDOff(mqtt_client.GreenLEDPubTop)
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.CloseBarrier()
	s.controllerImpl.DisplayShowText("Waiting for new user")
}

func (s *StateManager) onNoCarDetection() {
	s.noCarActions()
	s.state.reset()
}
