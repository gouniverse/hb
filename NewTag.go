package hb

// NewTag creates a tag, with the specified name
// useful for custom tags or ones that are not yet
// added to the hb library
func NewTag(tagName string) *Tag {
	return &Tag{
		TagName: tagName,
	}
}
