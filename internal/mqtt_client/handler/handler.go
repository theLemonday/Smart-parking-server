package handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	statemanager "github.com/thelemonday/smart-parking-iot-server/internal/state-manager"
)

type MQTTMsgHandler struct {
	c            mqtt.Client
	stateManager *statemanager.StateManager
}

func SetupHandler(c mqtt.Client, stateManager *statemanager.StateManager) *MQTTMsgHandler {
	return &MQTTMsgHandler{
		c:            c,
		stateManager: stateManager,
	}
}
