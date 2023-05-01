package hb

import (
	"bytes"
	"sort"
	"strings"

	"golang.org/x/net/html"
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

// Action shortcut for setting the "action" attribute
func (t *Tag) Action(action string) *Tag {
	return t.SetAttribute("action", action)
}

// AddClass adds a new class name to the tag attribute list.
func (t *Tag) AddClass(className string) *Tag {
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames, " ")
	classNamesArray = append(classNamesArray, className)
	classNames = strings.Join(classNamesArray, " ")
	classNames = strings.Trim(classNames, " ")
	return t.SetAttribute("class", classNames)
}

// AddStyle adds a new class name to the tag attribute list.
func (t *Tag) AddStyle(style string) *Tag {
	styles := t.GetAttribute("style")
	if styles == "" {
		styles = style
	} else if strings.HasSuffix(styles, ";") {
		styles = styles + style
	} else {
		styles = styles + ";" + style
	}
	return t.SetAttribute("style", styles)
}

// Attr shortcut for SetAttribute
func (t *Tag) Attr(key, value string) *Tag {
	return t.SetAttribute(key, value)
}

// AttrIf shortcut for setting an attribute if a condition is met
func (t *Tag) AttrIf(condition bool, key, value string) *Tag {
	if condition {
		return t.SetAttribute(key, value)
	}
	return t
}

// AttrIfElse shortcut for setting an attribute if a condition is met, otherwise adds another attribute
func (t *Tag) AttrIfElse(condition bool, key, valueIf string, valueElse string) *Tag {
	if condition {
		return t.SetAttribute(key, valueIf)
	}

	return t.SetAttribute(key, valueElse)
}

// Attrs shortcut for setting multiple attributes
func (t *Tag) Attrs(attrs map[string]string) *Tag {
	for key, value := range attrs {
		t.SetAttribute(key, value)
	}
	return t
}

// AttrsIf shortcut for setting multiple attributes if a condition is met
func (t *Tag) AttrsIf(condition bool, attrs map[string]string) *Tag {
	if condition {
		for key, value := range attrs {
			t.SetAttribute(key, value)
		}
	}
	return t
}

// AttrsIfElse shortcut for setting multiple attributes if a condition is met, otherwise adds another attribute
func (t *Tag) AttrsIfElse(condition bool, attrsIf map[string]string, attrsElse map[string]string) *Tag {
	if condition {
		for key, value := range attrsIf {
			t.SetAttribute(key, value)
		}
	} else {
		for key, value := range attrsElse {
			t.SetAttribute(key, value)
		}
	}
	return t
}

// AddChild adds a new child tag to this tag
func (t *Tag) AddChild(child *Tag) *Tag {
	if child == nil {
		return t
	}

	t.TagChildren = append(t.TagChildren, child)
	return t
}

// AddChildren adds an array of child tags to this tag
func (t *Tag) AddChildren(children []*Tag) *Tag {
	for _, child := range children {
		if child == nil {
			continue
		}

		t.AddChild(child)
	}
	return t
}

// AddHTML adds an HTML as child tags to this tag
func (t *Tag) AddHTML(html string) *Tag {
	t.AddChild(NewHTML(html))
	return t
}

// Childs shortcut for AddChild
func (t *Tag) Child(child *Tag) *Tag {
	return t.AddChild(child)
}

// ChildIf adds a child if a condition is met
func (t *Tag) ChildIf(condition bool, child *Tag) *Tag {
	if condition {
		return t.AddChild(child)
	}

	return t
}

// ChildIfElse adds a child if a condition is met, otherwise adds another child
func (t *Tag) ChildIfElse(condition bool, childIf *Tag, childElse *Tag) *Tag {
	if condition {
		return t.AddChild(childIf)
	}

	return t.AddChild(childElse)
}

// Children shortcut for AddChildren
func (t *Tag) Children(children []*Tag) *Tag {
	return t.AddChildren(children)
}

// ChildrenIf adds children if a condition is met
func (t *Tag) ChildrenIf(condition bool, children []*Tag) *Tag {
	if condition {
		return t.AddChildren(children)
	}

	return t
}

// ChildrenIfElse adds children if a condition is met
func (t *Tag) ChildrenIfElse(condition bool, childrenIf []*Tag, childrenElse []*Tag) *Tag {
	if condition {
		return t.AddChildren(childrenIf)
	}

	return t.AddChildren(childrenElse)
}

