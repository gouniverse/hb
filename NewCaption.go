package hb

// NewCaption is a shortcut to create a new CAPTION tag
func NewCaption() *Tag {
	tag := &Tag{
		TagName: "caption",
	}
	return tag
}
