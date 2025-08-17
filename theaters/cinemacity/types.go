package cinemacity

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

type CinemaCity struct {
	cinemaID int
	Movies   []CinemaCityMovie
	Events   []CinemaCityEvent
	mutex    sync.Mutex
}

type CinemaCityMovie struct {
	ID           string   `json:"id"`
	Title        string   `json:"name"`
	Length       int      `json:"length"`
	PosterLink   string   `json:"posterLink"`
	Link         string   `json:"link"`
	ReleaseYear  string   `json:"releaseYear"`
	ReleaseDate  string   `json:"releaseDate"`
	AttributeIds []string `json:"attributeIds"`
}

type CinemaCityEvent struct {
	ID                 string    `json:"id"`
	FilmID             string    `json:"filmId"`
	CinemaID           string    `json:"cinemaId"`
	DateTime           time.Time `json:"eventDateTime"`
	AttributeIDs       []string  `json:"attributeIds"`
	BookingLink        string    `json:"bookingLink"`
	SoldOut            bool      `json:"soldOut"`
	Auditorium         string    `json:"auditorium"`
	AuditoriumTinyName string    `json:"auditoriumTinyName"`
}

func (e *CinemaCityEvent) UnmarshalJSON(data []byte) error {
	type Alias CinemaCityEvent
	aux := &struct {
		EventDateTime string `json:"eventDateTime"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02T15:04:05", aux.EventDateTime)
	if err != nil {
		return fmt.Errorf("invalid eventDateTime: %w", err)
	}
	e.DateTime = t
	e.BookingLink = strings.Replace(e.BookingLink, "/api", "", 1)
	return nil
}
