package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

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
