package hb

import (
	"golang.org/x/net/html"
)

// escapeHTMLAttr escapes special characters in a string for safe use in HTML attributes
func escapeHtmlAttr(input string) string {
	return html.EscapeString(input)
}
