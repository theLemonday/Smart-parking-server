package state_manager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/internal/domain"
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"sync"
)

type StateManager struct {
	mqttClient                mqtt.Client
	controllerImpl            domain.ControllerRepo
	userUseCase               user.UseCase
	transfer2websocketService domain.StateManager2Websocket

	mu    sync.Mutex
	state *state
}

func NewStateManager(c mqtt.Client, _controller domain.ControllerRepo, useCase user.UseCase) *StateManager {
	return &StateManager{
		mqttClient:     c,
		controllerImpl: _controller,
		userUseCase:    useCase,
		state:          newCarParkState(),
	}
}

func (s *StateManager) SetWebsocketService(websocket domain.StateManager2Websocket) {
	s.transfer2websocketService = websocket
}

func (s *StateManager) onNoCarDetection() {
	s.controllerImpl.TurnLEDOff(mqtt_client.GreenLEDPubTop)
	s.controllerImpl.TurnLEDOff(mqtt_client.RedLEDPubTop)
	s.controllerImpl.CloseBarrier()

	s.state.reset()
}
