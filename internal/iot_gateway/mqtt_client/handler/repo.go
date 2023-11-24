package handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/state-manager"
)

type Repo interface {
	IRSensorGoInHandler(c mqtt.Client, m mqtt.Message)
}

type Impl struct {
	c            mqtt.Client
	stateManager *state_manager.StateManager
}

func SetupHandler(c mqtt.Client, stateManager *state_manager.StateManager) *Impl {
	return &Impl{
		c:            c,
		stateManager: stateManager,
	}
}
