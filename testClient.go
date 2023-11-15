package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/topic"
	"github.com/thelemonday/smart-parking-iot-server/util"
)

const optionsStmt = `
1: Publish car go in detected
2: Publish car go in detected (no)

3: Publish car go out detected
4: Publish car go out detected (no)

5: Publish RFID tag detected
Choice >
`

func main() {
	client := SetupMQTTClient(TestClientConfig)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic().Err(token.Error()).Msg("")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Println(optionsStmt); scanner.Scan(); fmt.Print("Choice >") {
		switch scanner.Text() {
		case "1":
			client.Publish(topic.IRGoInDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": true,
			}))
		case "2":
			client.Publish(topic.IRGoInDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": false,
			}))
		case "3":
			client.Publish(topic.IRGoOutDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": true,
			}))
		case "4":
			client.Publish(topic.IRGoOutDirection, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"detected": false,
			}))
		case "5":
			client.Publish(topic.RFIDSubTop, 2, false, util.MarshalJsonData2Byte(map[string]any{
				"uid": "B3AD9715",
			}))
		}
	}
}
