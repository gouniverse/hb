package hb

import (
	"strings"
	"testing"
)

func TestNewSelect(t *testing.T) {
	s := NewSelect()
	html := s.ToHTML()

	if strings.Contains(html, "<select>") == false {
		t.Error("Does not contain '<select>', Output:" + html)
	}

	if strings.Contains(html, "</select>") == false {
		t.Error("Does not contain '</select>', Output:" + html)
	}
}
