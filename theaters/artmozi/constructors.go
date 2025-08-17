package artmozi

import "sync"

func new(ID int) (*ArtMozi, error) {
	am := &ArtMozi{
		ID,
		baseUrls[ID],
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
	return new(CORVIN)
}

func Puskin() (*ArtMozi, error) {
	return new(PUSKIN)
}

func Muvesz() (*ArtMozi, error) {
	return new(MUVESZ)
}

func Taban() (*ArtMozi, error) {
	return new(TABAN)
}

func Kino() (*ArtMozi, error) {
	return new(KINO)
}

func Toldi() (*ArtMozi, error) {
	return new(TOLDI)
}