// Class shortcut for setting the "class" attribute
func (t *Tag) Class(clasName string) *Tag {
	return t.AddClass(clasName)
}

// ClassIf adds class name if a condition is met
func (t *Tag) ClassIf(condition bool, clasName string) *Tag {
	if condition {
		return t.AddClass(clasName)
	}

	return t
}

// ClassIfElse adds class name if a condition is met
func (t *Tag) ClassIfElse(condition bool, clasNameIf string, classNameElse string) *Tag {
	if condition {
		return t.AddClass(clasNameIf)
	}

	return t.AddClass(classNameElse)
}

// Data shortcut for setting a "data-" attribute
func (t *Tag) Data(name string, value string) *Tag {
	return t.SetAttribute("data-"+name, value)
}

// DataIf shortcut for setting a "data-" attribute if a condition is met
func (t *Tag) DataIf(condition bool, name string, value string) *Tag {
	if condition {
		return t.SetAttribute("data-"+name, value)
	}
	return t
}

// DataIfElse shortcut for setting a "data-" attribute if a condition is met
func (t *Tag) DataIfElse(condition bool, name string, valueIf string, valueElse string) *Tag {
	if condition {
		return t.SetAttribute("data-"+name, valueIf)
	}
	return t.SetAttribute("data-"+name, valueElse)
}

// Enctype shortcut for setting the "enctype" attribute
func (t *Tag) Enctype(enctype string) *Tag {
	return t.SetAttribute("enctype", enctype)
}

// GetAttribute returns the value of an attribute
func (t *Tag) GetAttribute(key string) string {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	return t.TagAttributes[key]
}

// HasClass returns true if the tag has a class with the specified name.
func (t *Tag) HasClass(className string) bool {
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames, " ")
	return inArrayString(classNamesArray, className)
}

// HTML shortcut for AddHTML
func (t *Tag) HTML(html string) *Tag {
	return t.AddHTML(html)
}

// HTMLIf adds html if a condition is met
func (t *Tag) HTMLIf(condition bool, html string) *Tag {
	if condition {
		return t.AddHTML(html)
	}

	return t
}

// HTMLIfElse adds html if a condition is met
func (t *Tag) HTMLIfElse(condition bool, htmlIf string, htmlElse string) *Tag {
	if condition {
		return t.AddHTML(htmlIf)
	}

	return t.AddHTML(htmlElse)
}

// Href shortcut for setting the "href" attribute
func (t *Tag) Href(href string) *Tag {
	return t.SetAttribute("href", href)
}

// ID shortcut for setting the "id" attribute
func (t *Tag) ID(id string) *Tag {
	return t.SetAttribute("id", id)
}

// Method shortcut for setting the "method" attribute
func (t *Tag) Method(method string) *Tag {
	return t.SetAttribute("method", method)
}

// Name shortcut for setting the "name" attribute
func (t *Tag) Name(name string) *Tag {
	return t.SetAttribute("name", name)
}

// Placeholder shortcut for setting the "placeholder" attribute
func (t *Tag) Placeholder(placeholder string) *Tag {
	return t.SetAttribute("placeholder", placeholder)
}

// OnBlur shortcut for setting the "onblur" attribute
func (t *Tag) OnBlur(js string) *Tag {
	return t.SetAttribute("onblur", html.EscapeString(js))
}

// OnChange shortcut for setting the "onchange" attribute
func (t *Tag) OnChange(js string) *Tag {
	return t.SetAttribute("onchange", html.EscapeString(js))
}

// OnClick shortcut for setting the "onclick" attribute
func (t *Tag) OnClick(js string) *Tag {
	return t.SetAttribute("onclick", html.EscapeString(js))
}

// OnDblClick shortcut for setting the "ondblclick" attribute
func (t *Tag) OnDblClick(js string) *Tag {
	return t.SetAttribute("ondblclick", html.EscapeString(js))
}

// OnFocus shortcut for setting the "onfocus" attribute
func (t *Tag) OnFocus(js string) *Tag {
	return t.SetAttribute("onfocus", html.EscapeString(js))
}

// OnInput shortcut for setting the "oninput" attribute
func (t *Tag) OnInput(js string) *Tag {
	return t.SetAttribute("oninput", html.EscapeString(js))
}

