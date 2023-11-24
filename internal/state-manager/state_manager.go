package state_manager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain/repo"
	"sync"
)

type StateManager struct {
	mqttClient                mqtt.Client
	controllerImpl            repo.ControllerRepo
	carParkStatusDb           domain.CarParkStatusDAO
	transfer2websocketService domain.StateManager2Websocket

	mu    *sync.Mutex
	state *state
}

func NewStateManager(c mqtt.Client, _controller repo.ControllerRepo, database domain.CarParkStatusDAO) *StateManager {
	return &StateManager{
		mqttClient:      c,
		controllerImpl:  _controller,
		carParkStatusDb: database,
		state:           newCarParkState(),
	}
}

func (s *StateManager) SetWebsocketService(websocket domain.StateManager2Websocket) {
	s.transfer2websocketService = websocket
}
