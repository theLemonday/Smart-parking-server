package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/thelemonday/smart-parking-iot-server/controller"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client/handler"
	state_manager "github.com/thelemonday/smart-parking-iot-server/internal/state-manager"
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"

	"github.com/thelemonday/smart-parking-iot-server/internal/server"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/database"
)

func main() {
	carParkingStatusDatabase := database.NewCurrentCarParkingStatusDatabase()
	accountsDatabase := database.NewAccountsDatabase(carParkingStatusDatabase)
	_server := server.NewSmartParkingIotServer(accountsDatabase)
	go _server.ListenAndServe(serverExposedPort)

	client := mqtt_client.SetupMQTTClient(clientConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic().Err(token.Error()).Msg("")
	}

	// _controller := controller.NewController(client)
	_controller := controller.NewTestController()
	stateManager := state_manager.NewStateManager(client, _controller, carParkingStatusDatabase, _server)
	_handler := handler.SetupHandler(client, stateManager)
	handlers := mqtt_client.MAPSubTopic2MessageHandler{
		topic.IRSensorSubTop: {QoS: 2, Handler: _handler.CarSensorDetectedHandler()},
		topic.RFIDSubTop:     {QoS: 2, Handler: _handler.RFID()},
	}
	mqtt_client.ClientSubTopics(client, handlers)

	go stateManager.HandleStatuses()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutdown client")
}
