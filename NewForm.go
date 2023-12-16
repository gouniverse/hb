package hb

// NewForm represents a FORM tag
func NewForm() *Tag {
	tag := &Tag{
		TagName: "form",
	}
	return tag
}
