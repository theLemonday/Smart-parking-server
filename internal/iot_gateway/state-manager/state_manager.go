package state_manager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/database"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client/controller"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain"
	"sync"
)

type StateManager struct {
	mqttClient                mqtt.Client
	controllerImpl            controller.Repo
	carParkStatusDb           database.CarParkStatusDAO
	transfer2websocketService domain.StateManager2Websocket

	mu    *sync.Mutex
	state *state
}

func NewStateManager(c mqtt.Client, _controller controller.Repo, database database.CarParkStatusDAO) *StateManager {
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
