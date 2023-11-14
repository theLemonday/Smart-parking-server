package controller

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ControllerRepo interface {
	OpenBarrier()
	CloseBarrier()
	DisplayShowText(string)
	DisplayShowQRCode(string)
	TurnLEDOn(string)
	TurnLEDOff(string)
}

type _ControllerImpl struct {
	client mqtt.Client
}

func SetupController(c mqtt.Client) _ControllerImpl {
	return _ControllerImpl{
		client: c,
	}
}
