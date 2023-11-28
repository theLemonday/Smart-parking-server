package domain

import (
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
)

type StateManager2Websocket interface {
	OnNewUserEnter(user *user.User)
	OnSlotStatusChanged(slotId string, detected bool)
	OnUserLeave(uid string)
	OnUserGoOutIdentified(bill *user.PaymentBill)
}

type Websocket2StateManager interface {
	OnUserIdentifiedByRFIDDonePayment(uid string)
}
