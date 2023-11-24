package controller

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
)

type TestImpl struct{}

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
