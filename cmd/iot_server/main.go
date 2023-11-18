package main

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic/handler"
	"github.com/thelemonday/smart-parking-iot-server/internal/viewmodel"
	"github.com/thelemonday/smart-parking-iot-server/internal/ws_server"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

func main() {
	_db := db.SetupCacheDb()

	go ws_server.RunBackendServer(_db)

	client := mqtt_client.SetupMQTTClient(clientConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic().Err(token.Error()).Msg("")
	}

	_controller := controller.SetupController(client)
	_viewmodel := viewmodel.SetupViewmodel(client, _controller, _db)
	_handler := handler.SetupHandler(client, &_viewmodel)
	handlers := mqtt_client.MAPSubTopic2MessageHandler{
		topic.IRGoInDirection: {QoS: 2, Handler: _handler.IRSensorGoInHandler()},
		topic.RFIDSubTop:      {QoS: 2, Handler: _handler.RFID()},
	}
	mqtt_client.ClientSubTopics(client, handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go _viewmodel.HandleStatuses()

	<-quit
	log.Info().Msg("Shutdown client")
}
