package hb

// NewWrap is a convenience tagless container to wrap multiple
// elements together. Any attributes added to the wrap tag will
// be lost. If you need to keep these better use a DIV tag
// Deprecated: replaced by the new method Wrap()
func NewWrap() *Tag {
	return &Tag{TagName: ""}
}
