package hb

// NewHTML creates pure HTML without surrounding tags
// Deprecated: replaced by the new method Raw()
func NewHTML(html string) *Tag {
	return &Tag{
		TagName:    "",
		TagContent: html,
	}
}
