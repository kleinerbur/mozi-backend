package artmozi

import "sync"

func new(id int) (*ArtMozi, error) {
	am := &ArtMozi{
		id,
		locationName[id],
		baseUrl[id],
		make(map[int]ArtMoziMovie, 0),
		make([]ArtMoziEvent, 0),
		sync.Mutex{},
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
