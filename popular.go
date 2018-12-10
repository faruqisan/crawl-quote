package quote

import (
	"errors"

	"github.com/gocolly/colly"
)

// GetPopulars function return quotes from popular page
func (e Engine) GetPopulars() ([]Data, error) {
	var (
		err    error
		quotes []Data
	)

	e.c.OnHTML(popularQuotesHTMLClass, func(el *colly.HTMLElement) {

		el.ForEach(popularQuoteHTMLClass, func(i int, q *colly.HTMLElement) {
			var (
				sanitizedQuote string
				quote          Data
			)

			sanitizedQuote, err = e.sanitizeQuote(q.ChildText(popularTextHTMLClass))
			if err != nil {
				return
			}

			if len(quotes) < defaultLimit {
				quote.Content = sanitizedQuote
				quote.Author = q.ChildText(popularAuthorHTMLClass)

				quotes = append(quotes, quote)
			}

		})

	})

	e.c.OnHTML(nextPageHTMLClass, func(el *colly.HTMLElement) {
		if len(quotes) < defaultLimit {
			el.Request.Visit(el.Attr("href"))
		}
	})

	e.c.Visit(popularQuotesURL)

	return quotes, err
}

func (e Engine) sanitizeQuote(input string) (string, error) {
	var (
		err    = errors.New("bad input")
		result string
		raw    [][]string
	)

	raw = e.contentRegexp.FindAllStringSubmatch(input, -1)

	if len(raw) < 1 {
		return result, err
	}

	if len(raw[0]) < 1 {
		return result, err
	}

	result = raw[0][0]

	return result, nil
}
