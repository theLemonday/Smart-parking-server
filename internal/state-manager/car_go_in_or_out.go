package viewmodel

import (
	"fmt"

	"github.com/thelemonday/smart-parking-iot-server/pkg/util"

	"github.com/rs/zerolog/log"
)

func (v *StateManager) generateID4Identify() {
	v.identificationID = util.GenerateNewNanoID(10)
	if !v.isGoIn {
		v.identificationID = fmt.Sprintf("tt%s", v.identificationID)
	}
}

func (v *StateManager) carGoInOrOutActions() {
	log.Info().Msg("Publish car in or out actions")

	// v.controllerImpl.TurnLEDOn(topic.RedLEDPubTop)
	// v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	// v.controllerImpl.DisplayShowQRCode(v._currentUID)
}

func (v *StateManager) onCarGoIn() {
	v.generateID4Identify()

	v.carGoInOrOutActions()

	v.newUserIdenitfyStatus = waittingToBeIdentified

	log.Info().Msg("Waiting the user to be identified")

}
