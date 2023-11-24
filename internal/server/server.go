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
	clients map[*websocketConnectionProfile]bool
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
	s.clients = make(map[*websocketConnectionProfile]bool)

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
	defer func(conn *websocketConnectionProfile) {
		s.removeWebsocketConnectionFromConnectionsList(conn)
		conn.close()
	}(connectionProfile)

	s.clients[connectionProfile] = true
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
	}
}

func (s *SmartParkingIotService) removeWebsocketConnectionFromConnectionsList(conn *websocketConnectionProfile) {
	log.Info().Msg("Delete websocket connection")
	delete(s.clients, conn)
}
