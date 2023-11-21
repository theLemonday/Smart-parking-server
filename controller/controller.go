package controller

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Impl struct {
	client mqtt.Client
}

func NewController(c mqtt.Client) Impl {
	return Impl{
		client: c,
	}
}
