package hb

// NewTag represents a tag
func NewTag(tagName string) *Tag {
	tag := &Tag{
		TagName: tagName,
	}
	return tag
}

// NewForm represents a FORM tag
func NewForm() *Tag {
	tag := &Tag{
		TagName: "form",
	}
	return tag
}

// NewHTML creates pure HTML without surrounding tags
func NewHTML(html string) *Tag {
	textTag := &Tag{
		TagName:    "",
		TagContent: html,
	}
	return textTag
}

// NewHR creates represents a HR tags
func NewHR() *Tag {
	textTag := &Tag{
		TagName: "hr",
	}
	return textTag
}

// NewHeading1 represents a H1 tag
func NewHeading1() *Tag {
	tag := &Tag{
		TagName: "h1",
	}
	return tag
}

// NewHeading2 represents a H1 tag
func NewHeading2() *Tag {
	tag := &Tag{
		TagName: "h2",
	}
	return tag
}

// NewHeading3 represents a H1 tag
func NewHeading3() *Tag {
	tag := &Tag{
		TagName: "h3",
	}
	return tag
}

// NewHeading4 represents a H1 tag
func NewHeading4() *Tag {
	tag := &Tag{
		TagName: "h4",
	}
	return tag
}

// NewHeading5 represents a H1 tag
func NewHeading5() *Tag {
	tag := &Tag{
		TagName: "h5",
	}
	return tag
}

// NewHeading6 represents a H1 tag
func NewHeading6() *Tag {
	tag := &Tag{
		TagName: "h6",
	}
	return tag
}

// NewHyperlink represents a H1 tag
func NewHyperlink() *Tag {
	tag := &Tag{
		TagName: "a",
	}
	return tag
}

// NewImage represents a IMG tag
func NewImage() *Tag {
	tag := &Tag{
		TagName: "img",
	}
	return tag
}

// NewInput represents a IMG tag
func NewInput() *Tag {
	tag := &Tag{
		TagName: "input",
	}
	return tag
}

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

// NewOption represents an OPTION tag
func NewOption() *Tag {
	tag := &Tag{
		TagName: "option",
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

// NewScript represents a SCRIPT tag
func NewScript(javascript string) *Tag {
	tag := &Tag{
		TagName: "script",
	}
	tag.AddChild(NewHTML(javascript))
	return tag
}

// NewScriptURL represents a SCRIPT tag with URL
func NewScriptURL(javascriptURL string) *Tag {
	tag := &Tag{
		TagName: "script",
	}
	tag.SetAttribute("src", javascriptURL)
	return tag
}

// NewSection represents a SECTION tag
func NewSection() *Tag {
	tag := &Tag{
		TagName: "section",
	}
	return tag
}

// NewSelect represents a SELECT tag
func NewSelect() *Tag {
	tag := &Tag{
		TagName: "select",
	}
	return tag
}

// NewSpan represents a SPAN tag
func NewSpan() *Tag {
	tag := &Tag{
		TagName: "span",
	}
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

// NewTemplate represents a TEMPLATE tag
func NewTemplate() *Tag {
	tag := &Tag{
		TagName: "template",
	}
	return tag
}

// NewTextArea represents a FORM tag
func NewTextArea() *Tag {
	tag := &Tag{
		TagName: "textarea",
	}
	return tag
}

// NewUL represents a UL tag
func NewUL() *Tag {
	tag := &Tag{
		TagName: "ul",
	}
	return tag
}

// NewWebpage returns a webpage instance
func NewWebpage() *Webpage {
	headTag := &Tag{TagName: "head"}
	bodyTag := &Tag{TagName: "body"}
	h := &Webpage{Charset: "utf-8"}
	h.Head = headTag
	h.Body = bodyTag
	return h
}
