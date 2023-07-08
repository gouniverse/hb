package hb

// NewFooter is a shortcut to create a new FOOTER tag
func NewFooter() *Tag {
	tag := &Tag{
		TagName: "footer",
	}
	return tag
}
