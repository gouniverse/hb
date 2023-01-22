package hb

import (
	"strings"
	"testing"
)

func TestNewFooter(t *testing.T) {
	footer := NewFooter()
	html := footer.ToHTML()

	if strings.Contains(html, "<footer>") == false {
		t.Error("Does not contain '<footer>', Output:" + html)
	}

	if strings.Contains(html, "</footer>") == false {
		t.Error("Does not contain '</footer>', Output:" + html)
	}
}
