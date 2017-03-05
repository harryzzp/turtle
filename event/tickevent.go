package event

import (
	"sync"
	"time"
)

type TickEvent struct {
	tickTime  time.Time
	eventName string
	event     string
}

type TickEventChannel struct {
	C    chan TickEvent
	once sync.Once
}

func NewTickEvent() *TickEventChannel {
	return &TickEventChannel{C: make(chan TickEvent)}
}

func (tc *TickEventChannel) SafeClose() {
	tc.once.Do(func() {
		close(tc.C)
	})
}
