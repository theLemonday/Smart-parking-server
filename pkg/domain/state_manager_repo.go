package domain

type StateManagerRepository interface {
	OnCarGoInDetected(detected bool)
	OnCarGoOutDetected(detected bool)
	OnCarGoInSlotDetected(sensorTopic string, detected bool)
	OnRFIDTagRead(uid string)
	OnQRCodeScanned(QRCode, username string)
}
