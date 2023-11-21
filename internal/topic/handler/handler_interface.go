package handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/internal/state-manager"
)

type HandlerRepo interface {
	IRSensorGoInHandler(c mqtt.Client, m mqtt.Message)
}

type HandlerImpl struct {
	c  mqtt.Client
	vm *state_manager.StateManager
}

func SetupHandler(c mqtt.Client, viewmodel *state_manager.StateManager) *HandlerImpl {
	return &HandlerImpl{
		c:  c,
		vm: viewmodel,
	}
}
