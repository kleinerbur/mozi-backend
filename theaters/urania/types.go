package urania

import (
	"sync"
	"time"
)

type Urania struct {
	baseUrl string
	Events  []UraniaEvent
	mutex   sync.Mutex
}

type UraniaEvent struct {
	BookingLink string
	MovieLink   string
	Title       string
	Auditorium  string
	IsSubbed    bool
	DateTime    time.Time
}
