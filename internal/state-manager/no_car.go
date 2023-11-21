package state_manager

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
)

func (v *StateManager) noCarActions() {
	v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	v.controllerImpl.CloseBarrier()
	v.controllerImpl.DisplayShowText("Waiting for new user")
}

func (v *StateManager) onNoCarDetection() {
	v.noCarActions()

	resetState(v.state)
}
