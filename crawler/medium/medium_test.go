package medium

import "testing"

func TestMediumCrawler_Parse(t *testing.T) {
	crawler := MediumCrawler{
		Url:"https://medium.com/browse/top",
	}
	feed,err := crawler.Parse()
	if err != nil {
		t.Error(err)
	}
	t.Error(feed)
	if len(feed.Items) <= 0 {
		t.Error("Can not get items")
	}
}
