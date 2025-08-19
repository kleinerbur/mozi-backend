package bem

import (
	"sync"

	"github.com/gocolly/colly"
)

func Mozi() (*Bem, error) {
	mozi := &Bem{
		collector: colly.NewCollector(),
		mutex:     sync.Mutex{},

		Events: make([]BemEvent, 0),
	}
	err := mozi.init()
	if err != nil {
		return nil, err
	}
	return mozi, nil
}
