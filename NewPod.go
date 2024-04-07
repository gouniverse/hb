package hb

// NewPod represents a wrapper tag
// It serves as a container for other tags,
// and is used to wrap other tags together.
func NewPod() *Tag {
	tag := &Tag{
		TagName: "",
	}
	return tag
}
