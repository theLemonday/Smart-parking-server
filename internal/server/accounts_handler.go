package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/database"
)

func unauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func (s *SmartParkingIotService) checkUserCredentials(w http.ResponseWriter, r *http.Request) *database.Account {
	username, password, ok := r.BasicAuth()
	if !ok {
		log.Info().Msgf("User %s failed to authenticate", username)
		unauthorizedResponse(w)
		return nil
	}

	account, err := s.AccountsDAO.AuthenticateUser(username, password)
	if err != nil {
		unauthorizedResponse(w)
		return nil
	}

	return &account
}
