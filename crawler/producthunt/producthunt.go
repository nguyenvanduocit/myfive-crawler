package Producthunt

import (
"github.com/mmcdole/gofeed"
"time"
	"net/http"
	"encoding/json"
	"fmt"
)


// JSON Struct for data recived from product hunt API
type Data struct {
	Posts []*Product `json:"posts"`
}

type Product struct {
	Name string `json:"name"`
	Tagline string `json:"tagline"`
	URL string `json:"url"`
}

type ProductHuntCrawler struct {
	Url string
	Feed *gofeed.Feed
}


func NewCrawler(Url string)(*ProductHuntCrawler){
	return &ProductHuntCrawler{
		Url: Url,
	}
}

func (crawler *ProductHuntCrawler)GetIdentifyURL()(string){
	return crawler.Url
}

func (crawler *ProductHuntCrawler)Parse()(*gofeed.Feed, error){
	if(crawler.Feed != nil){
		return crawler.Feed, nil
	}
	crawler.Feed = &gofeed.Feed{}
	resp, err := http.Get(crawler.Url)
	if err != nil {
		return nil, err
	}

	var respData Data
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	crawler.Feed.Title = "Product Hunt"
	posts := respData.Posts[:10]
	for _, post := range posts{
		item:= &gofeed.Item{}
		item.Title = fmt.Sprintf("%s: %s", post.Name, post.Tagline)
		item.Link = fmt.Sprintf("https://www.producthunt.com%s",post.URL)
		now := time.Now()
		item.PublishedParsed = &now
		crawler.Feed.Items = append(crawler.Feed.Items, item)
	}
	return crawler.Feed, nil
}
