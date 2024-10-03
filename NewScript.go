package hb

// NewScript represents a SCRIPT tag
// Deprecated: replaced by the new method Script()
func NewScript(javascript string) *Tag {
	return &Tag{
		TagName:     "script",
		TagChildren: []TagInterface{NewHTML(javascript)},
	}
}
