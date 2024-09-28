package hb

import (
	"bytes"
	"slices"
	"sort"
	"strings"
)

var _ TagInterface = (*Tag)(nil)

// Tag represents an HTML tag
type Tag struct {
	TagInterface
	TagName       string
	TagContent    string
	TagAttributes map[string]string
	TagChildren   []TagInterface
}

// Action shortcut for setting the "action" attribute
func (t *Tag) Action(action string) *Tag {
	return t.SetAttribute("action", action)
}

// AddClass adds a new class name to the tag attribute list.
func (t *Tag) AddClass(className string) *Tag {
	classNamesArray := strings.Split(t.GetAttribute("class"), " ")
	if !slices.Contains(classNamesArray, className) {
		classNamesArray = append(classNamesArray, className)
	}
	return t.SetAttribute("class", strings.Trim(strings.Join(classNamesArray, " "), " "))
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

// Alt shortcut for setting the "alt" attribute
func (t *Tag) Alt(text string) *Tag {
	return t.SetAttribute("alt", text)
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

func (t *Tag) AttrIfF(condition bool, key string, valueFunc func() string) *Tag {
	if condition {
		return t.SetAttribute(key, valueFunc())
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

// AttrsIf shortcut for setting multiple attributes via function if a condition is met
func (t *Tag) AttrsIfF(condition bool, attrsFunc func() map[string]string) *Tag {
	if condition {
		for key, value := range attrsFunc() {
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
func (t *Tag) AddChild(child TagInterface) *Tag {
	if child == nil {
		return t
	}

	t.TagChildren = append(t.TagChildren, child)
	return t
}

// AddChildren adds an array of child tags to this tag
func (t *Tag) AddChildren(children []TagInterface) *Tag {
	for _, child := range children {
		if child == nil {
			continue
		}

		t.AddChild(child)
	}
	return t
}

// AddHTML adds an HTML string as a child tag to this tag
func (t *Tag) AddHTML(html string) *Tag {
	t.AddChild(NewHTML(html))
	return t
}

// AddText adds an escaped text string as a child tag to this tag
func (t *Tag) AddText(text string) *Tag {
	t.AddChild(NewText(text))
	return t
}

// Child shortcut for AddChild
func (t *Tag) Child(child TagInterface) *Tag {
	return t.AddChild(child)
}

// ChildIf adds a child if a condition is met
func (t *Tag) ChildIf(condition bool, child TagInterface) *Tag {
	if condition {
		return t.AddChild(child)
	}

	return t
}

// ChildIf adds a child using function if a condition is met
func (t *Tag) ChildIfF(condition bool, childFunc func() TagInterface) *Tag {
	if condition {
		return t.AddChild(childFunc())
	}

	return t
}

// ChildIfElse adds a child if a condition is met, otherwise adds another child
func (t *Tag) ChildIfElse(condition bool, childIf TagInterface, childElse TagInterface) *Tag {
	if condition {
		return t.AddChild(childIf)
	}

	return t.AddChild(childElse)
}

// Children shortcut for AddChildren
func (t *Tag) Children(children []TagInterface) *Tag {
	return t.AddChildren(children)
}

// ChildrenIf adds children if a condition is met
func (t *Tag) ChildrenIf(condition bool, children []TagInterface) *Tag {
	if condition {
		return t.AddChildren(children)
	}

	return t
}

// ChildrenIf adds children using function if a condition is met
func (t *Tag) ChildrenIfF(condition bool, childrenFunc func() []TagInterface) *Tag {
	if condition {
		return t.AddChildren(childrenFunc())
	}

	return t
}

// ChildrenIfElse adds children if a condition is met
func (t *Tag) ChildrenIfElse(condition bool, childrenIf []TagInterface, childrenElse []TagInterface) *Tag {
	if condition {
		return t.AddChildren(childrenIf)
	}

	return t.AddChildren(childrenElse)
}

// ChildrenMap map a slice to a slice of tags and adds as children
func (t *Tag) ChildrenMap(items []any, callback func(item any, index int) TagInterface) *Tag {
	mappedItems := Map(items, callback)
	return t.Children(mappedItems)
}

// Class shortcut for setting the "class" attribute
func (t *Tag) Class(className string) *Tag {
	return t.AddClass(className)
}

// ClassIf adds class name if a condition is met
func (t *Tag) ClassIf(condition bool, className string) *Tag {
	if condition {
		return t.AddClass(className)
	}

	return t
}

// ClassIfElse adds class name if a condition is met
func (t *Tag) ClassIfElse(condition bool, classNameIf string, classNameElse string) *Tag {
	if condition {
		return t.AddClass(classNameIf)
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
	return slices.Contains(classNamesArray, className)
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
	return t.SetAttribute("onblur", js)
}

// OnChange shortcut for setting the "onchange" attribute
func (t *Tag) OnChange(js string) *Tag {
	return t.SetAttribute("onchange", js)
}

// OnClick shortcut for setting the "onclick" attribute
func (t *Tag) OnClick(js string) *Tag {
	return t.SetAttribute("onclick", js)
}

// OnDblClick shortcut for setting the "ondblclick" attribute
func (t *Tag) OnDblClick(js string) *Tag {
	return t.SetAttribute("ondblclick", js)
}

// OnFocus shortcut for setting the "onfocus" attribute
func (t *Tag) OnFocus(js string) *Tag {
	return t.SetAttribute("onfocus", js)
}

// OnInput shortcut for setting the "oninput" attribute
func (t *Tag) OnInput(js string) *Tag {
	return t.SetAttribute("oninput", js)
}

// OnKeyDown shortcut for setting the "onkeydown" attribute
func (t *Tag) OnKeyDown(js string) *Tag {
	return t.SetAttribute("onkeydown", js)
}

// OnKeyPress shortcut for setting the "onkeydown" attribute
func (t *Tag) OnKeyPress(js string) *Tag {
	return t.SetAttribute("onkeypress", js)
}

// OnKeyUp shortcut for setting the "onkeyup" attribute
func (t *Tag) OnKeyUp(js string) *Tag {
	return t.SetAttribute("onkeyup", js)
}

// OnLoad shortcut for setting the "onload" attribute
func (t *Tag) OnLoad(js string) *Tag {
	return t.SetAttribute("onload", js)
}

// OnMouseDown shortcut for setting the "onmousedown" attribute
func (t *Tag) OnMouseDown(js string) *Tag {
	return t.SetAttribute("onmousedown", js)
}

// OnMouseEnter shortcut for setting the "onmouseenter" attribute
func (t *Tag) OnMouseEnter(js string) *Tag {
	return t.SetAttribute("onmouseenter", js)
}

// OnMouseLeave shortcut for setting the "onmouseleave" attribute
func (t *Tag) OnMouseLeave(js string) *Tag {
	return t.SetAttribute("onmouseleave", js)
}

// OnMouseMove shortcut for setting the "onmousemove" attribute
func (t *Tag) OnMouseMove(js string) *Tag {
	return t.SetAttribute("onmousemove", js)
}

// OnMouseOut shortcut for setting the "onmouseout" attribute
func (t *Tag) OnMouseOut(js string) *Tag {
	return t.SetAttribute("onmouseout", js)
}

// OnMouseOver shortcut for setting the "onmouseover" attribute
func (t *Tag) OnMouseOver(js string) *Tag {
	return t.SetAttribute("onmouseover", js)
}

// OnMouseUp shortcut for setting the "onmouseup" attribute
func (t *Tag) OnMouseUp(js string) *Tag {
	return t.SetAttribute("onmouseup", js)
}

// OnSubmit shortcut for setting the "onsubmit" attribute
func (t *Tag) OnSubmit(js string) *Tag {
	return t.SetAttribute("onsubmit", js)
}

// Src shortcut for setting the "src" attribute
func (t *Tag) Role(role string) *Tag {
	return t.SetAttribute("role", role)
}

// SetAttribute sets the value of an attribute
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

// SrcIf sets the "src" attribute if a condition is met
func (t *Tag) SrcIf(condition bool, src string) *Tag {
	if condition {
		return t.Src(src)
	}

	return t
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

// TargetIf sets the target if a condition is met
func (t *Tag) TargetIf(condition bool, target string) *Tag {
	if condition {
		return t.Target(target)
	}

	return t
}

// Text shortcut for AddText
func (t *Tag) Text(text string) *Tag {
	return t.AddText(text)
}

// TextIf adds escaped text if a condition is met
func (t *Tag) TextIf(condition bool, text string) *Tag {
	if condition {
		return t.AddText(text)
	}

	return t
}

// TextIfElse adds escaped text if a condition is met
func (t *Tag) TextIfElse(condition bool, textIf string, textElse string) *Tag {
	if condition {
		return t.AddText(textIf)
	}

	return t.AddText(textElse)
}

// Title shortcut for setting the "title" attribute
func (t *Tag) Title(title string) *Tag {
	return t.SetAttribute("title", title)
}

// TitleIf sets the title if a condition is met
func (t *Tag) TitleIf(condition bool, title string) *Tag {
	if condition {
		return t.Title(title)
	}

	return t
}

var shortTags = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

func isShortTag(tagName string) bool {
	return shortTags[tagName]
}

// ToHTML returns HTML from Node
func (t *Tag) ToHTML() string {
	if t == nil {
		return ""
	}

	if t.TagName == "" {
		return t.TagContent + t.childrenToString()
	}

	if isShortTag(t.TagName) {
		return "<" + t.TagName + t.attrsToString() + " />"
	}

	return "<" + t.TagName + t.attrsToString() + ">" +
		t.TagContent +
		t.childrenToString() +
		"</" + t.TagName + ">"
}

// Type shortcut for setting the "type" attribute
func (t *Tag) Type(inputType string) *Tag {
	return t.SetAttribute("type", inputType)
}

// TypeIf sets the type if a condition is met
func (t *Tag) TypeIf(condition bool, inputType string) *Tag {
	if condition {
		return t.Type(inputType)
	}

	return t
}

// Value shortcut for setting the "value" attribute
func (t *Tag) Value(value string) *Tag {
	return t.SetAttribute("value", value)
}

// ValueIf sets the value if a condition is met
func (t *Tag) ValueIf(condition bool, value string) *Tag {
	if condition {
		return t.Value(value)
	}

	return t
}

func (t Tag) attrsToString() string {
	if len(t.TagAttributes) == 0 {
		return ""
	}

	keys := make([]string, 0, len(t.TagAttributes))
	for k := range t.TagAttributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer

	for _, key := range keys {
		buf.WriteString(t.attrToString(key, t.TagAttributes[key]))
	}

	return buf.String()
}

// attrToString converts a single key/value attribute to string
// the empty values are skipped
// the value attribute is kept, as it has special role for select element in forms
func (t Tag) attrToString(key string, value string) string {
	if strings.Trim(value, " ") == "" {
		if key != "value" {
			return ``
		}
	}
	return ` ` + key + `="` + escapeHtmlAttr(value) + `"`
}

func (t Tag) childrenToString() string {
	childrenString := ""

	for _, child := range t.TagChildren {
		if child == nil {
			continue
		}

		childrenString += child.ToHTML()
	}

	return childrenString
}
