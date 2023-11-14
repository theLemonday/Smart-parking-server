package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/topic"
	"github.com/thelemonday/smart-parking-iot-server/topic/handler"
	"github.com/thelemonday/smart-parking-iot-server/viewmodel"
)

func main() {
	client := SetupMQTTClient()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	_controller := controller.SetupController(client)
	_viewmodel := viewmodel.SetupViewmodel(client, _controller)
	_handler := handler.SetupHandler(client, &_viewmodel)
	handlers := topic.MAPSubTopic2MessageHandler{
		topic.IRGoInDirection: {QoS: 2, Handler: _handler.IRSensorGoInHandler()},
		topic.RFIDSubTop:      {QoS: 2, Handler: _handler.RFID()},
	}
	topic.ClientSubTopics(client, handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go _viewmodel.HandleStatuses()

	<-quit
	log.Info().Msg("Shutdown client")
}
