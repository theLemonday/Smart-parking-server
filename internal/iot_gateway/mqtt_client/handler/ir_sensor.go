package handler

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CarDetectionMsg struct {
	Detected bool `json:"detected"`
}

func (i Impl) CarSensorDetectedHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		msg, err := util.UnmarshalByte2JsonData[CarDetectionMsg](m.Payload())
		if err != nil {
			return
		}

		if m.Topic() == mqtt_client.IRGoInDirection {
			i.stateManager.OnCarGoInDetected(msg.Detected)
			return
		}

		if m.Topic() == mqtt_client.IRGoOutDirection {
			i.stateManager.OnCarGoOutDetected(msg.Detected)
			return
		}

		i.stateManager.OnCarGoInSlotDetected(strings.TrimPrefix(m.Topic(), mqtt_client.IRSensorPrefix), msg.Detected)
	}
}
