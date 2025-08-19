package bem

import (
	"strings"

	"github.com/forPelevin/gomoji"
)

func parseTitleAndTags(raw string) (string, string) {
	title := gomoji.RemoveEmojis(raw)
	if strings.Contains(title, ": ") {
		split := strings.Split(title, ": ")
		title = split[len(split)-1]
	}
	tags := ""
	if strings.Contains(title, "|") {
		split := strings.Split(title, "|")
		title = split[0]
		tags = strings.ToLower(split[1])
	}
	if strings.Contains(title, "/") {
		title = strings.Split(title, "/")[1]
	}
	return title, tags
}

func isAnActualMovie(rawTitle string) bool {
	return !strings.Contains(rawTitle, giftVoucherTag)
}
