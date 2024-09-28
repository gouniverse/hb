package hb

// NewWebpage returns a webpage instance
func NewWebpage() *Webpage {
	return &Webpage{
		charset: "utf-8",
		head:    &Tag{TagName: "head"},
		body:    &Tag{TagName: "body"},
	}
}
