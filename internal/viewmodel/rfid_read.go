package viewmodel

import (
	"github.com/rs/zerolog/log"
)

func (v *Viewmodel) onUserIdentified() {
	log.Info().Msg("publish action after user identified")
	// v.controllerImpl.OpenBarrier()
	// v.controllerImpl.DisplayShowText("Welcome to our smart parking system")
	// v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	// v.controllerImpl.TurnLEDOn(topic.GreenLEDPubTop)
}

func (v *Viewmodel) onRFIDRead(uid string) {
	if v._newUserIdenitfyStatus == _WaittingToBeIdentified {
		if !v.database.IsRFIDTagValid(uid) {
			log.Info().Msg("Invalid RFID tags, do nothing")
			return
		}

		log.Info().Msgf("New user identified by rfid: %s", uid)

		v.onUserIdentified()

		v._newUserIdenitfyStatus = _Identified
		v.database.InsertNewCurrentUser(uid, string(_IdentifiedByRFID))
	}
}
