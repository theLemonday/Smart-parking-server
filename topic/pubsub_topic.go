package topic

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type LEDTopic = string
type IRSensorTopic = string

const (
	OLEDPubTop                      = "smart-parking/gate/OLED"
	GateBarrierPubTop               = "smart-parking/gate/barrier"
	RFIDSubTop                      = "smart-parking/gate/RFID/out"
	IRTopicPrefix                   = "smart-parking/gate/IR/+"
	IRGoInDirection   IRSensorTopic = "smart-parking/gate/IR/in"
	IRGoOutDirection  IRSensorTopic = "smart-parking/gate/IR/out"
	GreenLEDPubTop    LEDTopic      = "smart-parking/gate/greenLED"
	RedLEDPubTop      LEDTopic      = "smart-parking/gate/redLED"
)

type QoSAndHandler struct {
	QoS     byte
	Handler mqtt.MessageHandler
}
type MAPSubTopic2MessageHandler = map[string]QoSAndHandler

func ClientSubTopics(c mqtt.Client, handlersMap MAPSubTopic2MessageHandler) {
	for k, v := range handlersMap {
		if token := c.Subscribe(k, v.QoS, v.Handler); token.Wait() && token.Error() != nil {
			log.Error().Err(token.Error()).Msg("")
		}
		log.Info().Msgf("Subscribed topic: %s", k)
	}
}
