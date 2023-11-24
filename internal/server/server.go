package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"

	"github.com/thelemonday/smart-parking-iot-server/database"
	"github.com/thelemonday/smart-parking-iot-server/internal/server/presenter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SmartParkingIotService struct {
	http.Handler
	database.AccountsDAO
	clients map[string]*websocketConnectionProfile
	monitor *websocket.Conn
	key     *ecdsa.PrivateKey
}

func NewSmartParkingIotServer(db database.AccountsDAO) *SmartParkingIotService {
	s := new(SmartParkingIotService)
	router := http.NewServeMux()

	router.Handle("/ws", http.HandlerFunc(s.webSocket))
	router.Handle("/authentication", http.HandlerFunc(s.userAuthenticationHandler))

	s.Handler = router
	s.AccountsDAO = db
	s.clients = make(map[string]*websocketConnectionProfile)

	curve := elliptic.P256()
	var err error
	s.key, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	return s
}

func (s *SmartParkingIotService) ListenAndServe(port string) {
	log.Info().Msgf("Started websocket server on %s", port)
	if err := http.ListenAndServe(port, s.Handler); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (s *SmartParkingIotService) webSocket(w http.ResponseWriter, r *http.Request) {
	account := s.checkUserCredentials(w, r)
	if account == nil {
		return
	}

	log.Info().Msgf("User %s connected", account.Username)
	log.Info().Msg("Upgrade connection to websocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	connectionProfile := newWebsocketConnectionProfile(conn, account)
	defer s.removeWebsocketConnectionFromConnectionsList(account.Username)

	s.clients[account.Username] = connectionProfile
	if err = s.onConnected(connectionProfile); err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	s.listenWebsocketAndServer(connectionProfile)
}

func (s *SmartParkingIotService) onConnected(conn *websocketConnectionProfile) error {
	return conn.Conn.WriteJSON(presenter.NewAccountAuthenticationSuccessResponse(conn.account))
}

type websocketMessage struct {
	Type string `json:"type"`
	Data json.RawMessage
}

func (s *SmartParkingIotService) listenWebsocketAndServer(conn *websocketConnectionProfile) {
	for {
		var msg websocketMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Info().Err(err).Msg("")
			return
		}

		switch msg.Type {
		case "pay-request":
			s.onUserConfirmPayment(&msg.Data)
		}
	}
}

func (s *SmartParkingIotService) removeWebsocketConnectionFromConnectionsList(username string) {
	log.Info().Msg("Delete websocket connection")
	s.clients[username].close()
	delete(s.clients, username)
}
