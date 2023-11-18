package mqtt_client

import (
	"fmt"
	secret "github.com/thelemonday/smart-parking-iot-server/configs"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

var msgPubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	log.Info().Msgf("Received from topic: %s message: %s\n", m.Payload(), m.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info().Msg("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Info().Msgf("Connect lost: %v", err)
}

func SetupMQTTClient(config secret.MQTTConfig) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%d", config.Protocol, config.Broker, config.Port))
	opts.SetClientID(config.ClientId)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetDefaultPublishHandler(msgPubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}

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
