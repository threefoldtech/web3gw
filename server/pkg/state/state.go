package state

import (
	"fmt"
	"sync"
	"time"
)

const (
	// Remove stale keys every minute
	keyCleanseInterval = time.Second * 60
	// A key is considered stale after 5 minutes
	keyStaleMark = time.Second * 300
)

type State interface{}

// state and metadata to manage it
type stateMeta[S State] struct {
	// timestamp when the state was last accessed
	accessed int64
	state    S
}

type StateManager[S State] struct {
	cleanseTicker *time.Ticker
	closeChan     chan struct{}
	conStates     sync.Map
}

// Set state for a connection ID. Previously saved state will be overriden.
func (sm *StateManager[S]) Set(conID string, state S) {
	sm.conStates.Store(conID, stateMeta[S]{accessed: time.Now().Unix(), state: state})
}

// Get a previously saved state for a connection ID. If the connection ID is not present in the map, a blank state will be returned
func (sm *StateManager[S]) Get(conID string) (S, bool) {
	val, exists := sm.conStates.Load(conID)
	var r stateMeta[S]
	if exists {
		// Conversion always succeeds as this is the only value we ever set
		r, _ = val.(stateMeta[S])
		//r = *t
		// Update and replace the metadata of the state
		r.accessed = time.Now().Unix()
		sm.conStates.Store(conID, r)

	}
	return r.state, exists
}

// Close the StateManager, cleaning up the background worker. It should not be used after this call.
func (sm *StateManager[S]) Close() {
	sm.cleanseTicker.Stop()
	close(sm.closeChan)
}

func NewStateManager[S State]() *StateManager[S] {
	sm := &StateManager[S]{
		cleanseTicker: time.NewTicker(keyCleanseInterval),
		closeChan:     make(chan struct{}, 1),
		conStates:     sync.Map{},
	}

	go func() {
		alive := true
		for alive {
			select {
			case _, open := <-sm.cleanseTicker.C:
				alive = open
				if !open {
					break
				}
				fmt.Println("Checking keys")
				//cleanse keys
				sm.conStates.Range(func(key any, value any) bool {
					meta, ok := value.(stateMeta[S])
					if !ok {
						fmt.Println("invalid state meta conversion in cleanup loop")
						return true
					}
					if meta.accessed < time.Now().Unix()-int64(keyStaleMark.Seconds()) {
						fmt.Println("Removing stale key", key)
						sm.conStates.Delete(key)
					}
					return true
				})
			case <-sm.closeChan:
				alive = false
				break
			}
		}
		fmt.Println("State manager background task closed")
	}()

	return sm
}
