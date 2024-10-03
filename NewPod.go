package hb

// NewPod represents a wrapper tag
// It serves as a container for other tags,
// and is used to wrap other tags together.
// Deprecated: replaced by the new method Wrap()
func NewPod() *Tag {
	return &Tag{TagName: ""}
}
