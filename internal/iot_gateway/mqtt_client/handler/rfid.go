package handler

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type _RFIDMsg struct {
	Uid string `json:"uid"`
}

func (i Impl) RFID() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		log.Info().Msg("New RFID card read")
		var msg _RFIDMsg
		if err := json.Unmarshal(m.Payload(), &msg); err != nil {
			log.Error().Err(err).Msg("")
		}

		i.stateManager.OnRFIDTagRead(msg.Uid)
	}
}
