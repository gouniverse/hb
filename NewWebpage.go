package hb

// NewWebpage returns a webpage instance
func NewWebpage() *Webpage {
	headTag := &Tag{TagName: "head"}
	bodyTag := &Tag{TagName: "body"}
	h := &Webpage{charset: "utf-8"}
	h.head = headTag
	h.body = bodyTag
	return h
}
