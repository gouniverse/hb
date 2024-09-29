package hb

// NewTextArea represents a FORM tag
// Deprecated: replaced by the new method TextArea()
func NewTextArea() *Tag {
	return &Tag{
		TagName: "textarea",
	}
}
