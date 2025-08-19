package bem

import (
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Bem struct {
	collector *colly.Collector
	mutex     sync.Mutex

	Events []BemEvent
}

type BemEvent struct {
	DateTime    time.Time
	Title       string
	BookingLink string
	HunDub      bool
	EnglishSubs bool
}
