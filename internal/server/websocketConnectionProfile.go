package server

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/database"
)

type websocketConnectionProfile struct {
	*websocket.Conn
	account *database.Account
}

func newWebsocketConnectionProfile(conn *websocket.Conn, account *database.Account) *websocketConnectionProfile {
	return &websocketConnectionProfile{
		Conn:    conn,
		account: account,
	}
}

func (c *websocketConnectionProfile) close() {
	err := c.Conn.Close()
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	log.Info().Msgf("Close websocket connection for user %s", c.account.Username)
}
