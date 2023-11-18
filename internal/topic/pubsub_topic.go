package topic

type LEDTopic = string
type IRSensorTopic = string

const (
	OLEDPubTop                      = "smart-parking/gate/OLED"
	GateBarrierPubTop               = "smart-parking/gate/barrier"
	RFIDSubTop                      = "smart-parking/gate/RFID/out"
	IRTopicPrefix                   = "smart-parking/gate/IR/+"
	IRGoInDirection   IRSensorTopic = "smart-parking/gate/IR/in"
	IRGoOutDirection  IRSensorTopic = "smart-parking/gate/IR/out"
	GreenLEDPubTop    LEDTopic      = "smart-parking/gate/greenLED"
	RedLEDPubTop      LEDTopic      = "smart-parking/gate/redLED"
)
