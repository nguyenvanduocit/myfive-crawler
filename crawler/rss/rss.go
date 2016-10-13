package RssCrawler

import (
	"github.com/mmcdole/gofeed"
)

type RssCrawler struct {
	RssUrl string
	Parser *gofeed.Parser
	Feed *gofeed.Feed
}


func NewCrawler(RssUrl string)(*RssCrawler){
	return &RssCrawler{
		RssUrl: RssUrl,
		Parser: gofeed.NewParser(),
	}
}

func (crawler *RssCrawler)GetIdentifyURL()(string){
	return crawler.RssUrl
}

func (crawler *RssCrawler)Parse()(*gofeed.Feed, error){
	if(crawler.Feed != nil){
		return crawler.Feed, nil
	}
	var err error
	crawler.Feed, err = crawler.Parser.ParseURL(crawler.RssUrl)
	if (err != nil){
		return nil, err
	}
	return crawler.Feed, nil

}
