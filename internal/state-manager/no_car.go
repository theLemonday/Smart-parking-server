package state_manager

import "github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"

func (s *StateManager) onNoCarDetection() {
	s.controllerImpl.TurnLEDOff(mqtt_client.GreenLEDPubTop)
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.CloseBarrier()

	s.state.reset()
}
