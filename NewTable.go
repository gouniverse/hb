package hb

// NewTable represents a TABLE tag
// Deprecated: replaced by the new method Table()
func NewTable() *Tag {
	return &Tag{TagName: "table"}
}

// NewTR represents a TR tag
func NewTR() *Tag {
	return &Tag{TagName: "tr"}
}

// NewTD represents a TD tag
func NewTD() *Tag {
	return &Tag{TagName: "td"}
}

// NewTH represents a TH tag
func NewTH() *Tag {
	return &Tag{TagName: "th"}
}

// NewTbody represents a TBODY tag
func NewTbody() *Tag {
	return &Tag{TagName: "tbody"}
}

// NewThead represents a THEAD tag
func NewThead() *Tag {
	return &Tag{TagName: "thead"}
}
