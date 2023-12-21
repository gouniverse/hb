package hb

import (
	"strings"
	"testing"
)

func TestNewImage(t *testing.T) {
	s := NewImage()
	html := s.ToHTML()

	if strings.Contains(html, "<img />") == false {
		t.Error("Does not contain '<img />', Output:" + html)
	}
}
