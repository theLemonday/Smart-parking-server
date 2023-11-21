package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/thelemonday/smart-parking-iot-server/internal/server"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

func main() {
	carParkingStatusDatabase := db.NewCurrentCarParkingStatusDatabase()
	accountsDatabase := db.NewAccountsDatabase(carParkingStatusDatabase)
	_server := server.NewSmartParkingIotServer(accountsDatabase)
	go _server.ListenAndServe(serverExposedPort)

	//client := mqtt_client.SetupMQTTClient(clientConfig)
	//if token := client.Connect(); token.Wait() && token.Error() != nil {
	//	log.Panic().Err(token.Error()).Msg("")
	//}
	//
	//_controller := controller.SetupController(client)
	//_viewmodel := state-manager.SetupViewmodel(client, _controller, carParkingStatusDatabase)
	//_handler := handler.SetupHandler(client, &_viewmodel)
	//handlers := mqtt_client.MAPSubTopic2MessageHandler{
	//	topic.IRGoInDirection: {QoS: 2, Handler: _handler.IRSensorGoInHandler()},
	//	topic.RFIDSubTop:      {QoS: 2, Handler: _handler.RFID()},
	//}
	//mqtt_client.ClientSubTopics(client, handlers)
	//
	//
	//go _viewmodel.HandleStatuses()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutdown client")
}
