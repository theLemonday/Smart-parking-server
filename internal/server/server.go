package ws_server

import (
	"net/http"

	"github.com/thelemonday/smart-parking-iot-server/db"
)

type SmartParkingIotWebsocketServer struct {
	http.Handler
	db.AccountsDatabaseRepo
	clients map[*userWSConn]bool
}

func NewSmartParkingIotServer(db db.AccountsDatabaseRepo) *SmartParkingIotWebsocketServer {
	s := new(SmartParkingIotWebsocketServer)
	router := http.NewServeMux()

	router.Handle("/ws", http.HandlerFunc(s.webSocket))

	s.Handler = router
	s.AccountsDatabaseRepo = db

	return s
}

func (s *SmartParkingIotWebsocketServer) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.Handler)
}
