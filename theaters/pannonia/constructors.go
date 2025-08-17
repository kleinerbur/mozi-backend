package pannonia

import (
	"sync"

	"github.com/gocolly/colly"
)

func new(id int) (*Pannonia, error) {
	p := &Pannonia{
		locationName[id],
		baseUrl[id],
		colly.NewCollector(),
		make([]PannoniaEvent, 0),
		&sync.Mutex{},
	}
	err := p.init()
	return p, err
}

func Mom() (*Pannonia, error) {
	return new(mom)
}

func Polus() (*Pannonia, error) {
	return new(polus)
}

func Lurdy() (*Pannonia, error) {
	return new(lurdy)
}

func GoBuda() (*Pannonia, error) {
	return new(gobuda)
}
