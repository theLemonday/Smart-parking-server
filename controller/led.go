package controller

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

func TurnLEDOn(c mqtt.Client, t topic.LEDTopic) {
	log.Info().Msgf("Turn %s on", t)

	c.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: on,
	}))
}

func TurnLEDOff(c mqtt.Client, t topic.LEDTopic) {
	log.Info().Msgf("Turn %s off", t)

	c.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: off,
	}))
}

func (c _ControllerImpl) TurnLEDOn(t topic.LEDTopic) {
	log.Info().Msgf("Turn %s on", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: on,
	}))
}

func (c _ControllerImpl) TurnLEDOff(t topic.LEDTopic) {
	log.Info().Msgf("Turn %s off", t)

	c.client.Publish(t, 0, false, util.MarshalJsonData2Byte(ledControlMsg{
		DesiredStatus: off,
	}))
}
