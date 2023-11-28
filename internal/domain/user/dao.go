package user

type UserDao interface {
	Create(RFIDTag string) *User
	Get(RFIDTag string) *User
	Delete(id string)
	IsRFIDTagValid(tag string) bool
}
