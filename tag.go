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

// NewDiv represents a DIV tag
func NewDiv() *Tag {
	section := &Tag{
		TagName: "div",
	}
	return section
}

// NewHTML creates pure HTML without surrounding tags
func NewHTML(html string) *Tag {
	textTag := &Tag{
		TagName:    "",
		TagContent: html,
	}
	return textTag
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
	section := &Tag{
		TagName: "section",
	}
	return section
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
	tagStart := `<` + t.TagName + t.attrToString() + `>`
	tagEnd := `</` + t.TagName + `>`
	if t.TagName == "" {
		tagStart = ""
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
		attrString += ` "` + key + `"="` + addslashes(value) + `"`
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
