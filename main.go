package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/db"
	"github.com/thelemonday/smart-parking-iot-server/topic"
	"github.com/thelemonday/smart-parking-iot-server/topic/handler"
	"github.com/thelemonday/smart-parking-iot-server/viewmodel"
)

func main() {
	_db := db.SetupCacheDb()
	client := SetupMQTTClient(MainClientConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic().Err(token.Error()).Msg("")
	}

	_controller := controller.SetupController(client)
	_viewmodel := viewmodel.SetupViewmodel(client, _controller, _db)
	_handler := handler.SetupHandler(client, &_viewmodel)
	handlers := MAPSubTopic2MessageHandler{
		topic.IRGoInDirection: {QoS: 2, Handler: _handler.IRSensorGoInHandler()},
		topic.RFIDSubTop:      {QoS: 2, Handler: _handler.RFID()},
	}
	ClientSubTopics(client, handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go _viewmodel.HandleStatuses()

	<-quit
	log.Info().Msg("Shutdown client")
}
