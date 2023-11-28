package server

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *SmartParkingIotService) OnNewUserEnter(user *user.User) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&struct {
		Uid           string    `json:"uid"`
		GoInTimestamp time.Time `json:"goInTimestamp"`
	}{user.Id, user.GoInTimestamp})
	if err != nil {
		log.Error().Err(err).Msg("failed to send new user enter msg to monitor")
		return
	}
}

func (s *SmartParkingIotService) OnSlotStatusChanged(slotId string, occupied bool) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&struct {
		SlotId   string `json:"slotId"`
		Occupied bool   `json:"occupied"`
	}{slotId, occupied})
	if err != nil {
		log.Error().Err(err).Msg("failed to send slot status change msg to monitor")
	}
}

func (s *SmartParkingIotService) OnUserLeave(uid string) {
	if s.monitor == nil {
		return
	}

	err := s.monitor.WriteJSON(&struct {
		UserUid string `json:"usernameUid"`
	}{uid})
	if err != nil {
		log.Error().Err(err).Msg("failed to send user leave msg to monitor")
	}
}
