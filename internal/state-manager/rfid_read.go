package viewmodel

import (
	"github.com/rs/zerolog/log"
)

func (v *StateManager) onUserIdentified() {
	log.Info().Msg("publish action after user identified")
	// v.controllerImpl.OpenBarrier()
	// v.controllerImpl.DisplayShowText("Welcome to our smart parking system")
	// v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	// v.controllerImpl.TurnLEDOn(topic.GreenLEDPubTop)
}

func (v *StateManager) newUserIdentifiedHandler(username string) {
	if v.newUserIdenitfyStatus != waittingToBeIdentified {
		return
	}

	if !v.database.IsRFIDTagValid(username) {
		log.Info().Msg("Invalid RFID tags, do nothing")
		return
	}

	log.Info().Msgf("New user identified by rfid: %s", username)

	v.onUserIdentified()

	v.newUserIdenitfyStatus = identified
	v.database.NewUserIdentifiedByRFID(username)

}

// If qr code matched, pass username to username reader
func (v *StateManager) onQRCodeScanned(QRCode string, username string) {
	if v.newUserIdenitfyStatus != waittingToBeIdentified {
		return
	}

	if v.identificationID != QRCode {
		log.Info().Msg("QRCode does not match")
	}

	log.Info().Msgf("New user identified by qr code: %s", QRCode)

	v.newUserIdentifiedHandler(username)
}
