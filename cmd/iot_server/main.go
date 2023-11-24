package main

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client/controller"
	"github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/mqtt_client/handler"
	state_manager "github.com/thelemonday/smart-parking-iot-server/internal/iot_gateway/state-manager"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/database"
	"github.com/thelemonday/smart-parking-iot-server/internal/server"
)

func main() {
	carParkingStatusDatabase := database.NewCurrentCarParkingStatusDatabase()
	accountsDatabase := database.NewAccountsDatabase(carParkingStatusDatabase)
	_server := server.NewSmartParkingIotServer(accountsDatabase)
	go _server.ListenAndServe(serverExposedPort)

	iotGateway := iot_gateway.NewIotGateway(clientConfig)

	_controller := controller.NewTestController()
	stateManager := state_manager.NewStateManager(iotGateway.GetMQTTClient(), _controller, carParkingStatusDatabase)
	stateManager.SetWebsocketService(_server)
	_handler := handler.SetupHandler(iotGateway.GetMQTTClient(), stateManager)
	handlers := iot_gateway.MAPSubTopic2MessageHandler{
		mqtt_client.IRSensorSubTop: {QoS: 2, Handler: _handler.CarSensorDetectedHandler()},
		mqtt_client.RFIDSubTop:     {QoS: 2, Handler: _handler.RFID()},
	}
	iotGateway.SubscribeTopics(handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutdown client")
}
