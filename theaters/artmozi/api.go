package artmozi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (am *ArtMozi) init() error {
	now := time.Now()
	if err := am.getData(now); err != nil {
		return err
	}
	return am.getData(now.AddDate(0, 0, 7))
}

func (am *ArtMozi) getData(date time.Time) error {
	resp, err := http.Get(
		am.baseUrl + fmt.Sprintf(
			"schedule/week/%d%d",
			date.Year(),
			ScheduleWeek(date),
		),
	)
	if err != nil {
		return fmt.Errorf("GET events @ ArtMozi %s failed: '%w'", am.name, err)
	}
	defer resp.Body.Close()
	var jsonResp ArtMoziResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return fmt.Errorf("GET events @ ArtMozi %s failed: '%w'", am.name, err)
	}
	am.addMovies(jsonResp.Movies)
	am.addEvents(jsonResp.Events)
	return nil
}

func (am *ArtMozi) addMovies(movies map[int]ArtMoziMovie) {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	for _, movie := range movies {
		if _, duplicate := am.Movies[movie.ID]; !duplicate {
			am.Movies[movie.ID] = movie
		}
	}
}

func (am *ArtMozi) addEvents(events []ArtMoziEvent) {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	am.Events = append(am.Events, events...)
}
