package hb

// NewImage represents a IMG tag
// Deprecated: replaced by the new method Image()
func NewImage() *Tag {
	return &Tag{TagName: "img"}
}
