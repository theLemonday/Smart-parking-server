package internal

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/internal/domain/user"
	"time"
)

var (
	uidTags = []string{
		"B3AD9715",
		"1365B515",
		"B3726B0F",
		"B3F1CO15",
		"536FBF15",
		"73C4B215",
	}
)

type UsersDatabase struct {
	currentUsers map[string]*user.User
}

func NewUsersDatabase() *UsersDatabase {
	return &UsersDatabase{
		currentUsers: make(map[string]*user.User),
	}
}

func (d *UsersDatabase) Create(RFIDTag string) *user.User {
	log.Info().Msgf("create new user uid: %s", RFIDTag)
	newUser := user.User{Id: RFIDTag, GoInTimestamp: time.Now()}
	d.currentUsers[RFIDTag] = &newUser

	return &newUser
}

func (d *UsersDatabase) IsRFIDTagValid(uid string) bool {
	for _, v := range uidTags {
		if v == uid {
			return true
		}
	}

	return false
}

func (d *UsersDatabase) Get(RFIDTag string) *user.User {
	return d.currentUsers[RFIDTag]
}

func (d *UsersDatabase) Delete(id string) {
	delete(d.currentUsers, id)
}
