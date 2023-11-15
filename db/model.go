package db

import "time"

type account struct {
	phoneNumber string
	password    string
	balance     uint
}

type User struct {
	id             string
	identifiedBy   string
	goInTimestamp  time.Time
	goOutTimestamp time.Time
}
