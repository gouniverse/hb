package hb

// NewTable represents a TABLE tag
func NewTable() *Tag {
	tag := &Tag{
		TagName: "table",
	}
	return tag
}

// NewTR represents a TR tag
func NewTR() *Tag {
	tag := &Tag{
		TagName: "tr",
	}
	return tag
}

// NewTD represents a TD tag
func NewTD() *Tag {
	tag := &Tag{
		TagName: "td",
	}
	return tag
}

// NewTH represents a TH tag
func NewTH() *Tag {
	tag := &Tag{
		TagName: "th",
	}
	return tag
}

// NewTbody represents a TBODY tag
func NewTbody() *Tag {
	tag := &Tag{
		TagName: "tbody",
	}
	return tag
}

// NewThead represents a THEAD tag
func NewThead() *Tag {
	tag := &Tag{
		TagName: "thead",
	}
	return tag
}
