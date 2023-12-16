package hb

import (
	"strings"
	"testing"
)

func TestNewForm(t *testing.T) {
	header := NewForm()
	html := header.ToHTML()

	if strings.Contains(html, "<form>") == false {
		t.Error("Does not contain '<form>', Output:" + html)
	}

	if strings.Contains(html, "</form>") == false {
		t.Error("Does not contain '</form>', Output:" + html)
	}
}
