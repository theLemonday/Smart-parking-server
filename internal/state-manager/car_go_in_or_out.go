package state_manager

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

func (v *StateManager) carGoInHandler() {
	if !v.carGoIn {
		if !v.carGoOut {
			v.onNoCarDetection()
			return
		}
		return
	}

	log.Info().Msg("Car goes in detected")
	v.isGoIn = true

	if v.newUserIdentifyStatus == unknown {
		v.generateID4Identify()

		v.carGoInOrOutActions()

		v.newUserIdentifyStatus = waitingToBeIdentified

		log.Info().Msg("Waiting the user to be identified")

		return
	}
}
