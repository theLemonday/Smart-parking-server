package controller

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
)

type Repo interface {
	OpenBarrier()
	CloseBarrier()
	DisplayShowText(string)
	DisplayShowQRCode(string)
	TurnLEDOn(ledTopic mqtt_client.LEDTopic)
	TurnLEDOff(ledTopic mqtt_client.LEDTopic)
}
