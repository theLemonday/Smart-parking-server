package main

import (
	"github.com/thelemonday/smart-parking-iot-server/internal"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client/controller"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client/handler"
	"github.com/thelemonday/smart-parking-iot-server/internal/state-manager"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/server"
)

func main() {
	_server := server.NewSmartParkingIotServer()
	go _server.ListenAndServe(serverExposedPort)

	iotGateway := internal.NewIotGateway(clientConfig)
	iotGateway.Connect()

	_controller := controller.NewTestController()
	carParkingStatusDatabase := internal.NewCarParkingStatusDatabase()
	stateManager := state_manager.NewStateManager(iotGateway.GetMQTTClient(), _controller, carParkingStatusDatabase)
	stateManager.SetWebsocketService(_server)
	_handler := handler.SetupHandler(iotGateway.GetMQTTClient(), stateManager)
	handlers := internal.MAPSubTopic2MessageHandler{
		mqtt_client.IRSensorSubTop: {QoS: 2, Handler: _handler.CarSensorDetectedHandler()},
		mqtt_client.RFIDSubTop:     {QoS: 2, Handler: _handler.RFID()},
	}
	iotGateway.SubscribeTopics(handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutdown client")
}
