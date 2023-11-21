package server

// type accountMsg struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type authenticationResponse struct {
// 	Token string `json:"token"`
// }

// func (s *SmartParkingIotWebsocketServer) userAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
// 	var msg accountMsg

// 	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	account, err := s.AccountsDAO.AuthenticateUser(msg.Username, msg.Password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
// 		"username": account.Username,
// 	})
// 	signedString, err := t.SignedString(s.key)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	log.Info().Msgf("Create new jwt token: %s", signedString)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(authenticationResponse{Token: signedString})
// }
