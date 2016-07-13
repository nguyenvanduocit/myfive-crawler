package util

import (
	"strings"
	"time"
	"fmt"
)

// TimeLayouts is contains a list of time.Parse() layouts that are used in
// attempts to convert item.Date and item.PubDate string to time.Time values.
// The layouts are attempted in ascending order until either time.Parse()
// does not return an error or all layouts are attempted.
var TimeLayouts = []string{
	"Mon, _2 Jan 2006 15:04:05 +0000",
	"Mon, _2 Jan 2006 15:04:05 MST",
	"Mon, _2 Jan 2006 15:04:05 -0700",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	"Mon, 2, Jan 2006 15:4",
	"02 Jan 2006 15:04:05 MST",
}



func ParseTime(s string) (time.Time, error) {
	timeString := strings.TrimSpace(s)
	var t time.Time
	for _, layout := range TimeLayouts {
		t, err := time.Parse(layout, timeString)
		if err != nil {
			return t, err
		}
		return t, nil
	}
	return t, fmt.Errorf("Can not found time layout for %s", timeString)
}
