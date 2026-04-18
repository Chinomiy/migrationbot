package fsm

import (
	"migtationbot/logger"
) // главный экран -> список стран -> выбор страны -> выбор типа поездки -> информация о поездке

const NoChange StateID = "no_change"

type StateID = string
type State struct {
	ID   StateID
	Data []any
}
type FSM struct {
	initialStateID StateID
	userStates     UserStateStorage
}

// map[string]mao[string]string
func New(initialStateName StateID,
) *FSM {
	s := &FSM{
		initialStateID: initialStateName,
		userStates:     initialUserStateStorage(),
	}
	return s
}
func (f *FSM) Transition(
	userID int64,
	stateID StateID,
	args ...any,
) error {
	if stateID == NoChange {
		return nil
	}
	state := State{ID: stateID, Data: args}
	if err := f.userStates.Push(userID, state); err != nil {
		return err
	}
	logger.Infof("user %d transitioned to state %s", userID, stateID)
	return nil
}
func (f *FSM) Current(userID int64) (State, error) {
	ok, err := f.userStates.Exists(userID)
	if err != nil {
		return State{}, err
	}
	if !ok {
		initial := State{ID: f.initialStateID, Data: nil}
		if err := f.userStates.Push(userID, initial); err != nil {
			return State{}, err
		}
		return initial, nil
	}
	return f.userStates.Get(userID)
}
func (f *FSM) Reset(userID int64) error {
	initial := State{ID: f.initialStateID, Data: nil}
	return f.userStates.Reset(userID, initial)
}
func (f *FSM) Back(userID int64) error {
	_, err := f.userStates.Back(userID)
	if err != nil {
		return err
	}
	return nil
}
