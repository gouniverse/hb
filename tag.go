package hb

import (
	"bytes"
	"sort"
	"strings"
)

// Addslashes addslashes()
func addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		//case '\'', '"', '\\':
		case '"', '\\':
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

// TagInterface represents an HTML tag interface
type TagInterface interface {
	ToHTML() string
}

// Tag represents an HTML tag
type Tag struct {
	TagInterface
	TagName       string
	TagContent    string
	TagAttributes map[string]string
	TagChildren   []*Tag
}

// HasClass returns true if the item is specified with the specified name.
func (n Node) HasClass(className string) (ok bool) {
func (t *Tag) HadClass(className string) bool
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames," ")
	return inArrayString(classNamesArray, className)
}

// AddClass adds a new class name to the tag attribute list.
func (t *Tag) AddClass(className string) *Tag {
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames," ")
	classNamesArray := append(classNamesArray, className)
	classNames = strings.Join(classNamesArray, " ")
	return t.SetAttribute("class", classNames)
}

// Attr shortcut for SetAttribute
func (t *Tag) Attr(key, value string) *Tag {
	return t.SetAttribute(key, value)
}

// Attrs shortcut for setting multiple attributes
func (t *Tag) Attrs(attrs map[string]string) *Tag {
	for key, value := range attrs {
		t.SetAttribute(key, value)
	}
	return t
}

// HTML shortcut for AddHTML
func (t *Tag) HTML(html string) *Tag {
	return t.AddHTML(html)
}

// AddChild adds a new child tag to this tag
func (t *Tag) AddChild(child *Tag) *Tag {
	t.TagChildren = append(t.TagChildren, child)
	return t
}

// AddChildren adds an array of child tags to this tag
func (t *Tag) AddChildren(children []*Tag) *Tag {
	for _, child := range children {
		t.AddChild(child)
	}
	return t
}

// AddHTML adds an HTML as child tags to this tag
func (t *Tag) AddHTML(html string) *Tag {
	t.AddChild(NewHTML(html))
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
		"meta",
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

	keys := make([]string, 0, len(t.TagAttributes))
	for k := range t.TagAttributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := t.TagAttributes[key]
		//for key, value := range t.TagAttributes {
		if strings.Trim(value, " ") == "" {
			continue
		}
		attrString += ` ` + key + `="` + addslashes(value) + `"`
	}

	if attrString != "" {
		attrString = " " + attrString
	}

	if len(attrString) < 1 {
		return ""
	}

	return " " + strings.Trim(attrString, " ")
}

func (t Tag) childrenToString() string {
	childrenString := ""

	for _, child := range t.TagChildren {
		childrenString += child.ToHTML()
	}

	return childrenString
}
