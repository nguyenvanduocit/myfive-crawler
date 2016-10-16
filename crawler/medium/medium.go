package medium

import (
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"time"
)

type MediumCrawler struct {
	Url string
	Feed *gofeed.Feed
}


func NewCrawler(Url string)(*MediumCrawler){
	return &MediumCrawler{
		Url: Url,
	}
}

func (crawler *MediumCrawler)GetIdentifyURL()(string){
	return crawler.Url
}

func (crawler *MediumCrawler)Parse()(*gofeed.Feed, error){
	if(crawler.Feed != nil){
		return crawler.Feed, nil
	}
	doc, err := goquery.NewDocument(crawler.Url)
	if err != nil {
		return nil, err
	}
	crawler.Feed = &gofeed.Feed{
		Title: doc.Find("title").Text(),
		Items:[]*gofeed.Item{},
	}
	doc.Find(".js-homeStream .streamItem").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".graf--title")
		item:= &gofeed.Item{}
		item.Title = title.Text()
		item.Link,_ = s.Find(".postArticle-content a").Attr("href")
		now := time.Now()
		item.PublishedParsed = &now
		crawler.Feed.Items = append(crawler.Feed.Items, item)
	})
	return crawler.Feed, nil

}
