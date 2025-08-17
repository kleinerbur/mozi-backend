package cinemacity

import "sync"

func new(ID int) (*CinemaCity, error) {
	cc := CinemaCity{
		ID,
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
	return new(ALLEE)
}

func Arena() (*CinemaCity, error) {
	return new(ARENA)
}

func Campona() (*CinemaCity, error) {
	return new(CAMPONA)
}

func DunaPlaza() (*CinemaCity, error) {
	return new(DUNAPLAZA)
}

func Mammut() (*CinemaCity, error) {
	return new(MAMMUT)
}

func WestEnd() (*CinemaCity, error) {
	return new(WESTEND)
}
