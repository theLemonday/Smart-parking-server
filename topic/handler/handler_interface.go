package handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/viewmodel"
)

type HandlerRepo interface {
	IRSensorGoInHandler(c mqtt.Client, m mqtt.Message)
}

type HandlerImpl struct {
	c  mqtt.Client
	vm *viewmodel.Viewmodel
}

func SetupHandler(c mqtt.Client, viewmodel *viewmodel.Viewmodel) *HandlerImpl {
	return &HandlerImpl{
		c:  c,
		vm: viewmodel,
	}
}
