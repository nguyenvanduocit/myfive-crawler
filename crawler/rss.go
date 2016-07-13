package RssCrawler

import (
	"net/http"
	"fmt"
	"encoding/xml"
	"github.com/nguyenvanduocit/myfive-crawler/model/rss"
	"github.com/nguyenvanduocit/myfive-crawler/model/site"
)

type RssCrawler struct {
	RssUrl string
}


func NewCrawler(rssUrl string)(*RssCrawler){
	return &RssCrawler{
		RssUrl: rssUrl,
	}
}

func (crawler *RssCrawler)SetRssUrl(url string){
	crawler.RssUrl = url
}

func (crawler *RssCrawler)ParseFeed()(*RssModel.RSS, error){
	response, err := http.Get(crawler.RssUrl)
	if err != nil {
		fmt.Println("Link is invalid")
		return nil, err
	}
	defer response.Body.Close()
	xmlDecoder := xml.NewDecoder(response.Body)

	var rss RssModel.RSS

	if err = xmlDecoder.Decode(&rss); err != nil {
		return nil, err
	}
	return &rss, nil

}

func (crawler *RssCrawler)GetSiteInfo()(*SiteModel.Site, error){
	rss, err := crawler.ParseFeed()
	if err != nil {
		return nil, err
	}
	return &SiteModel.Site{
		Title: rss.Channel.Title,
		Link: rss.Channel.Link,
	}, nil
}

func (crawler *RssCrawler)GetTopFive()([]*RssModel.Item, error){
	rss, err := crawler.ParseFeed()
	if err != nil {
		return nil, err
	}
	return rss.Channel.Item[:5], nil
}

