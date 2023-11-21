package state_manager

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/topic"
	"strings"
)

func (v *StateManager) OnCarGoInDetected(detected bool) {
	v._IRSensorIn <- detected
}

func (v *StateManager) OnCarGoOutDetected(detected bool) {
	v._IRSensorOut <- detected
}

func numberOfOccupiedSlots(slotsStatus map[string]bool) int {
	slotsOccupied := 0
	for _, v := range slotsStatus {
		if v {
			slotsOccupied++
		}
	}

	return slotsOccupied
}

func (v *StateManager) OnCarGoInSlotDetected(sensorTopic string, detected bool) {
	slotId := strings.TrimLeft(sensorTopic, topic.IRSlotPrefix)
	v.slotsStatus[slotId] = detected
	v.websocketService.OnSlotStatusChanged(slotId, detected)
}

func (v *StateManager) OnRFIDTagRead(uid string) {
	v._RFIDUid <- uid
}

func (v *StateManager) OnQRCodeScanned(QRCode, username string) {
	v._QRCodeScanner <- [2]string{QRCode, username}
}