// OnKeyDown shortcut for setting the "onkeydown" attribute
func (t *Tag) OnKeyDown(js string) *Tag {
	return t.SetAttribute("onkeydown", html.EscapeString(js))
}

// OnKeyPress shortcut for setting the "onkeydown" attribute
func (t *Tag) OnKeyPress(js string) *Tag {
	return t.SetAttribute("onkeypress", html.EscapeString(js))
}

// OnKeyUp shortcut for setting the "onkeyup" attribute
func (t *Tag) OnKeyUp(js string) *Tag {
	return t.SetAttribute("onkeyup", html.EscapeString(js))
}

// OnLoad shortcut for setting the "onload" attribute
func (t *Tag) OnLoad(js string) *Tag {
	return t.SetAttribute("onload", html.EscapeString(js))
}

// OnMouseDown shortcut for setting the "onmousedown" attribute
func (t *Tag) OnMouseDown(js string) *Tag {
	return t.SetAttribute("onmousedown", html.EscapeString(js))
}

// OnMouseEnter shortcut for setting the "onmouseenter" attribute
func (t *Tag) OnMouseEnter(js string) *Tag {
	return t.SetAttribute("onmouseenter", html.EscapeString(js))
}

// OnMouseLeave shortcut for setting the "onmouseleave" attribute
func (t *Tag) OnMouseLeave(js string) *Tag {
	return t.SetAttribute("onmouseleave", html.EscapeString(js))
}

// OnMouseMove shortcut for setting the "onmousemove" attribute
func (t *Tag) OnMouseMove(js string) *Tag {
	return t.SetAttribute("onmousemove", html.EscapeString(js))
}

// OnMouseOut shortcut for setting the "onmouseout" attribute
func (t *Tag) OnMouseOut(js string) *Tag {
	return t.SetAttribute("onmouseout", html.EscapeString(js))
}

// OnMouseOver shortcut for setting the "onmouseover" attribute
func (t *Tag) OnMouseOver(js string) *Tag {
	return t.SetAttribute("onmouseover", html.EscapeString(js))
}

// OnMouseUp shortcut for setting the "onmouseup" attribute
func (t *Tag) OnMouseUp(js string) *Tag {
	return t.SetAttribute("onmouseup", html.EscapeString(js))
}

// OnSubmit shortcut for setting the "onsubmit" attribute
func (t *Tag) OnSubmit(js string) *Tag {
	return t.SetAttribute("onsubmit", html.EscapeString(js))
}

// SetAttribute sets the valua of an attribute
func (t *Tag) SetAttribute(key, value string) *Tag {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	t.TagAttributes[key] = value
	return t
}

// Src shortcut for setting the "src" attribute
func (t *Tag) Src(src string) *Tag {
	return t.SetAttribute("src", src)
}

// Style shortcut for setting the "style" attribute
func (t *Tag) Style(style string) *Tag {
	return t.AddStyle(style)
}

// StyleIf adds style if a condition is met
func (t *Tag) StyleIf(condition bool, style string) *Tag {
	if condition {
		return t.AddStyle(style)
	}

	return t
}

// StyleIfElse adds style if a condition is met
func (t *Tag) StyleIfElse(condition bool, styleIf string, styleElse string) *Tag {
	if condition {
		return t.AddStyle(styleIf)
	}

	return t.AddStyle(styleElse)
}

// Target shortcut for setting the "target" attribute
func (t *Tag) Target(target string) *Tag {
	return t.SetAttribute("target", target)
}

// ToHTML returns HTML from Node
func (t *Tag) ToHTML() string {
	shortTags := []string{
		"br",
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
		tagStart = `<` + t.TagName + t.attrToString() + ` />`
		tagEnd = ""
	}

	var builder strings.Builder
	builder.WriteString(tagStart)
	builder.WriteString(t.TagContent)
	builder.WriteString(t.childrenToString())
	builder.WriteString(tagEnd)
	return builder.String()
}

// Type shortcut for setting the "type" attribute
func (t *Tag) Type(inputType string) *Tag {
	return t.SetAttribute("type", inputType)
}

// Value shortcut for setting the "value" attribute
func (t *Tag) Value(value string) *Tag {
	return t.SetAttribute("value", value)
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
