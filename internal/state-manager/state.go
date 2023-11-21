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

func resetState(_state *state) {
	log.Info().Msg("state-manager reset")
	_state.carGoIn = false
	_state.carGoOut = false
	_state.newUserIdentifyStatus = unknown
	_state.identificationID = ""
	_state.isGoIn = false
}

func newState() *state {
	s := new(state)
	resetState(s)

	return s
}
