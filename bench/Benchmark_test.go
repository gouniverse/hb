package main

import (
	"testing"

	"github.com/gouniverse/hb"
)

var user = struct {
	FirstName      string
	FavoriteColors []string
	RawContent     string
	EscapedContent string
}{
	FirstName:      "Bob",
	FavoriteColors: []string{"blue", "green", "mauve"},
	RawContent:     "<div><p>Raw Content to be displayed</p></div>",
	EscapedContent: "<div><div><div>Escaped</div></div></div>",
}

var navigation = []struct {
	Item string
	Link string
}{
	{
		Item: "Link 1",
		Link: "http://www.mytest.com/",
	}, {
		Item: "Link 2",
		Link: "http://www.mytest.com/",
	}, {
		Item: "Link 3",
		Link: "http://www.mytest.com/",
	},
}

// go test -bench=^Benchmark*  -benchtime 5s -count 2 -cpu 1,4
func BenchmarkPageWithHeaderAndUL(b *testing.B) {
	// run the benchmark function b.N times
	for n := 0; n < b.N; n++ {
		ul := hb.UL()
		for i := 0; i < len(user.FavoriteColors); i++ {
			ul.Child(hb.LI().HTML(user.FavoriteColors[i]))
		}

		page := hb.NewTag("html").
			Child(hb.NewTag("body").
				Child(hb.Heading1().HTML(user.FirstName)).
				Child(hb.Paragraph().HTML("Here's a list of your favorite colors:")).
				Child(ul))

		page.ToHTML()
	}

	b.ReportAllocs()
}
