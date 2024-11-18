package enum

import "fmt"

type EventType string

const (
	Game     EventType = "game"
	Practice EventType = "practice"
	Misc     EventType = "misc"
)

func (et EventType) String() string {
	return string(et)
}

func (et EventType) IsValid() bool {
	switch et {
	case Game, Practice, Misc:
		return true
	default:
		return false
	}
}

func ParseEventType(s string) (EventType, error) {
	switch EventType(s) {
	case Game, Practice, Misc:
		return EventType(s), nil
	default:
		return "", fmt.Errorf("invalid EventType: %s", s)
	}
}
