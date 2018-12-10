package quote

import (
	"regexp"

	"github.com/gocolly/colly"
)

// New function return configured quote engine
func New() Engine {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowURLRevisit(),
		colly.DetectCharset(),
	)

	return Engine{
		c:             c,
		contentRegexp: regexp.MustCompile(regex),
	}
}
