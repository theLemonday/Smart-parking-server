package controller

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

type DisplayMessageType string

const (
	Welcome     DisplayMessageType = "welcome"
	Bill        DisplayMessageType = "bill"
	SeeYouAgain DisplayMessageType = "see-you-again"
)

type DisplayControl struct {
	Type DisplayMessageType `json:"type"`
	Bill user.PaymentBill   `json:"bill,omitempty"`
}

func (i Impl) DisplayShowWelcome() {
	token := i.client.Publish(mqtt_client.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: Welcome,
	}))
	util.TokenWaitAndLog(token)
}

func (i Impl) DisplayShowBill(bill user.PaymentBill) {
	token := i.client.Publish(mqtt_client.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: Bill,
		Bill: bill,
	}))
	util.TokenWaitAndLog(token)
}

func (i Impl) DisplayShowSeeYouAgain() {
	token := i.client.Publish(mqtt_client.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: SeeYouAgain,
	}))
	util.TokenWaitAndLog(token)
}
