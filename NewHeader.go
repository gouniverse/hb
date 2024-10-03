package hb

// NewHeader is a shortcut to create a new HEADER tag
// Deprecated: replaced by the new method Header()
func NewHeader() *Tag {
	return &Tag{TagName: "header"}
}
