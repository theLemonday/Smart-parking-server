package viewmodel

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/topic"
	"github.com/thelemonday/smart-parking-iot-server/util"
)

type _IdentifyStatus int

const (
	_Unknown _IdentifyStatus = iota
	_WaittingToBeIdentified
	_RFID
	_QR
)

type _IdentifyMethod string

const (
	_IdentifiedByQr   _IdentifyMethod = "qr"
	_IdentifiedByRFID _IdentifyMethod = "rfid"
)

type _UserProfile struct {
	id         string
	identifyBy _IdentifyMethod
}

type state struct {
	carGoIn                bool
	carGoOut               bool
	_newUserIdenitfyStatus _IdentifyStatus
	_currentUID            string
	userProfiles           map[string]_UserProfile
	isGoIn                 bool
	// mu                     sync.Mutex
}

type Viewmodel struct {
	mqttClient     mqtt.Client
	controllerImpl controller.ControllerRepo
	state
	wg                sync.WaitGroup
	_IRSensorIn       chan bool
	_IRSensorOut      chan bool
	_IRSensorCarSlots chan [3]bool
	_RFIDReader       chan string
}

func SetupViewmodel(c mqtt.Client, _controller controller.ControllerRepo) Viewmodel {
	return Viewmodel{
		mqttClient:        c,
		controllerImpl:    _controller,
		_IRSensorIn:       make(chan bool),
		_IRSensorOut:      make(chan bool),
		_IRSensorCarSlots: make(chan [3]bool),
		_RFIDReader:       make(chan string),
		state: state{
			isGoIn:                 false,
			userProfiles:           make(map[string]_UserProfile),
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

func (v *Viewmodel) handleCarGoInOrGoOutDection(isGoIn bool) {
	uid := util.GenerateNewNanoID(10)
	if !isGoIn {
		uid = fmt.Sprintf("tt%s", uid)
	}

	// v.controllerImpl.TurnLEDOn(topic.RedLEDPubTop)
	// v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	// v.controllerImpl.DisplayShowQRCode(uid)

	v._currentUID = uid
	v._newUserIdenitfyStatus = _WaittingToBeIdentified
	v.isGoIn = isGoIn

	log.Info().Msg("Waiting the user to be identified")

	v.wg.Done()
}

func (v *Viewmodel) handleNoCar() {
	// actions
	// v.controllerImpl.TurnLEDOff(topic.GreenLEDPubTop)
	// v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
	// v.controllerImpl.CloseBarrier()
	// v.controllerImpl.DisplayShowText("Waiting for new user")

	v._newUserIdenitfyStatus = _Unknown
	v.carGoIn = false
	v.carGoOut = false

	log.Info().Msg("Reset viewmodel")

	v.wg.Done()
}

func (v *Viewmodel) HandleStatuses() {
	for {
		select {
		case detected := <-v._IRSensorIn:
			v.wg.Wait()
			log.Info().Msg("Car goes in detected")

			if !detected {
				v.wg.Add(1)
				v.handleNoCar()
				return
			}

			if v._newUserIdenitfyStatus == _Unknown {
				v.wg.Add(1)
				v.handleCarGoInOrGoOutDection(true)
				return
			}

		case rfidUId := <-v._RFIDReader:
			if v._newUserIdenitfyStatus == _WaittingToBeIdentified {
				v.userProfiles[v._currentUID] = _UserProfile{
					id:         rfidUId,
					identifyBy: _IdentifiedByRFID,
				}
				log.Info().Msgf("New user identified by rfid: %s", rfidUId)
				v.controllerImpl.OpenBarrier()
				v.controllerImpl.DisplayShowText("Welcome to our smart parking system")
				v.controllerImpl.TurnLEDOff(topic.RedLEDPubTop)
				v.controllerImpl.TurnLEDOn(topic.GreenLEDPubTop)
				v._newUserIdenitfyStatus = _RFID
			}

		case v.carGoOut = <-v._IRSensorOut:
			if !v.carGoOut {
				if !v.carGoIn {
					v.handleNoCar()
				}
				return
			}

			if v._newUserIdenitfyStatus == _Unknown {
				v.handleCarGoInOrGoOutDection(false)
			}
		}
	}
}
