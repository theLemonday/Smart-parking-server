package domain

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
)

type ControllerRepo interface {
	OpenBarrier()
	CloseBarrier()
	DisplayShowWelcome()
	DisplayShowBill(bill user.PaymentBill)
	DisplayShowSeeYouAgain()
	TurnLEDOn(ledTopic mqtt_client.LEDTopic)
	TurnLEDOff(ledTopic mqtt_client.LEDTopic)
}
