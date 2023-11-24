package internal

import (
	"github.com/rs/zerolog/log"
	"github.com/thelemonday/smart-parking-iot-server/pkg/domain"
	"math"
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

type CarParkingStatusDatabase struct {
	currentUsers map[string]*domain.User
}

func NewCarParkingStatusDatabase() *CarParkingStatusDatabase {
	return &CarParkingStatusDatabase{
		currentUsers: make(map[string]*domain.User),
	}
}

func (d *CarParkingStatusDatabase) NewUserIdentifiedByRFID(RFIDTag string) *domain.User {
	if !d.isRFIDTagValid(RFIDTag) {
		return nil
	}

	log.Info().Msgf("new user identified by RFID tag: %s", RFIDTag)
	newUser := domain.User{Id: RFIDTag, GoInTimestamp: time.Now()}
	d.currentUsers[RFIDTag] = &newUser

	return &newUser
}

func (d *CarParkingStatusDatabase) isRFIDTagValid(uid string) bool {
	for _, v := range uidTags {
		if v == uid {
			return true
		}
	}

	log.Error().Msg("RFID tag is not valid")
	return false
}

func (d *CarParkingStatusDatabase) CalculateCost(id string) *domain.PaymentBill {
	user := d.currentUsers[id]

	goOutTimestamp := time.Now()
	duration := goOutTimestamp.Sub(user.GoInTimestamp).Seconds()

	return &domain.PaymentBill{
		RFIDTag:        id,
		TotalCost:      int(math.Round(duration) + 0.5 + 10),
		GoInTimestamp:  user.GoInTimestamp,
		GoOutTimestamp: goOutTimestamp,
	}
}

func (d *CarParkingStatusDatabase) OnUserLeave(id string) {
	delete(d.currentUsers, id)
}
