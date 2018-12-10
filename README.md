# CRAWL-QUOTE

Quotes from [goodsread](https://www.goodreads.com) crawler built using golang

## Getting Started

- import the library
```go
import "github.com/faruqisan/crawl-quote"
```
- example
```go

quoteEngine := quote.New()

// get popular quotes
quotes, err := quoteEngine.GetPopulars()
if err != nil {
// handle err
}

// print the quotes
log.Println(quotes)

```
