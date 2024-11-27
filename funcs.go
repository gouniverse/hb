package hb

import "golang.org/x/net/html"

// ToTags is a functional helper to convert a slice of items
// to a slice of tags using a callback function
//
// Example:
//
// items := []string{"one", "two", "three"}
// tags := ToTags(items, func(item string, index int) TagInterface {
//   return NewSpan().Text(item)
// })
//
// Output: <span>one</span><span>two</span><span>three</span>
//
// Parameters:
// - items: The slice of items to convert to tags
// - callback: A function that takes an item and index and returns a TagInterface
//
// Returns:
// - []TagInterface: A slice of tags
func ToTags[T any](items []T, callback func(item T, index int) TagInterface) []TagInterface {
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
