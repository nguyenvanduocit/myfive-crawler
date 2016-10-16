package rss

import "testing"

func TestRssCrawler_Parse(t *testing.T) {
	crawler := NewCrawler("https://www.sitepoint.com/feed")
	feed,err := crawler.Parse()
	if err != nil {
		t.Error(err)
	}
	t.Error(feed)
	if len(feed.Items) <= 0 {
		t.Error("Can not get items")
	}
}