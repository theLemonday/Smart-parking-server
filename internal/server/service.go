package server

import (
	"time"

	"github.com/rs/zerolog/log"
)

type WebsocketService interface {
	OnNewUserEnter(uid string, identifiedMethod string, goInTimestamp time.Time)
	OnSlotStatusChanged(slotId string, occupied bool)
	OnUserLeave(uid string)
}

type newUserEnterMsg struct {
	Uid              string    `json:"uid"`
	IdentifiedMethod string    `json:"type"`
	GoInTimestamp    time.Time `json:"goInTimestamp"`
}

func (s *SmartParkingIotService) OnNewUserEnter(uid string, identifiedMethod string, goInTimestamp time.Time) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&newUserEnterMsg{
		Uid:              uid,
		IdentifiedMethod: identifiedMethod,
		GoInTimestamp:    goInTimestamp,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to send new user enter msg to monitor")
		return
	}
}

type slotStatusChangedMsg struct {
	SlotId   string `json:"slotId"`
	Occupied bool   `json:"occupied"`
}

func (s *SmartParkingIotService) OnSlotStatusChanged(slotId string, occupied bool) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&slotStatusChangedMsg{
		SlotId:   slotId,
		Occupied: occupied,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to send slot status change msg to monitor")
		return
	}
}

type userLeaveMsg struct {
	UserUid string `json:"usernameUid"`
}

func (s *SmartParkingIotService) OnUserLeave(uid string) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&userLeaveMsg{
		UserUid: uid,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to send user leave msg to monitor")
		return
	}
}
