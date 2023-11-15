package db

import "time"

type DatabaseRepo interface {
	InsertNewCurrentUser(id string, identifiedBy string)
	GetGoInTimestampOfUser(id string) (time.Time, error)
	DeleteUser(id string)
	GetAllUsers() []User
	IsRFIDTagValid(uid string) bool
}
