package database

import "time"

// CarParkStatusDAO If identified by RFID, id = RFID tag
// If identified by QR, id = username
type CarParkStatusDAO interface {
	NewUserIdentifiedByRFID(RFIDTag string) User
	NewUserIdentifiedByQRCode(username string) User
	GetGoInTimestampOfUser(id string) (time.Time, error)
	IsRFIDTagValid(uid string) bool
	GetAllUsers() []User
	DeleteUser(id string)
	CalculateCost(id string) *PaymentBill
}

type AccountsDAO interface {
	AuthenticateUser(username, password string) (Account, error)
	GetAccountProfile(username string) (Account, error)
	UserCreditAccount(username string, creditMoney int) error
	// ParkingCostBill(username string) (*PaymentBill, error)
}
