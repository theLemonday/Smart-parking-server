package db

import "time"

// CarParkStatusDatabaseRepo If identified by RFID, id = RFID tag
// If identified by QR, id = username
type CarParkStatusDatabaseRepo interface {
	NewUserIdentifiedByRFID(RFIDTag string)
	GetGoInTimestampOfUser(id string) (time.Time, error)
	IsRFIDTagValid(uid string) bool
	GetAllUsers() []User
	DeleteUser(id string)
	CalculateCost(id string) int
}

type AccountsDatabaseRepo interface {
	AuthenticateUser(username, password string) (Account, error)
	GetAccountProfile(username string) (Account, error)
	UserCreditAccount(username string, creditMoney int) error
	UserPayParkingCost(username string) (bool, error)
}
