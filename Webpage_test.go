package hb

import (
	"strings"
	"testing"
)

func TestWebpage(t *testing.T) {
	wp := NewWebpage()
	wpHtml := wp.ToHTML()

	prefix := `<!DOCTYPE html><html><head><meta charset="utf-8" />`

	if !strings.HasPrefix(wpHtml, prefix) {
		t.Error("Does not start with: ", prefix, ", Output: ", wpHtml)
	}

	suffix := `</head><body></body></html>`

	if !strings.HasSuffix(wpHtml, suffix) {
		t.Error("Does not end with: ", suffix, ", Output: ", wpHtml)
	}
}

func TestWebpageAddLanguageAttribute(t *testing.T) {
	wp := NewWebpage()
	wp.Attr("lang", "en")
	wpHtml := wp.ToHTML()

	prefix := `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8" />`

	if !strings.HasPrefix(wpHtml, prefix) {
		t.Error("Does not start with: ", prefix, ", Output: ", wpHtml)
	}

	suffix := `</head><body></body></html>`

	if !strings.HasSuffix(wpHtml, suffix) {
		t.Error("Does not end with: ", suffix, ", Output: ", wpHtml)
	}
}
