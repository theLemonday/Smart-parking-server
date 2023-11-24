package mqtt_client

type LEDTopic = string
type IRSensorTopic = string

const (
	OLEDPubTop                      = "smart-parking/gate/OLED"
	GateBarrierPubTop               = "smart-parking/gate/barrier"
	RFIDSubTop                      = "smart-parking/gate/RFID/out"
	IRSensorSubTop                  = "smart-parking/gate/IR/#"
	IRSensorPrefix                  = "smart-parking/gate/IR/"
	IRSlotPrefix                    = "smart-parking/gate/IR/slot/"
	IRGoInDirection   IRSensorTopic = "smart-parking/gate/IR/in"
	IRGoOutDirection  IRSensorTopic = "smart-parking/gate/IR/out"
	GreenLEDPubTop    LEDTopic      = "smart-parking/gate/greenLED"
	RedLEDPubTop      LEDTopic      = "smart-parking/gate/redLED"
)
