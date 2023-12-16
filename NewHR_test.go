package hb

import (
	"strings"
	"testing"
)

func TestNewHR(t *testing.T) {
	header := NewHR()
	html := header.ToHTML()

	if strings.Contains(html, "<hr />") == false {
		t.Error("Does not contain '<hr />', Output:" + html)
	}
}
