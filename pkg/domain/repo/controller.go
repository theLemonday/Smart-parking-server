package repo

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain"
)

type ControllerRepo interface {
	OpenBarrier()
	CloseBarrier()
	DisplayShowWelcome()
	DisplayShowBill(bill domain.PaymentBill)
	DisplayShowSeeYouAgain()
	TurnLEDOn(ledTopic mqtt_client.LEDTopic)
	TurnLEDOff(ledTopic mqtt_client.LEDTopic)
}
