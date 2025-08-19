package artmozi

import "sync"

func new(id int) (*ArtMozi, error) {
	am := &ArtMozi{
		cinemaID: id,
		name:     locationName[id],
		baseUrl:  baseUrl[id],
		mutex:    sync.Mutex{},

		Movies: make(map[int]ArtMoziMovie, 0),
		Events: make([]ArtMoziEvent, 0),
	}
	err := am.init()
	if err != nil {
		return nil, err
	}
	return am, nil
}

func Corvin() (*ArtMozi, error) {
	return new(corvin)
}

func Puskin() (*ArtMozi, error) {
	return new(puskin)
}

func Muvesz() (*ArtMozi, error) {
	return new(muvesz)
}

func Taban() (*ArtMozi, error) {
	return new(taban)
}

func Kino() (*ArtMozi, error) {
	return new(kino)
}

func Toldi() (*ArtMozi, error) {
	return new(toldi)
}
