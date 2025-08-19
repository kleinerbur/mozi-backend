package urania

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

func (u *Urania) init() error {
	now := time.Now()
	nDays := 14
	var grp errgroup.Group
	for i := range nDays {
		deltaDays := i
		grp.Go(func() error {
			return u.getEvents(now.AddDate(0, 0, deltaDays))
		})
	}
	if err := grp.Wait(); err != nil {
		return err
	}
	sort.SliceStable(u.Events, func(i, j int) bool {
		return u.Events[i].DateTime.Before(u.Events[j].DateTime)
	})
	return nil
}

func getIsSubbed(movieLink string) bool {
	resp, err := http.Get(movieLink)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false
	}
	return strings.Contains(doc.Find(".film_fontos_infok").Text(), "felirattal")
}

func (u *Urania) getEvents(date time.Time) error {
	resp, err := http.Post(
		"https://urania-nf.hu/ajax.php",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("module=Musor&c=napiMusor&date=%d", date.Unix())),
	)
	if err != nil {
		return fmt.Errorf("GET events @ Urania failed: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("GET events @ Urania failed: %w", err)
	}

	var bookingLink string
	var movieLink string
	var title string
	var auditorium string
	var dateTime time.Time
	var isSubbed bool

	schedule := doc.Find(".musor_tartalom")
	schedule.Children().Each(func(i int, node *goquery.Selection) {
		switch node.Nodes[0].Data {
		case "a":
			if node.AttrOr("class", "") != "jegy_mob_call" {
				text := node.Text()
				if text == "Jegyvásárlás" {
					bookingLink = node.AttrOr("href", "")
				} else {
					title = strings.Split(text, "\u00a0")[0]
					movieLink = u.baseUrl + node.AttrOr("href", "")
					isSubbed = getIsSubbed(movieLink)
				}
			}
		case "span":
			text := strings.Split(node.Text(), "\n")[0]
			dateTime, _ = time.Parse(
				"2006-01-02 -07:00 15:04",
				fmt.Sprintf(
					"%s %s",
					date.Format("2006-01-02 07:00"),
					strings.Split(text, "\u00a0")[0],
				),
			)
			auditorium = strings.Join(strings.Split(text, " ")[1:], " ")
		case "div":
			u.mutex.Lock()
			u.Events = append(u.Events, UraniaEvent{
				BookingLink: bookingLink,
				MovieLink:   movieLink,
				Title:       title,
				Auditorium:  auditorium,
				DateTime:    dateTime,
				IsSubbed:    isSubbed,
			})
			u.mutex.Unlock()
		}
	})
	return nil
}
