package main

import (
	"bufio"
	"fmt"
	mqtt_client2 "github.com/thelemonday/smart-parking-iot-server/internal/mqtt_client"
	"github.com/thelemonday/smart-parking-iot-server/pkg/util"
	"os"

	"github.com/rs/zerolog/log"
)

const optionsStmt = `
1: Publish car go in detected
2: Publish car go in detected (no)

3: Publish car go out detected
4: Publish car go out detected (no)

5: Publish RFID tag detected
Choice >`

func main() {
	client := mqtt_client2.SetupMQTTClient(TestClientConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic().Err(token.Error()).Msg("")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Println(optionsStmt); scanner.Scan(); fmt.Print("Choice >") {
		switch scanner.Text() {
		case "1":
			client.Publish(mqtt_client2.IRGoInDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": true,
			}))
		case "2":
			client.Publish(mqtt_client2.IRGoInDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": false,
			}))
		case "3":
			client.Publish(mqtt_client2.IRGoOutDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": true,
			}))
		case "4":
			client.Publish(mqtt_client2.IRGoOutDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": false,
			}))
		case "5":
			client.Publish(mqtt_client2.RFIDSubTop, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"uid": "B3AD9715",
			}))
		}
	}
}
