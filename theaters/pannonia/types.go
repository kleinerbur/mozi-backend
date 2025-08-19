package pannonia

import (
	"time"

	"github.com/gocolly/colly"
)

type Pannonia struct {
	name      string
	baseUrl   string
	collector *colly.Collector

	Movies map[string]*PannoniaMovie
	Events map[string]*PannoniaEvent
}

type PannoniaMovie struct {
	MovieLink     string
	Title         string
	OriginalTitle string
	Year          int
}

type PannoniaEvent struct {
	DateTime    time.Time
	MovieLink   string
	BookingLink string
	Title       string
	Auditorium  string
	IsSubbed    bool
}
