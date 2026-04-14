package fsm

import (
	"context"
	"migtationbot/logger"
) // главный экран -> список стран -> выбор страны -> выбор типа поездки -> информация о поездке

type StateID string
type State struct {
	ID   StateID
	Data []any
}
type Callback func(ctx context.Context, args ...any) error
type FSM struct {
	initialStateID StateID
	callbacks      map[StateID]Callback
	userStates     UserStateStorage
}

func New(initialStateName StateID,
	callbacks map[StateID]Callback,
) *FSM {
	s := &FSM{
		initialStateID: initialStateName,
		callbacks:      make(map[StateID]Callback),
		userStates:     initialUserStateStorage(),
	}
	for stateID, callback := range callbacks {
		s.callbacks[stateID] = callback
	}
	return s
}
func (f *FSM) Transition(
	ctx context.Context,
	userID int64,
	stateID StateID,
	args ...any,
) error {
	state := State{ID: stateID, Data: args}
	if err := f.userStates.Push(userID, state); err != nil {
		// добавить логи чувствую есть подводные
		return err
	}
	logger.Infof("user %d transitioned to state %s", userID, stateID)
	if cb, ok := f.callbacks[stateID]; ok {
		return cb(ctx, args...)
	}
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
func (f *FSM) Back(ctx context.Context, userID int64) error {
	state, err := f.userStates.Back(userID)
	if err != nil {
		return err
	}
	if cb, ok := f.callbacks[state.ID]; ok {
		return cb(ctx, state.Data...)
	}
	return nil
}
