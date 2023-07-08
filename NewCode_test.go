package hb

import (
	"strings"
	"testing"
)

func TestNewCode(t *testing.T) {
	code := NewCode()
	html := code.ToHTML()

	if strings.Contains(html, "<code>") == false {
		t.Error("Does not contain '<code>', Output:" + html)
	}

	if strings.Contains(html, "</code>") == false {
		t.Error("Does not contain '</code>', Output:" + html)
	}
}