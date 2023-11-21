package handler

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CarDetectionMsg struct {
	Detected bool `json:"detected"`
}

func (h Impl) CarSensorDetectedHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		msg, err := util.UnmarshalByte2JsonData[CarDetectionMsg](m.Payload())
		if err != nil {
			return
		}

		if m.Topic() == topic.IRGoInDirection {
			h.stateManager.OnCarGoInDetected(msg.Detected)
			return
		}

		if m.Topic() == topic.IRGoOutDirection {
			h.stateManager.OnCarGoOutDetected(msg.Detected)
			return
		}

		h.stateManager.OnCarGoInSlotDetected(m.Topic(), msg.Detected)
	}
}
