package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"

	"github.com/thelemonday/smart-parking-iot-server/db"
	"github.com/thelemonday/smart-parking-iot-server/internal/server/presenter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SmartParkingIotWebsocketServer struct {
	http.Handler
	db.AccountsDAO
	clients map[*userWSConn]bool
	key     *ecdsa.PrivateKey
}

func NewSmartParkingIotServer(db db.AccountsDAO) *SmartParkingIotWebsocketServer {
	s := new(SmartParkingIotWebsocketServer)
	router := http.NewServeMux()

	router.Handle("/ws", http.HandlerFunc(s.webSocket))

	s.Handler = router
	s.AccountsDAO = db
	s.clients = make(map[*userWSConn]bool)

	curve := elliptic.P256()
	s.key, _ = ecdsa.GenerateKey(curve, rand.Reader)

	return s
}

func (s *SmartParkingIotWebsocketServer) ListenAndServe(port string) {
	log.Info().Msgf("Started websocket server on %s", port)
	if err := http.ListenAndServe(port, s.Handler); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (s *SmartParkingIotWebsocketServer) webSocket(w http.ResponseWriter, r *http.Request) {
	account := s.authenticateUser(w, r)
	if account == nil {
		return
	}

	log.Info().Msgf("User %s connected", account.Username)
	log.Info().Msg("Upgrade connection to websocket")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	conn := userWSConn{Conn: ws, account: account}
	defer func(conn *userWSConn) {
		s.removeWebsocketConnectionFromConnectionsList(conn)
		conn.close()
	}(&conn)

	s.clients[&conn] = true
	if err = s.onConnected(&conn); err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	s.listenWebsocketAndServer(&conn)
}

func (s *SmartParkingIotWebsocketServer) onConnected(conn *userWSConn) error {
	return conn.Conn.WriteJSON(presenter.NewAccountAuthenticationSuccessResponse(conn.account))
}

type websocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (s *SmartParkingIotWebsocketServer) listenWebsocketAndServer(conn *userWSConn) {
	for {
		var msg websocketMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Info().Err(err).Msg("")
			return
		}
	}
}

func (s *SmartParkingIotWebsocketServer) removeWebsocketConnectionFromConnectionsList(conn *userWSConn) {
	log.Info().Msg("Delete websocket connection")
	delete(s.clients, conn)
}
