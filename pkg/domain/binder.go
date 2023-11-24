package domain

type StateManager2Websocket interface {
	OnNewUserEnter(user *User)
	OnSlotStatusChanged(slotId string, detected bool)
	OnUserLeave(uid string)
	OnUserGoOutIdentified(bill *PaymentBill)
}

type Websocket2StateManager interface {
	//OnUserDonePayment(username string)
	OnUserIdentifiedByRFIDDonePayment(uid string)
}
