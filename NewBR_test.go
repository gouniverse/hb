package hb

import (
	"strings"
	"testing"
)

func TestNewBR(t *testing.T) {
	br := NewBR()
	html := br.ToHTML()

	if strings.Contains(html, "<br />") == false {
		t.Error("Does not contain '<br />', Output:" + html)
	}
}
