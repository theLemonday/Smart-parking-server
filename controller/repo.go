package controller

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
)

type Repo interface {
	OpenBarrier()
	CloseBarrier()
	DisplayShowText(string)
	DisplayShowQRCode(string)
	TurnLEDOn(ledTopic topic.LEDTopic)
	TurnLEDOff(ledTopic topic.LEDTopic)
}
