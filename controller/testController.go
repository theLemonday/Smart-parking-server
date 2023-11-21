package controller

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
)

type TestImpl struct{}

func (t TestImpl) OpenBarrier() {
	log.Info().Msg("open barrier")
}

func (t TestImpl) CloseBarrier() {
	log.Info().Msg("close barrier")
}

func (t TestImpl) DisplayShowText(s string) {
	log.Info().Msgf("Display show text: %s", s)
}

func (t TestImpl) DisplayShowQRCode(s string) {
	log.Info().Msgf("Display show qrcode: %s", s)
}

func (t TestImpl) TurnLEDOn(ledTopic topic.LEDTopic) {
	log.Info().Msgf("Turn on led %s", ledTopic)
}

func (t TestImpl) TurnLEDOff(ledTopic topic.LEDTopic) {
	log.Info().Msgf("Turn off led %s", ledTopic)
}

func NewTestController() *TestImpl {
	return &TestImpl{}
}
