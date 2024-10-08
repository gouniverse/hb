package hb

import (
	"strings"
	"testing"
)

func TestHtmlWebpage(t *testing.T) {
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

func TestHtmlWebpageSetCharset(t *testing.T) {
	wp := NewWebpage()
	wp.SetCharset("testing")
	wpHtml := wp.ToHTML()

	prefix := `<!DOCTYPE html><html><head><meta charset="testing" />`

	if !strings.HasPrefix(wpHtml, prefix) {
		t.Error("Does not start with: ", prefix, ", Output: ", wpHtml)
	}

	suffix := `</head><body></body></html>`

	if !strings.HasSuffix(wpHtml, suffix) {
		t.Error("Does not end with: ", suffix, ", Output: ", wpHtml)
	}
}

func TestHtmlWebpageSetLanguage(t *testing.T) {
	wp := NewWebpage()
	wp.SetLanguage("en")
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

func TestHtmlWebpageAddScripts(t *testing.T) {
	html := NewWebpage().
		AddScriptURL("http://example.com/example1.js").
		AddScriptURL("http://example.com/example2.js").
		AddScript(`alert("Hello world 1")`).
		AddScript(`alert("Hello world 2")`).
		ToHTML()

	expectedExpressions := []string{
		`<script src="http://example.com/example1.js"></script>`,
		`<script src="http://example.com/example2.js"></script>`,
		`<script>alert("Hello world 1")</script>`,
		`<script>alert("Hello world 2")</script>`,
	}

	for _, expected := range expectedExpressions {
		if !strings.Contains(html, expected) {
			t.Error("Does not contain: ", expected, ", Output: ", html)
		}
	}
}

func TestHtmlWebpageAddStyles(t *testing.T) {
	html := NewWebpage().
		AddStyleURL("http://example.com/example1.css").
		AddStyleURL("http://example.com/example2.css").
		AddStyle(`body { color: red; }`).
		AddStyle(`body { color: blue; }`).
		ToHTML()

	expectedExpressions := []string{
		`<link href="http://example.com/example1.css" rel="stylesheet" />`,
		`<link href="http://example.com/example2.css" rel="stylesheet" />`,
		`<style>body { color: red; }</style>`,
		`<style>body { color: blue; }</style>`,
	}

	for _, expected := range expectedExpressions {
		if !strings.Contains(html, expected) {
			t.Error("Does not contain: ", expected, ", Output: ", html)
		}
	}
}

func TestHtmlWebpageAddMetas(t *testing.T) {
	html := NewWebpage().
		AddMeta(NewMeta().Name("NAME 1").Value("VALUE 1")).
		Meta(NewMeta().Name("NAME 2").Value("VALUE 2")).
		ToHTML()

	expectedExpressions := []string{
		`<meta name="NAME 1" value="VALUE 1" />`,
		`<meta name="NAME 2" value="VALUE 2" />`,
	}

	for _, expected := range expectedExpressions {
		if !strings.Contains(html, expected) {
			t.Error("Does not contain: ", expected, ", Output: ", html)
		}
	}
}

func TestHtmlWebpageHead(t *testing.T) {
	wp := NewWebpage()
	wp.Head().Child(NewScript(`window.onload="alert('Hello')"`))
	wpHtml := wp.ToHTML()

	prefix := `<!DOCTYPE html><html><head><meta charset="utf-8" />`

	if !strings.HasPrefix(wpHtml, prefix) {
		t.Error("Does not start with: ", prefix, ", Output: ", wpHtml)
	}

	suffix := `<script>window.onload="alert('Hello')"</script></head><body></body></html>`

	if !strings.HasSuffix(wpHtml, suffix) {
		t.Error("Does not end with: ", suffix, ", Output: ", wpHtml)
	}
}
