package hb

import (
	"strings"
	"testing"
)

func TestNewHeader(t *testing.T) {
	header := NewHeader()
	html := header.ToHTML()

	if strings.Contains(html, "<header>") == false {
		t.Error("Does not contain '<header>', Output:" + html)
	}

	if strings.Contains(html, "</header>") == false {
		t.Error("Does not contain '</header>', Output:" + html)
	}
}
