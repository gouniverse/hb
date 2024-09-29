package hb

// NewTemplate represents a TEMPLATE tag
// Deprecated: replaced by the new method Template()
func NewTemplate() *Tag {
	return &Tag{TagName: "template"}
}
