package hb

// NewSpan represents a SPAN tag
// Deprecated: replaced by the new method Span()
func NewSpan() *Tag {
	return &Tag{TagName: "span"}
}
