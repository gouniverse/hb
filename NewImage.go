package hb

// NewImage represents a IMG tag
func NewImage() *Tag {
	tag := &Tag{
		TagName: "img",
	}
	return tag
}
