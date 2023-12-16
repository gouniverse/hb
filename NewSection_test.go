package hb

import (
	"strings"
	"testing"
)

func TestNewSection(t *testing.T) {
	s := NewSection()
	html := s.ToHTML()

	if strings.Contains(html, "<section>") == false {
		t.Error("Does not contain '<section>', Output:" + html)
	}

	if strings.Contains(html, "</section>") == false {
		t.Error("Does not contain '</section>', Output:" + html)
	}
}
