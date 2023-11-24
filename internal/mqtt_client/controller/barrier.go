package controller

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
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

func (i Impl) OpenBarrier() {
	i.client.Publish(mqtt_client.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: open,
	}))
}

func (i Impl) CloseBarrier() {
	i.client.Publish(mqtt_client.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: closed,
	}))
}
