package hb

// NewLabel represents a LABEL tag
func NewLabel() *Tag {
	tag := &Tag{TagName: "label"}
	return tag
}

// NewLI represents a LI tag
func NewLI() *Tag {
	tag := &Tag{TagName: "li"}
	return tag
}

// NewLink represents a LINK tag
func NewLink() *Tag {
	tag := &Tag{TagName: "link"}
	return tag
}

// NewMeta represents a META tag
func NewMeta() *Tag {
	tag := &Tag{TagName: "meta"}
	return tag
}

// NewNav represents a NAV tag
func NewNav() *Tag {
	tag := &Tag{
		TagName: "nav",
	}
	return tag
}

// NewNavbar represents a NAVBAR tag
func NewNavbar() *Tag {
	tag := &Tag{
		TagName: "navbar",
	}
	return tag
}

// NewOL represents a UL tag
func NewOL() *Tag {
	tag := &Tag{
		TagName: "ol",
	}
	return tag
}

// NewParagraph represents a IMG tag
func NewParagraph() *Tag {
	tag := &Tag{
		TagName: "p",
	}
	return tag
}

// NewPRE represents a PRE tag
func NewPRE() *Tag {
	return NewTag("pre")
}

// NewScriptURL represents a SCRIPT tag with URL
func NewScriptURL(javascriptURL string) *Tag {
	tag := &Tag{
		TagName: "script",
	}
	tag.SetAttribute("src", javascriptURL)
	return tag
}

// NewStyle represents a STYLE tag
func NewStyle(css string) *Tag {
	tag := &Tag{
		TagName: "style",
	}
	tag.AddChild(NewHTML(css))
	return tag
}

// NewStyleURL represents a LINK tag with URL
func NewStyleURL(styleURL string) *Tag {
	tag := NewLink()
	tag.SetAttribute("href", styleURL)
	tag.SetAttribute("rel", "stylesheet")
	return tag
}

// NewSub represents a SUB tag
func NewSub() *Tag {
	tag := &Tag{
		TagName: "sub",
	}
	return tag
}

// NewSup represents a SUP tag
func NewSup() *Tag {
	tag := &Tag{
		TagName: "sup",
	}
	return tag
}
