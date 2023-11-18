package ws_server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

type WsClient struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	database  db.DatabaseRepo
}

var upgrader = websocket.Upgrader{}

func SetupWebsocketServer(database db.DatabaseRepo) *WsClient {
	return &WsClient{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		database:  database,
	}
}

func (wsClient *WsClient) handleNewWsConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}(ws)

	wsClient.clients[ws] = true
	wsClient.onConnected(ws)
}

type authenticationMessage struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (wsClient *WsClient) onConnected(conn *websocket.Conn) {
	var msg authenticationMessage
	if err := conn.ReadJSON(&msg); err != nil {
		log.Info().Err(err).Msg("")
		wsClient.removeWebsocketConnectionFromConnectionsList(conn)
		return
	}

	if wsClient.database.UserAuthentication(msg.PhoneNumber, msg.Password) {
		wsClient.onAuthenticated(conn)
		wsClient.listenWebsocketAndServer(conn)
	} else {
		log.Info().Msg("Authentication failed: wrong phone number or password")
		wsClient.removeWebsocketConnectionFromConnectionsList(conn)
	}
}

func (wsClient *WsClient) onAuthenticated(conn *websocket.Conn) {
	err := conn.WriteJSON(wsClient.database.GetAllUsers())
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}

type websocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (wsClient *WsClient) listenWebsocketAndServer(conn *websocket.Conn) {
	for {
		var msg websocketMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Info().Err(err).Msg("")
			wsClient.removeWebsocketConnectionFromConnectionsList(conn)
			break
		}
	}
}

func (wsClient *WsClient) removeWebsocketConnectionFromConnectionsList(conn *websocket.Conn) {
	log.Info().Msg("Delete websocket connection")
	delete(wsClient.clients, conn)
}
