package bem

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

func (b *Bem) parseEvent(node *goquery.Selection) error {
	rawTitle := node.Find(classTitle).Text()
	if isAnActualMovie(rawTitle) {
		title, tags := parseTitleAndTags(rawTitle)
		dateTime, err := time.Parse(
			"Monday, Jan 2, 15:04 PM 2006 -07:00",
			fmt.Sprintf(
				"%s %s",
				node.Find(classDateTime).Text(),
				time.Now().Format("2006 -07:00"),
			),
		)
		if err != nil {
			return err
		}
		b.mutex.Lock()
		b.Events = append(b.Events, BemEvent{
			Title:       strings.TrimSpace(title),
			DateTime:    dateTime,
			BookingLink: baseUrl + node.AttrOr("href", ""),
			HunDub:      strings.Contains(tags, hunDubTag),
			EnglishSubs: strings.Contains(tags, engSubTag),
		})
		b.mutex.Unlock()
	}
	return nil
}

func (b *Bem) init() error {
	resp, err := http.Get(baseUrl + schedulePath)
	if err != nil {
		return fmt.Errorf("GET events @ Bem failed: %w", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("GET events @ Bem failed: %w", err)
	}
	errGroup := errgroup.Group{}
	doc.Find("main a").Each(func(i int, node *goquery.Selection) {
		errGroup.Go(func() error {
			return b.parseEvent(node)
		})
	})
	return errGroup.Wait()
}
