package hb

// NewSection represents a SECTION tag
// Deprecated: replaced by the new method Section()
func NewSection() *Tag {
	return &Tag{TagName: "section"}
}
