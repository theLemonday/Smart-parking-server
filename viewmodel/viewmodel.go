package viewmodel

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

type _IdentifyStatus int

const (
	_Unknown _IdentifyStatus = iota
	_WaittingToBeIdentified
	_Identified
)

type _IdentifyMethod string

const (
	_IdentifiedByQr   _IdentifyMethod = "qr"
	_IdentifiedByRFID _IdentifyMethod = "rfid"
)

type state struct {
	carGoIn                bool
	carGoOut               bool
	_newUserIdenitfyStatus _IdentifyStatus
	_currentUID            string
	isGoIn                 bool
}

type Viewmodel struct {
	mqttClient     mqtt.Client
	controllerImpl controller.ControllerRepo
	*state
	database          db.DatabaseRepo
	_IRSensorIn       chan bool
	_IRSensorOut      chan bool
	_IRSensorCarSlots chan [3]bool
	_RFIDReader       chan string
	resetting         chan bool
}

func SetupViewmodel(c mqtt.Client, _controller controller.ControllerRepo, database db.DatabaseRepo) Viewmodel {
	return Viewmodel{
		mqttClient:        c,
		controllerImpl:    _controller,
		database:          database,
		_IRSensorIn:       make(chan bool),
		_IRSensorOut:      make(chan bool),
		_IRSensorCarSlots: make(chan [3]bool),
		_RFIDReader:       make(chan string),
		resetting:         make(chan bool),
		state: &state{
			isGoIn:                 false,
			_newUserIdenitfyStatus: _Unknown,
			carGoOut:               false,
			carGoIn:                false,
		},
	}
}

func (v *Viewmodel) OnCarGoInDetected(detected bool) {
	v._IRSensorIn <- detected
}

func (v *Viewmodel) OnRFIDRead(uid string) {
	v._RFIDReader <- uid
}

func (v *Viewmodel) HandleStatuses() {
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

			if v._newUserIdenitfyStatus == _Unknown {
				v.onCarGoIn()
				continue
			}

		case uid := <-v._RFIDReader:
			v.onRFIDRead(uid)

		case v.carGoOut = <-v._IRSensorOut:
			if !v.carGoOut {
				if !v.carGoIn {
					v.onNoCarDetection()
					continue
				}
				continue
			}

			if v._newUserIdenitfyStatus == _Unknown {
				v.onCarGoIn()
				continue
			}
		}
	}
}
