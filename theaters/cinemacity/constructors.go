package cinemacity

import "sync"

func new(id int) (*CinemaCity, error) {
	cc := CinemaCity{
		id,
		locationName[id],
		make([]CinemaCityMovie, 0),
		make([]CinemaCityEvent, 0),
		sync.Mutex{},
	}
	err := cc.init()
	if err != nil {
		return nil, err
	}
	return &cc, nil
}

func Allee() (*CinemaCity, error) {
	return new(allee)
}

func Arena() (*CinemaCity, error) {
	return new(arena)
}

func Campona() (*CinemaCity, error) {
	return new(campona)
}

func DunaPlaza() (*CinemaCity, error) {
	return new(dunaplaza)
}

func Mammut() (*CinemaCity, error) {
	return new(mammut)
}

func WestEnd() (*CinemaCity, error) {
	return new(westend)
}
