package iot_gateway

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
)

type IotGateway struct {
	c mqtt.Client
}

func NewIotGateway(conf mqtt_client.MQTTConfig) *IotGateway {
	return &IotGateway{
		c: mqtt_client.SetupMQTTClient(conf),
	}
}

func (ig *IotGateway) Connect() {
	if token := ig.c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal().Err(token.Error()).Msg("")
	}
}

func (ig *IotGateway) GetMQTTClient() mqtt.Client {
	return ig.c
}

type MAPSubTopic2MessageHandler = map[string]struct {
	QoS     byte
	Handler mqtt.MessageHandler
}

func (ig *IotGateway) SubscribeTopics(handlersMap MAPSubTopic2MessageHandler) {
	for k, v := range handlersMap {
		if token := ig.c.Subscribe(k, v.QoS, v.Handler); token.Wait() && token.Error() != nil {
			log.Error().Err(token.Error()).Msg("")
		}
		log.Info().Msgf("Subscribed topic: %s", k)
	}
}
