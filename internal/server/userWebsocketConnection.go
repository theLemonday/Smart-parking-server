package ws_server

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

type userWSConn struct {
	*websocket.Conn
	profile *db.Account
}

func (c *userWSConn) close() {
	err := c.Conn.Close()
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	log.Info().Msgf("Close websocket connection for user %s", c.profile.Username)
}
