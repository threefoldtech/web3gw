package nostr

import (
	"sync"

	"github.com/nbd-wtf/go-nostr"
)

// Size of the buffer
const BUFFER_SIZE = 1000

// A thread-safe ringbuffer based container of events
type eventBuffer struct {
	idx uint
	buf [BUFFER_SIZE]*nostr.Event

	mutex sync.RWMutex
}

// create a new event ring buffer
func newEventBuffer() *eventBuffer {
	return &eventBuffer{
		idx: 0,
		buf: [BUFFER_SIZE]*nostr.Event{},
	}
}

// push an event to the ring buffer
func (eb *eventBuffer) push(event *nostr.Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	eb.buf[eb.idx] = event
	eb.idx = (eb.idx + 1) % BUFFER_SIZE
}

// Get a slice with the available messages
func (eb *eventBuffer) slice() []nostr.Event {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	s := make([]nostr.Event, 0, BUFFER_SIZE)
	for i := eb.idx; i < eb.idx+BUFFER_SIZE; i++ {
		if eb.buf[i%BUFFER_SIZE] != nil {
			s = append(s, *eb.buf[i%BUFFER_SIZE])
		}
	}

	return s
}

// Get a slice with the available messages, and clear it
func (eb *eventBuffer) take() []nostr.Event {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	s := make([]nostr.Event, 0, BUFFER_SIZE)
	for i := eb.idx; i < eb.idx+BUFFER_SIZE; i++ {
		if eb.buf[i%BUFFER_SIZE] != nil {
			s = append(s, *eb.buf[i%BUFFER_SIZE])
		}
	}

	eb.idx = 0
	eb.buf = [BUFFER_SIZE]*nostr.Event{}

	return s
}

// Consume and return `cnt` events
func (eb *eventBuffer) consume(cnt uint32) []nostr.Event {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	var s []nostr.Event
	for i := eb.idx; i < eb.idx+BUFFER_SIZE && cnt > 0; i++ {
		if eb.buf[i%BUFFER_SIZE] != nil {
			s = append(s, *eb.buf[i%BUFFER_SIZE])
			eb.buf[i%BUFFER_SIZE] = nil
			cnt--
		}
	}

	return s
}
