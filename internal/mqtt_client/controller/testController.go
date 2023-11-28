package controller

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
)

type TestImpl struct{}

func (t TestImpl) DisplayShowWelcome() {
	log.Info().Msg("display show welcome")
}

func (t TestImpl) DisplayShowBill(bill user.PaymentBill) {
	log.Info().Msgf("display show bill uid: %s, %d", bill.RFIDTag, bill.TotalCost)
}

func (t TestImpl) DisplayShowSeeYouAgain() {
	log.Info().Msg("display show see you again")
}

func (t TestImpl) OpenBarrier() {
	log.Info().Msg("open barrier")
}

func (t TestImpl) CloseBarrier() {
	log.Info().Msg("closed barrier")
}

func (t TestImpl) DisplayShowText(s string) {
	log.Info().Msgf("Display show text: %s", s)
}

func (t TestImpl) DisplayShowQRCode(s string) {
	log.Info().Msgf("Display show qrcode: %s", s)
}

func (t TestImpl) TurnLEDOn(ledTopic mqtt_client.LEDTopic) {
	log.Info().Msgf("Turn on led %s", ledTopic)
}

func (t TestImpl) TurnLEDOff(ledTopic mqtt_client.LEDTopic) {
	log.Info().Msgf("Turn off led %s", ledTopic)
}

func NewTestController() *TestImpl {
	return &TestImpl{}
}
