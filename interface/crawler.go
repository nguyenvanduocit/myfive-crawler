package CrawlerInterface

import (
	"github.com/nguyenvanduocit/myfive-crawler/model/rss"
	"github.com/nguyenvanduocit/myfive-crawler/model/site"
)

type Crawler interface {
	GetTopFive()([]*RssModel.Item, error)
	GetSiteInfo()(*SiteModel.Site, error)
}
