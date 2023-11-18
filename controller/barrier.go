package controller

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

type action string

const (
	open  action = "open"
	close action = "close"
)

type barrierControlMsg struct {
	Action action `json:"action"`
}

func OpenBarrier(c mqtt.Client) {
	c.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: open,
	}))
}

func CloseBarrier(c mqtt.Client) {
	c.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: close,
	}))
}

func (c _ControllerImpl) OpenBarrier() {
	c.client.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: open,
	}))
}

func (c _ControllerImpl) CloseBarrier() {
	c.client.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: close,
	}))
}
