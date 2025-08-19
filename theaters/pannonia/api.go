package pannonia

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

func (p *Pannonia) init() error {
	p.collector.OnHTML("div.order-header div.meta", func(eventData *colly.HTMLElement) {
		r := regexp.MustCompile(`\d. terem`)
		if r.MatchString(eventData.Text) {
			bookingLink := eventData.Request.URL.Path[1:]
			auditorium := r.FindString(eventData.Text)
			p.Events[bookingLink].Auditorium = auditorium
		}
	})

	p.collector.OnHTML("div#details-wrapper", func(movieData *colly.HTMLElement) {
		movieLink := movieData.Request.URL.Path[1:]
		p.Movies[movieLink] = &PannoniaMovie{
			movieLink,
			parseTitle(movieData),
			parseOriginalTitle(movieData),
			parseYear(movieData),
		}
	})

	p.collector.OnHTML(".day-wrapper", func(dayWrapper *colly.HTMLElement) {
		date := parseDate(dayWrapper)
		if date.Before(time.Now().AddDate(0, 0, 15)) {
			dayWrapper.ForEach(".movie-wrapper", func(idx int, movieWrapper *colly.HTMLElement) {
				title := parseTitle(movieWrapper)
				movieLink := parseMovieLink(movieWrapper)
				if isAnActualMovie(title) {
					movieWrapper.ForEach(".movie-time", func(idx int, movieTime *colly.HTMLElement) {
						bookingLink := parseBookingLink(movieTime)

						p.Events[bookingLink] = &PannoniaEvent{
							MovieLink:   movieLink,
							BookingLink: bookingLink,
							Title:       title,
							Auditorium:  "",
							DateTime:    parseDateTime(date, movieTime),
							IsSubbed:    parseIsSubbed(movieTime),
						}
						p.collector.Visit(p.baseUrl + bookingLink)
						p.collector.Visit(p.baseUrl + movieLink)
					})
				}
			})
		}
	})
	return p.collector.Visit(p.baseUrl + "#musorlista")
}
