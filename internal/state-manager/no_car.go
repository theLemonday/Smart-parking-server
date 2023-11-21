package viewmodel

import (
	"github.com/rs/zerolog/log"
)

func (v *StateManager) noCarActions() {
	log.Info().Msg("Publish no car actions")

	// v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	// v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	// v.controllerImpl.CloseBarrier()
	// v.controllerImpl.DisplayShowText("Waiting for new user")
}

func (v *StateManager) onNoCarDetection() {
	v.noCarActions()

	v.newUserIdenitfyStatus = unknown
	v.carGoIn = false
	v.carGoOut = false

	log.Info().Msg("viewmodel reset")
}
