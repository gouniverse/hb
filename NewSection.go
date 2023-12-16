package hb

// NewSection represents a SECTION tag
func NewSection() *Tag {
	tag := &Tag{
		TagName: "section",
	}
	return tag
}
