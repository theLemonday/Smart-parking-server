package secret

type MQTTConfig struct {
	Protocol string
	Broker   string
	Port     int
	ClientId string
	Username string
	Password string
}
