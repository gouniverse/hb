package hb

// NewHeader is a shortcut to create a new HEADER tag
func NewHeader() *Tag {
	tag := &Tag{
		TagName: "header",
	}
	return tag
}
