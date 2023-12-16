package hb

// NewSpan represents a SPAN tag
func NewSpan() *Tag {
	tag := &Tag{
		TagName: "span",
	}
	return tag
}
