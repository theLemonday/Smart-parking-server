package viewmodel

import (
	"fmt"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"

	"github.com/rs/zerolog/log"
)

func (v *Viewmodel) generateID4Identify() {
	v._currentUID = util.GenerateNewNanoID(10)
	if !v.isGoIn {
		v._currentUID = fmt.Sprintf("tt%s", v._currentUID)
	}
}

func (v *Viewmodel) carGoInOrOutActions() {
	log.Info().Msg("Publish car in or out actions")

	// v.controllerImpl.TurnLEDOn(topic.RedLEDPubTop)
	// v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	// v.controllerImpl.DisplayShowQRCode(v._currentUID)
}

func (v *Viewmodel) onCarGoIn() {
	v.generateID4Identify()

	v.carGoInOrOutActions()

	v._newUserIdenitfyStatus = _WaittingToBeIdentified

	log.Info().Msg("Waiting the user to be identified")

}
