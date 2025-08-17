package pannonia

import (
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Pannonia struct {
	name      string
	baseUrl   string
	collector *colly.Collector
	Events    []PannoniaEvent
	mutex     *sync.Mutex
}

type PannoniaEvent struct {
	Title       string
	DateTime    time.Time
	BookingLink string
	Subbed      bool
	Premiere    bool
}
