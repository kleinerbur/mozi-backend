package artmozi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ArtMoziMovieInfo struct {
	Title    string `json:"title"`
	AgeLimit string `json:"ageLimit"`
	ImgUrl   string `json:"imgUrl"`
	Url      string `json:"cannonicalUrl"`
}

type ArtMoziMovie struct {
	ID          int
	Title       string
	AgeLimit    string
	ImgUrl      string
	BookingLink string
}

type ArtMoziEventInfo struct {
	CinemaID     int    `json:"cinema"`
	Auditorium   string `json:"cinema_room"`
	Link         string `json:"link"`
	VisualEffect string `json:"visualEffect"`
	DubSub       string `json:"dubSub"`
	DubSubCode   string `json:"dubSubCode"`
	Premiere     string `json:"premiere"`
	Accessible   bool   `json:"accessible"`
}

type ArtMoziEvent struct {
	FilmID       string
	DateTime     time.Time
	EventID      string
	CinemaID     int
	Auditorium   string
	Link         string
	VisualEffect string
	DubSub       string
	DubSubCode   string
	Premiere     bool
	Accessible   bool
}

type ArtMoziResponse struct {
	Movies map[int]ArtMoziMovie // quasi set
	Events []ArtMoziEvent
}

type ArtMoziEventRawDate string
type ArtMoziEventRawID string
type ArtMoziEventRaw map[ArtMoziEventRawID]ArtMoziEventInfo

func (r *ArtMoziResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Movies   map[string]ArtMoziMovieInfo                                  `json:"movies"`
		Schedule map[string]map[string]map[string]map[string]ArtMoziEventInfo `json:"schedule"`
		// date (str) -> filmID -> time (str) -> eventID -> eventInfo
	}
	json.Unmarshal(data, &raw)
	r.Movies = make(map[int]ArtMoziMovie)
	for movieId, movieInfo := range raw.Movies {
		id, err := strconv.Atoi(movieId)
		if err != nil {
			return fmt.Errorf("invalid movie ID: %w", err)
		}
		r.Movies[id] = ArtMoziMovie{
			ID:          id,
			Title:       movieInfo.Title,
			AgeLimit:    movieInfo.AgeLimit,
			ImgUrl:      movieInfo.ImgUrl,
			BookingLink: movieInfo.Url,
		}
	}

	var entries []ArtMoziEvent
	for dateStr, films := range raw.Schedule {
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			return fmt.Errorf("parse date: %w", err)
		}
		for filmId, times := range films {
			for timeStr, events := range times {
				t, err := time.Parse("15:04", timeStr)
				if err != nil {
					return fmt.Errorf("parse time: %w", err)
				}
				dateTime := time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
				for eventId, event := range events {
					entries = append(entries, ArtMoziEvent{
						FilmID:       filmId,
						DateTime:     dateTime,
						EventID:      eventId,
						CinemaID:     event.CinemaID,
						Auditorium:   event.Auditorium,
						Link:         event.Link,
						VisualEffect: event.VisualEffect,
						DubSub:       event.DubSub,
						DubSubCode:   event.DubSubCode,
						Premiere:     event.Premiere == "Premiere",
						Accessible:   event.Accessible,
					})
				}
			}
		}
	}
	r.Events = entries
	return nil
}

type ArtMozi struct {
	cinemaID int
	name     string
	baseUrl  string
	Movies   map[int]ArtMoziMovie
	Events   []ArtMoziEvent
	mutex    sync.Mutex
}
