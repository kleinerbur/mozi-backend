package cinemacity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"sort"
	"time"

	"golang.org/x/sync/errgroup"
)

func (cc *CinemaCity) getMovies(until time.Time) error {
	resp, err := http.Get(
		baseUrl + fmt.Sprintf(
			"films/until/%s",
			until.Format("2006-01-02"),
		))
	if err != nil {
		return fmt.Errorf("GET films @ CinemaCity %s failed: '%w'", locationNames[cc.cinemaID], err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("GET films @ CinemaCity %s failed: '%w'", locationNames[cc.cinemaID], err)
	}
	var jsonResp struct {
		Body struct {
			Films []CinemaCityMovie `json:"films"`
		} `json:"body"`
	}
	json.Unmarshal(body, &jsonResp)
	cc.Movies = append(cc.Movies, jsonResp.Body.Films...)
	return nil
}

func (cc *CinemaCity) getEvents(date time.Time) error {
	resp, err := http.Get(
		baseUrl + fmt.Sprintf(
			"film-events/in-cinema/%d/at-date/%s",
			cc.cinemaID,
			date.Format("2006-01-02"),
		),
	)
	if err != nil {
		return fmt.Errorf("GET events @ CinemaCity %s failed: '%w'", locationNames[cc.cinemaID], err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("GET events @ CinemaCity %s failed: '%w'", locationNames[cc.cinemaID], err)
	}
	var jsonResp struct {
		Body struct {
			Events []CinemaCityEvent `json:"events"`
		} `json:"body"`
	}
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return fmt.Errorf("GET events @ CinemaCity %s failed: '%w'", locationNames[cc.cinemaID], err)
	}
	cc.mutex.Lock()
	cc.Events = append(cc.Events, jsonResp.Body.Events...)
	cc.mutex.Unlock()
	return nil
}

func (cc *CinemaCity) EventsOf(movie *CinemaCityMovie) []CinemaCityEvent {
	filtered := slices.Collect(
		func(yield func(CinemaCityEvent) bool) {
			for _, event := range cc.Events {
				if event.FilmID == movie.ID {
					if !yield(event) {
						return
					}
				}
			}
		},
	)
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].DateTime.Before(filtered[j].DateTime)
	})
	return filtered
}

func (cc *CinemaCity) init() error {
	now := time.Now()
	nDays := 14
	err := cc.getMovies(now.AddDate(0, 0, nDays))
	if err != nil {
		return err
	}
	var grp errgroup.Group
	for i := range nDays {
		deltaDays := i
		grp.Go(func() error {
			return cc.getEvents(now.AddDate(0, 0, deltaDays))
		})
	}
	if err := grp.Wait(); err != nil {
		return err
	}
	return nil
}
