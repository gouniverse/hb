package hb

import (
	"strings"
	"testing"
)

func TestNewHTML(t *testing.T) {
	s := NewHTML("<script>alert('Hello')</script>")
	html := s.ToHTML()

	if strings.Contains(html, "<script>") == false {
		t.Error("Does not contain '<script>', Output:" + html)
	}

	if strings.Contains(html, "alert('Hello')") == false {
		t.Error("Does not contain 'alert('Hello')', Output:" + html)
	}

	if strings.Contains(html, "</script>") == false {
		t.Error("Does not contain '</script>', Output:" + html)
	}
}
