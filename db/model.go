package db

import "time"

type Account struct {
	Username string `json:"username"`
	password string
	Balance  int `json:"balance"`
}

type User struct {
	id             string
	identifiedBy   string
	goInTimestamp  time.Time
	goOutTimestamp time.Time
}
