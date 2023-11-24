package server

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/database"
)

type paymentRequestMsg struct {
	Username string `json:"username"`
}

func (s *SmartParkingIotService) onUserConfirmPayment(payload *json.RawMessage) {

}

func (s *SmartParkingIotService) OnUserGoOutIdentified(bill *database.PaymentBill) {
	err := s.clients[bill.Username].Conn.WriteJSON(bill)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
}
