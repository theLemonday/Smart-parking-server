package db

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
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

type currentCarParkingStatusDatabase struct {
	currentUsers map[string]User
}

func NewCurrentCarParkingStatusDatabase() *currentCarParkingStatusDatabase {
	return &currentCarParkingStatusDatabase{
		currentUsers: map[string]User{},
	}
}

func (d *currentCarParkingStatusDatabase) GetAllUsers() []User {
	var users = make([]User, 0, len(d.currentUsers))
	for _, v := range d.currentUsers {
		users = append(users, v)
	}

	return users
}

func (d *currentCarParkingStatusDatabase) AuthenticateUser(username, password string) (*Account, bool) {

}

func (d *currentCarParkingStatusDatabase) IsRFIDTagValid(uid string) bool {
	for _, v := range uidTags {
		if v == uid {
			return true
		}
	}

	return false
}

func (d *currentCarParkingStatusDatabase) InsertNewCurrentUser(id string, identifiedBy string) {
	log.Info().Msgf("Insert new user into database")
	d.currentUsers[id] = User{id: id, identifiedBy: identifiedBy, goInTimestamp: time.Now()}
}

func (d *currentCarParkingStatusDatabase) GetGoInTimestampOfUser(id string) (time.Time, error) {
	if v, ok := d.currentUsers[id]; ok {
		return v.goInTimestamp, nil
	}

	return time.Time{}, errors.New("no user with given Id")
}

func (d *currentCarParkingStatusDatabase) DeleteUser(id string) {
	delete(d.currentUsers, id)
}
