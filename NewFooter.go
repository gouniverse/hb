package hb

// NewFooter is a shortcut to create a new FOOTER tag
// Deprecated: replaced by the new method Footer()
func NewFooter() *Tag {
	return &Tag{TagName: "footer"}
}
