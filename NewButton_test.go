package hb

import (
	"strings"
	"testing"
)

func TestNewButton(t *testing.T) {
	header := NewButton()
	html := header.ToHTML()

	if strings.Contains(html, "<button>") == false {
		t.Error("Does not contain '<button>', Output:" + html)
	}

	if strings.Contains(html, "</button>") == false {
		t.Error("Does not contain '</button>', Output:" + html)
	}
}
