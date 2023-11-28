package state_manager

import "github.com/rs/zerolog/log"

type IdentifyStatus int

const (
	unknown IdentifyStatus = iota
	waitingToBeIdentified
	identified
)

type state struct {
	carGoIn               bool
	carGoOut              bool
	newUserIdentifyStatus IdentifyStatus
	identificationID      string
	isGoIn                bool
	slotsStatus           map[string]bool
}

func (s *state) reset() {
	log.Info().Str("service", "state-manager").Msg("reset state")
	s.carGoIn = false
	s.carGoOut = false
	s.newUserIdentifyStatus = unknown
	s.identificationID = ""
	s.isGoIn = false
}

func newCarParkState() *state {
	s := new(state)
	s.reset()

	return s
}
