package pannonia

import (
	"github.com/gocolly/colly"
)

func new(id int) (*Pannonia, error) {
	p := &Pannonia{
		locationName[id],
		baseUrl[id],
		colly.NewCollector(),
		make(map[string]*PannoniaMovie, 0),
		make(map[string]*PannoniaEvent, 0),
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
