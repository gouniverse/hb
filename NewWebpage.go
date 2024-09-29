package hb

// NewWebpage returns a webpage instance
// Deprecated: replaced by the new method Webpage()
func NewWebpage() *HtmlWebpage {
	return &HtmlWebpage{
		charset: "utf-8",
		head:    &Tag{TagName: "head"},
		body:    &Tag{TagName: "body"},
	}
}
