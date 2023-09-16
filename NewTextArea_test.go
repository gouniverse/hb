package hb

import (
	"strings"
	"testing"
)

func TestNewTextArea(t *testing.T) {
	e := NewTextArea()
	html := e.ToHTML()

	if strings.Contains(html, "<textarea>") == false {
		t.Error("Does not contain '<textarea>', Output:" + html)
	}

	if strings.Contains(html, "</textarea>") == false {
		t.Error("Does not contain '</textarea>', Output:" + html)
	}
}
