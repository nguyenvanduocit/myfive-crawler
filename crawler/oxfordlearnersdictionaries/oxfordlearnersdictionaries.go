package oxfordlearnersdictionaries

import (
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
	"time"
	"fmt"
)

type RssCrawler struct {
	Url string
	Feed *gofeed.Feed
}


func NewCrawler(RssUrl string)(*RssCrawler){
	return &RssCrawler{
		Url: RssUrl,
	}
}

func (crawler *RssCrawler)GetIdentifyURL()(string){
	return crawler.Url
}

func (crawler *RssCrawler)Parse()(*gofeed.Feed, error){
	if(crawler.Feed != nil){
		return crawler.Feed, nil
	}
	doc, err := goquery.NewDocument(crawler.Url)
	if err != nil {
		return nil, err
	}
	crawler.Feed = &gofeed.Feed{
		Title: "Word Of Day",
		Items:[]*gofeed.Item{},
	}
	linkword := doc.Find(".daywordmain .linkword")
	detailUrl, exist := linkword.Attr("href")
	detailUrl = fmt.Sprintf("https://en.oxforddictionaries.com%s", detailUrl)
	if !exist {
		return nil, fmt.Errorf("Can not found word's detail link: %s", linkword.Text())
	}
	crawler.Feed.Items, err = crawler.ParseDetail(detailUrl)
	if err != nil{
		return nil, err
	}

	return crawler.Feed, nil

}

func (crawler *RssCrawler)ParseDetail(url string)([]*gofeed.Item, error){
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	word := doc.Find(".hw").Text()
	prononciation := doc.Find("section.pronSection.etym > div > span.phoneticspelling").Text()
	var items []*gofeed.Item
	doc.Find(".entryWrapper section.gramb").Each(func(i int, s *goquery.Selection) {
		item:= &gofeed.Item{}
		wordType := s.Find("h3.pos").Text()
		firstMean := s.Find(".semb .ind").Text()
		item.Title = fmt.Sprintf("%s(%s)(%s): %s", word, wordType, prononciation, firstMean)
		item.Link =  url
		now := time.Now()
		item.PublishedParsed = &now

		items = append(items, item)
	})
	return items, nil
}
