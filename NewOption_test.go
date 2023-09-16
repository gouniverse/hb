package hb

import (
	"strings"
	"testing"
)

func TestNewOption(t *testing.T) {
	e := NewOption()
	html := e.ToHTML()

	if strings.Contains(html, "<option>") == false {
		t.Error("Does not contain '<option>', Output:" + html)
	}

	if strings.Contains(html, "</option>") == false {
		t.Error("Does not contain '</option>', Output:" + html)
	}
}

func TestNewOptionEmptyValue(t *testing.T) {
	e := NewOption().Value("").HTML("- select -")
	html := e.ToHTML()

	if strings.Contains(html, `<option value="">- select -</option>`) == false {
		t.Error(`Does not contain '<option value="">- select -</option>', Output:` + html)
	}
}
