package hb

import (
	"strings"
	"testing"
)

func TestNewDiv(t *testing.T) {
	header := NewDiv()
	html := header.ToHTML()

	if strings.Contains(html, "<div>") == false {
		t.Error("Does not contain '<div>', Output:" + html)
	}

	if strings.Contains(html, "</div>") == false {
		t.Error("Does not contain '</div>', Output:" + html)
	}
}
