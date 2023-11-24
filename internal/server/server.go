package server

import (
	"encoding/json"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SmartParkingIotService struct {
	http.Handler
	monitor *websocket.Conn
}

func (s *SmartParkingIotService) OnUserGoOutIdentified(bill *domain.PaymentBill) {
	err := s.monitor.WriteJSON(bill)
	if err != nil {
		log.Error().Err(err).Msg("send bill to monitor")
		return
	}
}

func NewSmartParkingIotServer() *SmartParkingIotService {
	s := new(SmartParkingIotService)
	router := http.NewServeMux()

	router.Handle("/ws", http.HandlerFunc(s.webSocket))

	s.Handler = router

	return s
}

func (s *SmartParkingIotService) ListenAndServe(port string) {
	log.Info().Msgf("Started websocket server on %s", port)
	if err := http.ListenAndServe(port, s.Handler); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (s *SmartParkingIotService) webSocket(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Upgrade connection to websocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}(conn)

	s.monitor = conn
	s.listenWebsocketAndServer(conn)
}

type websocketMessage struct {
	Type string `json:"type"`
	Data json.RawMessage
}

func (s *SmartParkingIotService) listenWebsocketAndServer(conn *websocket.Conn) {
	for {
		var msg websocketMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Info().Err(err).Msg("")
			return
		}

		//switch msg.Type {
		//case "pay-request":
		//	s.onUserConfirmPayment(&msg.Data)
		//}
	}
}
