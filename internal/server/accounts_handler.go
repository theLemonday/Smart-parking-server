package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

type authenticationResponse struct {
	Token string `json:"token"`
}

func (s *SmartParkingIotService) userAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	account := s.checkUserCredentials(w, r)
	if account == nil {
		return
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": account.Username,
	})

	signedString, err := t.SignedString(s.key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Msgf("Create new jwt token: %s", signedString)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authenticationResponse{Token: signedString})
}
