package domain

import "time"

type User struct {
	Id             string
	GoInTimestamp  time.Time
	GoOutTimestamp time.Time
}

type PaymentBill struct {
	RFIDTag        string
	GoInTimestamp  time.Time
	GoOutTimestamp time.Time
	TotalCost      int
}
