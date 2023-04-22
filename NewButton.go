package hb

// NewButton represents a BUTTON tag
func NewButton() *Tag {
	tag := &Tag{
		TagName: "button",
	}
	return tag
}
