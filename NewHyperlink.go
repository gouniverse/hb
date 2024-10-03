package hb

// NewHyperlink represents a H1 tag
// Deprecated: replaced by the new method Hyperlink()
func NewHyperlink() *Tag {
	return &Tag{
		TagName: "a",
	}
}
