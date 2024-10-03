package hb

// NewInput represents a IMG tag
// Deprecated: replaced by the new method Input()
func NewInput() *Tag {
	return &Tag{TagName: "input"}
}
