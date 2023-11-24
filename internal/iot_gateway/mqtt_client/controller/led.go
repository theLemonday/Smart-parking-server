package controller

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
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

func (c Impl) TurnLEDOn(t mqtt_client.LEDTopic) {
	log.Info().Msgf("Turn %s on", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: on,
	}))
}

func (c Impl) TurnLEDOff(t mqtt_client.LEDTopic) {
	log.Info().Msgf("Turn %s off", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: off,
	}))
}
