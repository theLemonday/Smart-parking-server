package controller

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
)

type action string

const (
	open   action = "open"
	closed action = "closed"
)

type barrierControlMsg struct {
	Action action `json:"action"`
}

func (c Impl) OpenBarrier() {
	c.client.Publish(mqtt_client.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: open,
	}))
}

func (c Impl) CloseBarrier() {
	c.client.Publish(mqtt_client.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: closed,
	}))
}
