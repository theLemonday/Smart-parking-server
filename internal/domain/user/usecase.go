package user

import (
	"github.com/rs/zerolog/log"
	"math"
	"time"
)

type UseCase interface {
	NewUserEnter(id string) *User
	UserLeave(id string)
	CalculateCarParkBill(id string) *PaymentBill
}

type userUseCase struct {
	carParkStatusDAO UserDao
}

func (s userUseCase) NewUserEnter(id string) *User {
	if !s.carParkStatusDAO.IsRFIDTagValid(id) {
		log.Error().Msgf("invalid rfid tag: %s", id)
		return nil
	}

	user := s.carParkStatusDAO.Create(id)
	log.Info().Msgf("new user enter: %+v\n", user)
	return user
}

func (s userUseCase) UserLeave(id string) {
	if !s.carParkStatusDAO.IsRFIDTagValid(id) {
		log.Error().Msgf("invalid rfid tag: %s", id)
	}

	log.Info().Msgf("delete user %s", id)
	s.carParkStatusDAO.Delete(id)
}

func (s userUseCase) CalculateCarParkBill(id string) *PaymentBill {
	user := s.carParkStatusDAO.Get(id)

	goOutTimestamp := time.Now()
	duration := goOutTimestamp.Sub(user.GoInTimestamp).Seconds()

	return &PaymentBill{
		RFIDTag:        id,
		TotalCost:      int(math.Round(duration) + 0.5 + 10),
		GoInTimestamp:  user.GoInTimestamp,
		GoOutTimestamp: goOutTimestamp,
	}
}

func NewUserUseCase(dao UserDao) UseCase {
	return &userUseCase{
		carParkStatusDAO: dao,
	}
}
