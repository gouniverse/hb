package html

import (
	"bytes"
	"strings"
)

// Addslashes addslashes()
func addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func inArrayString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// NewButton represents a BUTTON tag
func NewButton() *Tag {
	tag := &Tag{
		TagName: "button",
	}
	return tag
}

// NewDiv represents a DIV tag
func NewDiv() *Tag {
	tag := &Tag{
		TagName: "div",
	}
	return tag
}

// NewForm represents a IMG tag
func NewForm() *Tag {
	tag := &Tag{
		TagName: "input",
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

// NewParagraph represents a IMG tag
func NewParagraph() *Tag {
	tag := &Tag{
		TagName: "p",
	}
	return tag
}

// NewScript represents a STYLE tag
func NewScript(javascript string) *Tag {
	tag := &Tag{
		TagName: "script",
	}
	tag.AddChild(NewHTML(javascript))
	return tag
}

// NewSection represents a SECTION tag
func NewSection() *Tag {
	tag := &Tag{
		TagName: "section",
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

// Tag represents an HTML tag
type Tag struct {
	TagName       string
	TagContent    string
	TagAttributes map[string]string
	TagChildren   []*Tag
}

// Attr shortcut for SetAttribute
func (t *Tag) Attr(key, value string) *Tag {
	return t.SetAttribute(key, value)
}

// AddChild returns HTML from Node
func (t *Tag) AddChild(child *Tag) *Tag {
	t.TagChildren = append(t.TagChildren, child)
	return t
}

// AddChildren returns HTML from Node
func (t *Tag) AddChildren(children []*Tag) *Tag {
	for _, child := range children {
		t.AddChild(child)
	}
	return t
}

// SetAttribute sets the valua of an attribute
func (t *Tag) SetAttribute(key, value string) *Tag {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	t.TagAttributes[key] = value
	return t
}

// GetAttribute returns the value of an attribute
func (t *Tag) GetAttribute(key string) string {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	return t.TagAttributes[key]
}

// ToHTML returns HTML from Node
func (t *Tag) ToHTML() string {
	shortTags := []string{
		"img",
		"input",
		"link",
	}

	isShortTag := inArrayString(shortTags, t.TagName)

	tagStart := `<` + t.TagName + t.attrToString() + `>`
	tagEnd := `</` + t.TagName + `>`
	if t.TagName == "" {
		tagStart = ""
		tagEnd = ""
	}
	if isShortTag {
		tagEnd = ""
	}
	return tagStart + t.TagContent + t.childrenToString() + tagEnd
}

func (t Tag) attrToString() string {
	attrString := ""

	for key, value := range t.TagAttributes {
		if strings.Trim(value, " ") == "" {
			continue
		}
		attrString += ` ` + key + `="` + addslashes(value) + `"`
	}

	if attrString != "" {
		attrString = " " + attrString
	}

	return attrString
}

func (t Tag) childrenToString() string {
	childrenString := ""

	for _, child := range t.TagChildren {
		childrenString += child.ToHTML()
	}

	return childrenString
}
