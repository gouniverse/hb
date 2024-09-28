package hb

// NewHTML creates pure HTML without surrounding tags
func NewHTML(html string) *Tag {
	return &Tag{
		TagName:    "",
		TagContent: html,
	}
}
