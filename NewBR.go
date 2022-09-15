package hb

// NewBR represents a BR tag
func NewBR() *Tag {
	tag := &Tag{
		TagName: "br",
	}
	return tag
}
