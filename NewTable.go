package hb

// NewTable represents a TABLE tag
// Deprecated: replaced by the new method Table()
func NewTable() *Tag {
	return &Tag{TagName: "table"}
}

// NewTR represents a TR tag
// Deprecated: replaced by the new method TR()
func NewTR() *Tag {
	return &Tag{TagName: "tr"}
}

// NewTD represents a TD tag
// Deprecated: replaced by the new method TD()
func NewTD() *Tag {
	return &Tag{TagName: "td"}
}

// NewTH represents a TH tag
// Deprecated: replaced by the new method TH()
func NewTH() *Tag {
	return &Tag{TagName: "th"}
}

// NewTbody represents a TBODY tag
// Deprecated: replaced by the new method Tbody()
func NewTbody() *Tag {
	return &Tag{TagName: "tbody"}
}

// NewThead represents a THEAD tag
// Deprecated: replaced by the new method Thead()
func NewThead() *Tag {
	return &Tag{TagName: "thead"}
}
