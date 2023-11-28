package handler

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CarDetectionMsg struct {
	Detected bool `json:"detected"`
}

func (h MQTTMsgHandler) CarSensorDetectedHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		msg, err := util.UnmarshalByte2JsonData[CarDetectionMsg](m.Payload())
		if err != nil {
			return
		}

		if m.Topic() == mqtt_client.IRGoInDirection {
			h.stateManager.OnCarGoInDetected(msg.Detected)
			return
		}

		if m.Topic() == mqtt_client.IRGoOutDirection {
			h.stateManager.OnCarGoOutDetected(msg.Detected)
			return
		}

		h.stateManager.OnCarGoInSlotDetected(strings.TrimPrefix(m.Topic(), mqtt_client.IRSensorPrefix), msg.Detected)
	}
}
