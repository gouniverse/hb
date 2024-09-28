package hb

// NewScript represents a SCRIPT tag
func NewScript(javascript string) *Tag {
	return &Tag{
		TagName:     "script",
		TagChildren: []TagInterface{NewHTML(javascript)},
	}
}
