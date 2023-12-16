package hb

// NewHTML creates pure HTML without surrounding tags
func NewHTML(html string) *Tag {
	textTag := &Tag{
		TagName:    "",
		TagContent: html,
	}
	return textTag
}
