package RssModel

type RSS struct {
	Channel Channel `xml:"channel"`
}

//Channel struct for RSS
type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Item          []*Item `xml:"item"`
}

type Item struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Comments    string          `xml:"comments"`
	PubDate     string        `xml:"pubDate"`
	GUID        string          `xml:"guid"`
	Category    []string        `xml:"category"`
	Enclosure   []*ItemEnclosure `xml:"enclosure"`
	Description string          `xml:"description"`
	Content     string          `xml:"content"`
}

type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}
