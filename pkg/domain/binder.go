package domain

import "github.com/thelemonday/smart-parking-iot-server/database"

type StateManager2Websocket interface {
	OnNewUserEnter(user database.User)
	OnSlotStatusChanged(slotId string, detected bool)
	OnUserLeave(uid string)
	OnUserGoOutIdentified(bill *database.PaymentBill)
}

type Websocket2StateManager interface {
	//OnUserDonePayment(username string)
	OnUserIdentifiedByRFIDDonePayment(uid string)
}
