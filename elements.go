package hb

// NewLabel represents a LABEL tag
func NewLabel() *Tag {
	return &Tag{TagName: "label"}
}

// NewLI represents a LI tag
func NewLI() *Tag {
	return &Tag{TagName: "li"}
}

// NewLink represents a LINK tag
func NewLink() *Tag {
	return &Tag{TagName: "link"}
}

// NewMeta represents a META tag
func NewMeta() *Tag {
	return &Tag{TagName: "meta"}
}

// NewNav represents a NAV tag
func NewNav() *Tag {
	return &Tag{TagName: "nav"}
}

// NewNavbar represents a NAVBAR tag
func NewNavbar() *Tag {
	return &Tag{TagName: "navbar"}
}

// NewOL represents a OL tag
func NewOL() *Tag {
	return &Tag{TagName: "ol"}
}

// NewParagraph represents a IMG tag
func NewParagraph() *Tag {
	return &Tag{TagName: "p"}
}

// NewPRE represents a PRE tag
func NewPRE() *Tag {
	return &Tag{TagName: "pre"}
}

// NewScriptURL represents a SCRIPT tag with URL
func NewScriptURL(javascriptURL string) *Tag {
	return &Tag{
		TagName: "script",
		TagAttributes: map[string]string{
			"src": javascriptURL,
		},
	}
}

// NewStyle represents a STYLE tag
func NewStyle(css string) *Tag {
	return &Tag{
		TagName:     "style",
		TagChildren: []TagInterface{NewHTML(css)},
	}
}

// NewStyleURL represents a LINK tag with URL
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
func NewSub() *Tag {
	return &Tag{TagName: "sub"}
}

// NewSup represents a SUP tag
func NewSup() *Tag {
	return &Tag{TagName: "sup"}
}
