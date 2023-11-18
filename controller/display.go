package controller

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

type DisplayMessageType string

const (
	Text DisplayMessageType = "text"
	QR   DisplayMessageType = "qr"
)

type DisplayControl struct {
	Type DisplayMessageType `json:"type"`
	Msg  string             `json:"msg"`
}

func DisplayShowText(c mqtt.Client, text string) {
}

func DisplayShowQRCode(c mqtt.Client, text string) {
	token := c.Publish(topic.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: QR,
		Msg:  text,
	}))
	util.TokenWaitAndLog(token)
}

func (c _ControllerImpl) DisplayShowText(text string) {
	token := c.client.Publish(topic.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: Text,
		Msg:  text,
	}))
	util.TokenWaitAndLog(token)
}

func (c _ControllerImpl) DisplayShowQRCode(text string) {
	token := c.client.Publish(topic.OLEDPubTop, 0, false, util.MarshalJsonData2Byte(DisplayControl{
		Type: QR,
		Msg:  text,
	}))
	util.TokenWaitAndLog(token)
}
