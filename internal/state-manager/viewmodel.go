package state_manager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

type StateManager struct {
	mqttClient     mqtt.Client
	controllerImpl controller.ControllerRepo
	database       db.CarParkStatusDAO
	*state
	_IRSensorIn        chan bool
	_IRSensorOut       chan bool
	_IRSensorCarSlots  chan [3]bool
	usernameIdentified chan string
	_QRCodeScanner     chan string
}

func NewStateManager(c mqtt.Client, _controller controller.ControllerRepo, database db.CarParkStatusDAO) *StateManager {
	return &StateManager{
		mqttClient:         c,
		controllerImpl:     _controller,
		database:           database,
		_IRSensorIn:        make(chan bool),
		_IRSensorOut:       make(chan bool),
		_IRSensorCarSlots:  make(chan [3]bool),
		usernameIdentified: make(chan string),
		_QRCodeScanner:     make(chan string),
		state:              newState(),
	}
}

func (v *StateManager) OnCarGoInDetected(detected bool) {
	v._IRSensorIn <- detected
}

func (v *StateManager) OnNewUserIdentified(username string) {
	v.usernameIdentified <- username
}

func (v *StateManager) OnQRCodeScanned(QRCode, username string) {
	v._QRCodeScanner <- QRCode
}

func (v *StateManager) HandleStatuses() {
	for {
		select {
		case v.carGoIn = <-v._IRSensorIn:
			if !v.carGoIn {
				if !v.carGoOut {
					v.onNoCarDetection()
					continue
				}
				continue
			}

			log.Info().Msg("Car goes in detected")
			v.isGoIn = true

			if v.newUserIdentifyStatus == unknown {
				v.onCarGoIn()
				continue
			}

		case uid := <-v.usernameIdentified:
			v.newUserIdentifiedHandler(uid)

		case code := <-v._QRCodeScanner:
			v.onQRCodeScanned(code, "")

		case v.carGoOut = <-v._IRSensorOut:
			if !v.carGoOut {
				if !v.carGoIn {
					v.onNoCarDetection()
					continue
				}
				continue
			}

			if v.newUserIdentifyStatus == unknown {
				v.onCarGoIn()
				continue
			}
		}
	}
}
