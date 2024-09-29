package hb

// NewUL represents a UL tag
// Deprecated: replaced by the new method UL()
func NewUL() *Tag {
	return &Tag{
		TagName: "ul",
	}
}
