package hb

// NewScript represents a SCRIPT tag
func NewScript(javascript string) *Tag {
	tag := &Tag{
		TagName: "script",
	}
	tag.AddChild(NewHTML(javascript))
	return tag
}
