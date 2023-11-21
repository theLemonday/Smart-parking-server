package state_manager

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
}

func resetState(_state *state) {
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
