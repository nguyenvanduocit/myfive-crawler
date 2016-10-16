package oxfordlearnersdictionaries

import "testing"

func TestMediumCrawler_Parse(t *testing.T) {
	crawler := NewCrawler("https://en.oxforddictionaries.com/")
	feed,err := crawler.Parse()
	if err != nil {
		t.Error(err)
	}
	t.Error(feed)
	if len(feed.Items) <= 0 {
		t.Error("Can not get items")
	}
}
