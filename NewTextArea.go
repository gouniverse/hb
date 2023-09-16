package hb

// NewTextArea represents a FORM tag
func NewTextArea() *Tag {
	tag := &Tag{
		TagName: "textarea",
	}
	return tag
}
