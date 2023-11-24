package domain

type CarParkStatusDAO interface {
	NewUserIdentifiedByRFID(RFIDTag string) *User
	CalculateCost(id string) *PaymentBill
	OnUserLeave(id string)
}
