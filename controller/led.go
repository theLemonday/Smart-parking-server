package controller

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

type desiredStatus string

const (
	on  desiredStatus = "on"
	off desiredStatus = "off"
)

type ledControlMsg struct {
	DesiredStatus desiredStatus `json:"status"`
	Timeout       uint          `json:"timeout,omitempty"`
}

func (c Impl) TurnLEDOn(t topic.LEDTopic) {
	log.Info().Msgf("Turn %s on", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: on,
	}))
}

func (c Impl) TurnLEDOff(t topic.LEDTopic) {
	log.Info().Msgf("Turn %s off", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: off,
	}))
}
