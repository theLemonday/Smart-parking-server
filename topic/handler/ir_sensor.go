package handler

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type CarDectectionMsg struct {
	Detected bool `json:"detected"`
}

func (h HandlerImpl) IRSensorGoInHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		var msg CarDectectionMsg
		if err := json.Unmarshal(m.Payload(), &msg); err != nil {
			log.Error().Err(err).Msg("")
			return
		}

		if msg.Detected {
			// log.Info().Msg("Car goes in detected")
			h.vm.OnCarGoInDetected(true)
		} else {
			// log.Info().Msg("No car goes in detected")
			h.vm.OnCarGoInDetected(false)
		}

	}
}
