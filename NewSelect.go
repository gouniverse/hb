package hb

// NewSelect represents a SELECT tag
func NewSelect() *Tag {
	tag := &Tag{
		TagName: "select",
	}
	return tag
}
