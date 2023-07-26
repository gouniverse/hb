package hb

import (
	"strings"
	"testing"
)

func TestNewCaption(t *testing.T) {
	caption := NewCaption()
	html := caption.ToHTML()

	if strings.Contains(html, "<caption>") == false {
		t.Error("Does not contain '<caption>', Output:" + html)
	}

	if strings.Contains(html, "</caption>") == false {
		t.Error("Does not contain '</caption>', Output:" + html)
	}
}