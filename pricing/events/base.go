package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"event_id"`
	Type      string    `json:"event_type"`
	Timestamp string    `json:"event_timestamp"`
	Version   string    `json:"version"`
}

func newBaseEvent(eventType string, version string) Event {
	return Event{
		ID:        uuid.New(),
		Type:      eventType,
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Version:   version,
	}
}

func (e *Event) EventID() uuid.UUID {
	return e.ID
}
func (e *Event) EventType() string {
	return e.Type
}
func (e *Event) EventVersion() string {
	return e.Version

}
func (e *Event) EventTimeStamp() string {
	return e.Timestamp
}
