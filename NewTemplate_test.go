package hb

import (
	"strings"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	s := NewTemplate()
	html := s.ToHTML()

	if strings.Contains(html, "<template>") == false {
		t.Error("Does not contain '<template>', Output:" + html)
	}

	if strings.Contains(html, "</template>") == false {
		t.Error("Does not contain '</template>', Output:" + html)
	}
}
