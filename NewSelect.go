package hb

// NewSelect represents a SELECT tag
// Deprecated: replaced by the new method Select()
func NewSelect() *Tag {
	return &Tag{TagName: "select"}
}
