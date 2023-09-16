package hb

// NewOption represents an OPTION tag
func NewOption() *Tag {
	tag := &Tag{
		TagName: "option",
	}
	return tag
}
