package hb

// NewForm represents a FORM tag
// Deprecated: replaced by the new method Form()
func NewForm() *Tag {
	return &Tag{TagName: "form"}
}
