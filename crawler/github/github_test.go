package github

import (
	"testing"
)

func TestGithubCrawler_Parse(t *testing.T) {
	crawler := GithubCrawler{
		Url:"https://github.com/trending",
	}
	feed,err := crawler.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(feed.Items) <= 0 {
		t.Error("Can not get items")
	}
}