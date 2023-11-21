package util

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

func MarshalJsonData2Byte(data interface{}) []byte {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return payload
}

func UnmarshalByte2JsonData[D any](payload []byte) (*D, error) {
	var jsonData D
	if err := json.Unmarshal(payload, &jsonData); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	return &jsonData, nil
}

func TokenWaitAndLog(token mqtt.Token) {
	if token.Wait(); token.Error() != nil {
		log.Error().Err(token.Error()).Msg("")
	}
}

func GenerateNewNanoID(length int) string {
	id, err := gonanoid.New(length)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return id
}
