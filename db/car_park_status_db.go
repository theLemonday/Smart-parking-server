package db

import (
	"errors"
	"math"
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

type CurrentCarParkingStatusDatabase struct {
	currentUsers map[string]User
}

func (d *CurrentCarParkingStatusDatabase) CalculateCost(id string) int {
	user := d.currentUsers[id]

	duration := time.Now().Sub(user.goInTimestamp).Seconds()

	return int(math.Round(duration) + 0.5 + 10)
}

func NewCurrentCarParkingStatusDatabase() *CurrentCarParkingStatusDatabase {
	return &CurrentCarParkingStatusDatabase{
		currentUsers: make(map[string]User),
	}
}

func (d *CurrentCarParkingStatusDatabase) GetAllUsers() []User {
	var users = make([]User, 0, len(d.currentUsers))
	for _, v := range d.currentUsers {
		users = append(users, v)
	}

	return users
}

func (d *CurrentCarParkingStatusDatabase) IsRFIDTagValid(uid string) bool {
	for _, v := range uidTags {
		if v == uid {
			return true
		}
	}

	return false
}

func (d *CurrentCarParkingStatusDatabase) NewUserIdentifiedByRFID(RFIDTag string) {
	log.Info().Msgf("New user identified by RFID tag: %s", RFIDTag)
	d.currentUsers[RFIDTag] = User{id: RFIDTag, identifiedBy: "rfid", goInTimestamp: time.Now()}
}

func (d *CurrentCarParkingStatusDatabase) NewUserIdentifiedByQRCode(username string) {
	log.Info().Msgf("New user identified by QRCode: %s", username)
	d.currentUsers[username] = User{id: username, identifiedBy: "qr", goInTimestamp: time.Now()}
}

func (d *CurrentCarParkingStatusDatabase) GetGoInTimestampOfUser(id string) (time.Time, error) {
	if v, ok := d.currentUsers[id]; ok {
		return v.goInTimestamp, nil
	}

	return time.Time{}, errors.New("no user with given Id")
}

func (d *CurrentCarParkingStatusDatabase) DeleteUser(id string) {
	delete(d.currentUsers, id)
}
