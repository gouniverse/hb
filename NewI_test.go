package hb

import (
	"strings"
	"testing"
)

func TestNewI(t *testing.T) {
	i := NewI()
	html := i.ToHTML()

	if strings.Contains(html, "<i>") == false {
		t.Error("Does not contain '<i>', Output:" + html)
	}
	if strings.Contains(html, "</i>") == false {
		t.Error("Does not contain '</i>', Output:" + html)
	}
}

func TestNewIWithContent(t *testing.T) {
	i := NewI().HTML("italic")
	html := i.ToHTML()

	if strings.Contains(html, "<i>") == false {
		t.Error("Does not contain '<i>', Output:" + html)
	}
	if strings.Contains(html, "</i>") == false {
		t.Error("Does not contain '</i>', Output:" + html)
	}
	if strings.Contains(html, "italic") == false {
		t.Error("Does not contain 'italic', Output:" + html)
	}
}
