package fsm

import (
	"fmt"
	"sync"
)

type UserStateStorage interface {
	Exists(userID int64) (bool, error)
	Get(userID int64) (State, error)
	Push(userID int64,
		stateID State) error
	Reset(userID int64, initial State) error
	Back(userID int64) (State, error)
}
type userStateStorage struct {
	mu      sync.RWMutex
	Storage map[int64][]State
}

func initialUserStateStorage() *userStateStorage {
	return &userStateStorage{Storage: make(map[int64][]State)}
}
func (u *userStateStorage) Push(userID int64, state State) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Storage[userID] = append(u.Storage[userID], state)
	return nil
}
func (u *userStateStorage) Get(userID int64) (State, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	stack, ok := u.Storage[userID]
	if !ok || len(stack) == 0 {
		return State{}, fmt.Errorf("no state for userID: %d", userID)
	}
	return stack[len(stack)-1], nil
}
func (u *userStateStorage) Back(userID int64) (State, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	stack, ok := u.Storage[userID]
	if !ok || len(stack) == 0 {
		return State{}, fmt.Errorf("no state for userID: %d", userID)
	}
	if len(stack) == 1 {
		return stack[0], nil
	}
	u.Storage[userID] = stack[:len(stack)-1]
	return u.Storage[userID][len(u.Storage[userID])-1], nil
}
func (u *userStateStorage) Exists(userID int64) (bool, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	
	stack, ok := u.Storage[userID]
	return ok && len(stack) > 0, nil
}
func (u *userStateStorage) Reset(userID int64, initial State) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Storage[userID] = []State{initial}
	return nil
}
