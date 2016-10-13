package CrawlerInterface

import (
	"github.com/mmcdole/gofeed"
)

type Crawler interface {
	Parse()(*gofeed.Feed, error)
	GetIdentifyURL()(string)
}
