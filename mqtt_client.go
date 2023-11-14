package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/secret"
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

func SetupMQTTClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%d", secret.Protocal, secret.Broker, secret.Port))
	opts.SetClientID(secret.ClientId)
	opts.SetUsername(secret.Username)
	opts.SetPassword(secret.Password)
	opts.SetDefaultPublishHandler(msgPubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}
