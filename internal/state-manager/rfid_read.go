package state_manager

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
)

func (v *StateManager) onUserIdentified() {
	v.controllerImpl.OpenBarrier()
	v.controllerImpl.DisplayShowText("Welcome to our smart parking system")
	v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	v.controllerImpl.TurnLEDOn(topic.GreenLEDPubTop)
}

func (v *StateManager) rfidTagReadHandler(uid string) {
	if v.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	if !v.carParkStatusDb.IsRFIDTagValid(uid) {
		log.Info().Msg("Invalid RFID tags, do nothing")
		return
	}

	log.Info().Msgf("New user identified by rfid: %s", uid)

	v.newUserIdentifyStatus = identified
	user := v.carParkStatusDb.NewUserIdentifiedByRFID(uid)

	v.websocketService.OnNewUserEnter(user.Id, user.IdentifiedBy, user.GoInTimestamp)

	v.onUserIdentified()
}

func (v *StateManager) qrCodeScannedHandler(QRCode string, username string) {
	if v.newUserIdentifyStatus != waitingToBeIdentified {
		return
	}

	if v.identificationID != QRCode {
		log.Info().Msg("QRCode does not match")
		return
	}

	log.Info().Msgf("New user identified by qr code: %s", QRCode)

	v.newUserIdentifyStatus = identified
	user := v.carParkStatusDb.NewUserIdentifiedByQRCode(username)

	v.websocketService.OnNewUserEnter(user.Id, user.IdentifiedBy, user.GoInTimestamp)

	v.onUserIdentified()
}
