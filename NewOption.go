package hb

// NewOption represents an OPTION tag
// Deprecated: replaced by the new method Option()
func NewOption() *Tag {
	return &Tag{TagName: "option"}
}
