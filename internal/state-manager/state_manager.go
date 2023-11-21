package state_manager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/database"
	"github.com/thelemonday/smart-parking-iot-server/internal/server"
)

type StateManager struct {
	mqttClient       mqtt.Client
	controllerImpl   controller.Repo
	carParkStatusDb  database.CarParkStatusDAO
	websocketService server.WebsocketService
	*state
	_IRSensorIn       chan bool
	_IRSensorOut      chan bool
	_IRSensorCarSlots chan [3]bool
	_RFIDUid          chan string
	_QRCodeScanner    chan [2]string
}

func NewStateManager(c mqtt.Client, _controller controller.Repo, database database.CarParkStatusDAO, websocketService server.WebsocketService) *StateManager {
	return &StateManager{
		mqttClient:        c,
		controllerImpl:    _controller,
		carParkStatusDb:   database,
		websocketService:  websocketService,
		_IRSensorIn:       make(chan bool),
		_IRSensorOut:      make(chan bool),
		_IRSensorCarSlots: make(chan [3]bool),
		_RFIDUid:          make(chan string),
		_QRCodeScanner:    make(chan [2]string),
		state:             newState(),
	}
}

func (v *StateManager) HandleStatuses() {
	for {
		select {
		case v.carGoIn = <-v._IRSensorIn:
			v.carGoInHandler()

		case uid := <-v._RFIDUid:
			v.rfidTagReadHandler(uid)

		case info := <-v._QRCodeScanner:
			qrcode, username := info[0], info[1]
			v.qrCodeScannedHandler(qrcode, username)

		case v.carGoOut = <-v._IRSensorOut:
			if !v.carGoOut {
				if !v.carGoIn {
					v.onNoCarDetection()
					continue
				}
				continue
			}

			if v.newUserIdentifyStatus == unknown {
				v.carGoInHandler()
				continue
			}
		}
	}
}
