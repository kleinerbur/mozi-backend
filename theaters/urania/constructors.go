package urania

import (
	"sync"
)

func Mozi() (*Urania, error) {
	urania := &Urania{
		baseUrl: "https://urania-nf.hu",
		mutex:   sync.Mutex{},

		Events: make([]UraniaEvent, 0),
	}
	err := urania.init()
	if err != nil {
		return nil, err
	}
	return urania, nil
}
