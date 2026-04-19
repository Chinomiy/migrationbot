package fsm

import (
	"fmt"
	"migtationbot/logger"
) // главный экран -> список стран -> выбор страны -> выбор типа поездки -> информация о поездке

const NoChange StateID = "no_change"

type StateID = string
type State struct {
	ID      StateID
	LastMsg int
	Data    []any
}
type FSM struct {
	initialStateID StateID
	userStates     UserStateStorage
}

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
	lstMsg int,
	args ...any,
) error {
	if stateID == NoChange {
		return nil
	}
	state := State{ID: stateID, Data: args, LastMsg: lstMsg}
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
		initial := State{ID: f.initialStateID, Data: nil, LastMsg: 0}
		if err := f.userStates.Push(userID, initial); err != nil {
			return State{}, err
		}
		return initial, nil
	}
	return f.userStates.Get(userID)
}
func (f *FSM) Reset(userID int64) error {
	initial := State{ID: f.initialStateID, Data: nil, LastMsg: 0}
	fmt.Println(f.userStates.Reset(userID, initial))
	return f.userStates.Reset(userID, initial)
}
func (f *FSM) Back(userID int64) error {
	_, err := f.userStates.Back(userID)
	if err != nil {
		return err
	}
	return nil
}
