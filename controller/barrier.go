package controller

import (
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

func (c Impl) OpenBarrier() {
	c.client.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: open,
	}))
}

func (c Impl) CloseBarrier() {
	c.client.Publish(topic.GateBarrierPubTop, 2, false, util.MarshalJsonData2Byte(barrierControlMsg{
		Action: close,
	}))
}
