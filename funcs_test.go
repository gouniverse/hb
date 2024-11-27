package hb

import (
	"strconv"
	"strings"
	"testing"
)

func TestToTags(t *testing.T) {
	items := []struct {
		text string
	}{
		{
			text: "One",
		},
		{
			text: "Two",
		},
	}

	spans := ToTags(items, func(item struct{ text string }, index int) TagInterface {
		return NewSpan().
			Text(strconv.Itoa(index + 1)).
			Text(". ").
			Text(item.text)
	})

	html := NewWrap().Children(spans).ToHTML()

	if strings.Contains(html, "<span>1. One</span><span>2. Two</span>") == false {
		t.Error("Does not contain '<span>1. One</span><span>2. Two</span>', Output:" + html)
	}
}
