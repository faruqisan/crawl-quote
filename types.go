package quote

import (
	"regexp"

	"github.com/gocolly/colly"
)

// Engine struct act as function receiver
type Engine struct {
	c             *colly.Collector
	contentRegexp *regexp.Regexp
}

// Data struct define model of quote data
type Data struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
