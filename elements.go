package hb

// NewLabel represents a LABEL tag
// Deprecated: replaced by the new method Label()
func NewLabel() *Tag {
	return &Tag{TagName: "label"}
}

// NewLI represents a LI tag
// Deprecated: replaced by the new method LI()
func NewLI() *Tag {
	return &Tag{TagName: "li"}
}

// NewLink represents a LINK tag
// Deprecated: replaced by the new method Link()
func NewLink() *Tag {
	return &Tag{TagName: "link"}
}

// NewMeta represents a META tag
// Deprecated: replaced by the new method Meta()
func NewMeta() *Tag {
	return &Tag{TagName: "meta"}
}

// NewNav represents a NAV tag
// Deprecated: replaced by the new method Nav()
func NewNav() *Tag {
	return &Tag{TagName: "nav"}
}

// NewNavbar represents a NAVBAR tag
// Deprecated: replaced by the new method Navbar()
func NewNavbar() *Tag {
	return &Tag{TagName: "navbar"}
}

// NewOL represents a OL tag
// Deprecated: replaced by the new method OL()
func NewOL() *Tag {
	return &Tag{TagName: "ol"}
}

// NewParagraph represents a IMG tag
// Deprecated: replaced by the new method Paragraph()
func NewParagraph() *Tag {
	return &Tag{TagName: "p"}
}

// NewPRE represents a PRE tag
// Deprecated: replaced by the new method PRE()
func NewPRE() *Tag {
	return &Tag{TagName: "pre"}
}

// NewScriptURL represents a SCRIPT tag with URL
// Deprecated: replaced by the new method ScriptURL()
func NewScriptURL(javascriptURL string) *Tag {
	return &Tag{
		TagName: "script",
		TagAttributes: map[string]string{
			"src": javascriptURL,
		},
	}
}

// NewStyle represents a STYLE tag
// Deprecated: replaced by the new method Style()
func NewStyle(css string) *Tag {
	return &Tag{
		TagName:     "style",
		TagChildren: []TagInterface{NewHTML(css)},
	}
}

// NewStyleURL represents a LINK tag with URL
// Deprecated: replaced by the new method StyleURL()
func NewStyleURL(styleURL string) *Tag {
	return &Tag{
		TagName: "link",
		TagAttributes: map[string]string{
			"href": styleURL,
			"rel":  "stylesheet",
		},
	}
}

// NewSub represents a SUB tag
// Deprecated: replaced by the new method Sub()
func NewSub() *Tag {
	return &Tag{TagName: "sub"}
}

// NewSup represents a SUP tag
// Deprecated: replaced by the new method Sup()
func NewSup() *Tag {
	return &Tag{TagName: "sup"}
}
