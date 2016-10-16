package Github

import (
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"time"
)

type GithubCrawler struct {
	Url string
	Feed *gofeed.Feed
}


func NewCrawler(Url string)(*GithubCrawler){
	return &GithubCrawler{
		Url: Url,
	}
}

func (crawler *GithubCrawler)GetIdentifyURL()(string){
	return crawler.Url
}

func (crawler *GithubCrawler)Parse()(*gofeed.Feed, error){
	if(crawler.Feed == nil){
		return crawler.Feed, nil
	}

	doc, err := goquery.NewDocument(crawler.Url)
	if err != nil {
		return nil, err
	}
	crawler.Feed = &gofeed.Feed{
		Title: doc.Find("title").Text(),
	}
	doc.Find(".explore-content .repo-list .repo-list-item").Each(func(i int, s *goquery.Selection) {
		item:= &gofeed.Item{}
		item.Title, _ = s.Find(".repo-list-name a").Attr("href")
		item.Link = "https://github.com" + item.Title
		now := time.Now()
		item.PublishedParsed = &now
		crawler.Feed.Items = append(crawler.Feed.Items, item)
	})
	return crawler.Feed, nil

}
