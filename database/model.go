package database

import "time"

type Account struct {
	Username string `json:"username"`
	password string
	Balance  int `json:"balance"`
}

type User struct {
	Id             string
	IdentifiedBy   string
	GoInTimestamp  time.Time
	GoOutTimestamp time.Time
}

type PaymentBill struct {
	Username       string
	GoInTimestamp  time.Time
	GoOutTimestamp time.Time
	TotalCost      int
}
