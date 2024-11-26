package hb

import "golang.org/x/net/html"

// Map a slice of anything to a slice of tags.
func Map[T any](items []T, callback func(item T, index int) TagInterface) []TagInterface {
	tags := make([]TagInterface, len(items))
	for i, item := range items {
		tags[i] = callback(item, i)
	}
	return tags
}

// escapeHTMLAttr escapes special characters in a string for safe use in HTML attributes
func escapeHtmlAttr(input string) string {
	return html.EscapeString(input)
}

var shortTags = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

func isShortTag(tagName string) bool {
	return shortTags[tagName]
}
