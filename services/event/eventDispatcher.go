package event

import (
	"sync"

	"github.com/yahyaammar-dev/pacebe/types"
)

var (
	// Global instance
	globalDispatcher *EventDispatcher
	once             sync.Once
)

type EventDispatcher struct {
	listeners map[string][]types.Listener
	mu        sync.Mutex
}

// GetDispatcher returns the singleton instance of EventDispatcher
func GetDispatcher() *EventDispatcher {
	once.Do(func() {
		globalDispatcher = &EventDispatcher{
			listeners: make(map[string][]types.Listener),
		}
	})
	return globalDispatcher
}

// Register adds a listener for an event
func Register(eventName string, listener types.Listener) {
	dispatcher := GetDispatcher()
	dispatcher.mu.Lock()
	defer dispatcher.mu.Unlock()
	dispatcher.listeners[eventName] = append(dispatcher.listeners[eventName], listener)
}

// Dispatch triggers an event
func Dispatch(event types.Event) {
	dispatcher := GetDispatcher()
	dispatcher.mu.Lock()
	defer dispatcher.mu.Unlock()
	for _, listener := range dispatcher.listeners[event.Name] {
		listener(event)
	}
}
