package hb

import "text/template"

// NewText creates pure escaped text without surrounding tags
func NewText(text string) *Tag {
	escapedText := template.HTMLEscapeString(text)

	return &Tag{
		TagName:    "",
		TagContent: escapedText,
	}
}
