package hb

// NewCaption is a shortcut to create a new CAPTION tag
// Deprecated: replaced by the new method
func NewCaption() *Tag {
	return &Tag{TagName: "caption"}
}
