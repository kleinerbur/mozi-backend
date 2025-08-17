package pannonia

import (
	"time"

	"github.com/gocolly/colly"
)

func (p *Pannonia) init() error {
	p.collector.OnHTML(".day-wrapper", func(dayWrapper *colly.HTMLElement) {
		date := parseDate(dayWrapper.DOM.Find(".date").Text())
		if date.Before(time.Now().AddDate(0, 0, 15)) {
			dayWrapper.ForEach(".movie-wrapper", func(idx int, movieWrapper *colly.HTMLElement) {
				title := trim(movieWrapper.DOM.Find(".title").Text())
				if isAnActualMovie(title) {
					premiere := movieWrapper.DOM.Find(".premiere").Text() == "Premier"
					movieWrapper.ForEach(".movie-time", func(idx int, movieTime *colly.HTMLElement) {
						bookingLink := movieTime.ChildAttr("a", "href")
						timestamp, _ := time.Parse("15:04", movieTime.DOM.Find(".time").Text())
						subbed := movieTime.DOM.Find(".type").Text() == "F"

						p.mutex.Lock()
						p.Events = append(p.Events, PannoniaEvent{
							title,
							combine(date, timestamp),
							p.baseUrl + bookingLink,
							subbed,
							premiere,
						})
						p.mutex.Unlock()
					})
				}
			})
		}
	})
	return p.collector.Visit(p.baseUrl + "#musorlista")
}
