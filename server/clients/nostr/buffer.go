package nostr

import (
	"sync"

	"github.com/nbd-wtf/go-nostr"
)

// Size of the buffer
const BUFFER_SIZE = 1000

// A thread-safe ringbuffer based container of events
type eventBuffer struct {
	idx      uint
	msgCount int //really a uint but int for compatibility with different code
	buf      [BUFFER_SIZE]*nostr.Event

	mutex sync.RWMutex
}

// create a new event ring buffer
func newEventBuffer() *eventBuffer {
	return &eventBuffer{
		idx:      0,
		msgCount: 0,
		buf:      [BUFFER_SIZE]*nostr.Event{},
	}
}

// push an event to the ring buffer
func (eb *eventBuffer) push(event *nostr.Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	var score int
	if eb.buf[eb.idx] != nil {
		score = -1
	}
	if event != nil {
		score++
	}
	eb.buf[eb.idx] = event
	eb.idx = (eb.idx + 1) % BUFFER_SIZE
	eb.msgCount += score
}

// Get a slice with the available messages
func (eb *eventBuffer) slice() []nostr.Event {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	s := make([]nostr.Event, eb.msgCount)
	var i uint
	for i < uint(eb.msgCount) {
		if eb.buf[i+eb.idx] != nil {
			s[i] = *eb.buf[i+eb.idx]
			i++
		}
	}
	return s
}

// Get a slice with the available messages, and clear it
func (eb *eventBuffer) take() []nostr.Event {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	s := make([]nostr.Event, eb.msgCount)
	var i, j uint
	for i < uint(eb.msgCount) {
		if eb.buf[(i+eb.idx)%BUFFER_SIZE] != nil {
			s[j] = *eb.buf[i+eb.idx]
			j++
		}
		i++
	}

	eb.idx = 0
	eb.msgCount = 0
	eb.buf = [BUFFER_SIZE]*nostr.Event{}

	return s

}
