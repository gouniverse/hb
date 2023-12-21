package hb

// NewHyperlink represents a H1 tag
func NewHyperlink() *Tag {
	tag := &Tag{
		TagName: "a",
	}
	return tag
}
