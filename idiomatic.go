package hb

// This file contains idiomatic functions for creating HTML tags
// These functions all start with the New prefix
// Shortcuts are defined in shortcuts.go

import (
	"strings"
	"testing"
	"text/template"
)

func NewA() *Tag {
	return NewHyperlink()
}

// NewArticle represents an ARTICLE tag
// Shortcut method exists: Article()
func NewArticle() *Tag {
	return &Tag{TagName: "article"}
}

// NewAside represents an ASIDE tag
// Shortcut method exists: Aside()
func NewAside() *Tag {
	return &Tag{TagName: "aside"}
}

// NewBR represents a BR tag
// Shortcut method exists: BR()
func NewBR() *Tag {
	return &Tag{TagName: "br"}
}

// NewButton represents a BUTTON tag
// Shortcut method exists: Button()
func NewButton() *Tag {
	return &Tag{TagName: "button"}
}

// NewCaption is a shortcut to create a new CAPTION tag
// Shortcut method exists: Caption()
func NewCaption() *Tag {
	return &Tag{TagName: "caption"}
}

// NewCode represents a CODE tag
// Shortcut method exists: Code()
func NewCode() *Tag {
	return NewTag("code")
}

// NewDiv represents a DIV tag
// Shortcut method exists: Div()
func NewDiv() *Tag {
	return &Tag{TagName: "div"}
}

// NewFigure represents a FIGURE tag
// Shortcut method exists: Figure()
func NewFigure() *Tag {
	return &Tag{TagName: "figure"}
}

// NewFooter is a shortcut to create a new FOOTER tag
// Shortcut method exists: Footer()
func NewFooter() *Tag {
	return &Tag{TagName: "footer"}
}

// NewForm represents a FORM tag
// Shortcut method exists: Form()
func NewForm() *Tag {
	return &Tag{TagName: "form"}
}

// NewHeader is a shortcut to create a new HEADER tag
// Deprecated: replaced by the new method Header()
func NewHeader() *Tag {
	return &Tag{TagName: "header"}
}

// NewHeading1 represents a H1 tag
// Shortcut method exists: Heading1()
func NewHeading1() *Tag {
	return &Tag{
		TagName: "h1",
	}
}

// NewHeading2 represents a H2 tag
// Shortcut method exists: Heading2()
func NewHeading2() *Tag {
	return &Tag{
		TagName: "h2",
	}
}

// NewHeading3 represents a H3 tag
// Shortcut method exists: Heading3()
func NewHeading3() *Tag {
	return &Tag{
		TagName: "h3",
	}
}

// NewHeading4 represents a H4 tag
// Shortcut method exists: Heading4()
func NewHeading4() *Tag {
	return &Tag{
		TagName: "h4",
	}
}

// NewHeading5 represents a H5 tag
// Shortcut method exists: Heading5()
func NewHeading5() *Tag {
	return &Tag{
		TagName: "h5",
	}
}

// NewHeading6 represents a H6 tag
// Shortcut method exists: Heading6()
func NewHeading6() *Tag {
	return &Tag{
		TagName: "h6",
	}
}

// NewH1 represents a H1 tag
// Shortcut method exists: H1()
func NewH1() *Tag {
	return NewHeading1()
}

// NewH2 represents a H2 tag
// Shortcut method exists: H2()
func NewH2() *Tag {
	return NewHeading2()
}

// NewH3 represents a H3 tag
// Shortcut method exists: H3()
func NewH3() *Tag {
	return NewHeading3()
}

// NewH4 represents a H4 tag
// Shortcut method exists: H4()
func NewH4() *Tag {
	return NewHeading4()
}

// NewH5 represents a H5 tag
// Shortcut method exists: H5()
func NewH5() *Tag {
	return NewHeading5()
}

// NewH6 represents a H6 tag
// Shortcut method exists: H6()
func NewH6() *Tag {
	return NewHeading6()
}

// NewHR creates represents a HR tags
// Shortcut method exists: HR()
func NewHR() *Tag {
	return &Tag{TagName: "hr"}
}

// NewHTML creates pure HTML without surrounding tags
// for safe escaped outut use NewText()
// Shortcut method exists: Raw()
// Deprecated: replaced by the new method NewRaw()
func NewHTML(html string) *Tag {
	return &Tag{
		TagName:    "",
		TagContent: html,
	}
}

// NewHyperlink represents a H1 tag
// Shortcut method exists: Hyperlink()
func NewHyperlink() *Tag {
	return &Tag{
		TagName: "a",
	}
}

// NewI represents an I tag
// Shortcut method exists: I()
func NewI() *Tag {
	return &Tag{
		TagName: "i",
	}
}

// NewImage represents a IMG tag
// Shortcut method exists: Image()
func NewImage() *Tag {
	return &Tag{TagName: "img"}
}

// NewInput represents a IMG tag
// Shortcut method exists: Input()
func NewInput() *Tag {
	return &Tag{TagName: "input"}
}

// NewLabel represents a LABEL tag
// Shortcut method exists: Label()
func NewLabel() *Tag {
	return &Tag{TagName: "label"}
}

// NewLI represents a LI tag
// Shortcut method exists: LI()
func NewLI() *Tag {
	return &Tag{TagName: "li"}
}

// NewLink represents a LINK tag
// Shortcut method exists: Link()
func NewLink() *Tag {
	return &Tag{TagName: "link"}
}

// NewMain represents a MAIN tag
// Shortcut method exists: Main()
func NewMain() *Tag {
	return &Tag{TagName: "main"}
}

// NewMeta represents a META tag
// Shortcut method exists: Meta()
func NewMeta() *Tag {
	return &Tag{TagName: "meta"}
}

// NewNav represents a NAV tag
// Shortcut method exists: Nav()
func NewNav() *Tag {
	return &Tag{TagName: "nav"}
}

// NewNavbar represents a NAVBAR tag
// Deprecated: replaced by the new method Navbar()
func NewNavbar() *Tag {
	return &Tag{TagName: "navbar"}
}

// NewOL represents a OL tag
// Shortcut method exists: OL()
func NewOL() *Tag {
	return &Tag{TagName: "ol"}
}

// NewOption represents an OPTION tag
// Shortcut method exists: Option()
func NewOption() *Tag {
	return &Tag{TagName: "option"}
}

// NewP represents a P tag
// Shortcut method exists: P()
func NewP() *Tag {
	return &Tag{TagName: "p"}
}

// NewParagraph represents a IMG tag
// Shortcut method exists: Paragraph()
func NewParagraph() *Tag {
	return &Tag{TagName: "p"}
}

// NewPod represents a wrapper tag
// It serves as a container for other tags,
// and is used to wrap other tags together.
// Deprecated: replaced by the new method Wrap()
func NewPod() *Tag {
	return &Tag{TagName: ""}
}

// NewPRE represents a PRE tag
// Shortcut method exists: PRE()
func NewPRE() *Tag {
	return &Tag{TagName: "pre"}
}

// NewRaw creates pure escaped HTML without surrounding tags
// for safe escaped outut use NewText()
// Shortcut method exists: Raw()
func NewRaw(html string) *Tag {
	return &Tag{
		TagName:    "",
		TagContent: html,
	}
}

// NewScript represents a SCRIPT tag
// Shortcut method exists: Script()
func NewScript(javascript string) *Tag {
	return &Tag{
		TagName:     "script",
		TagChildren: []TagInterface{NewHTML(javascript)},
	}
}

// NewScriptURL represents a SCRIPT tag with URL
// Shortcut method exists: ScriptURL()
func NewScriptURL(javascriptURL string) *Tag {
	return &Tag{
		TagName: "script",
		TagAttributes: map[string]string{
			"src": javascriptURL,
		},
	}
}

// NewSection represents a SECTION tag
// Shortcut method exists: Section()
func NewSection() *Tag {
	return &Tag{TagName: "section"}
}

// NewSelect represents a SELECT tag
// Shortcut method exists: Select()
func NewSelect() *Tag {
	return &Tag{TagName: "select"}
}

// NewSpan represents a SPAN tag
// Shortcut method exists: Span()
func NewSpan() *Tag {
	return &Tag{TagName: "span"}
}

// NewStrong represents a STRONG tag
// Shortcut method exists: Strong()
func NewStrong() *Tag {
	return &Tag{TagName: "strong"}
}

// NewStyle represents a STYLE tag
// Shortcut method exists: Style()
func NewStyle(css string) *Tag {
	return &Tag{
		TagName:     "style",
		TagChildren: []TagInterface{NewHTML(css)},
	}
}

// NewStyleURL represents a LINK tag with URL
// Shortcut method exists: StyleURL()
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
// Shortcut method exists: Sub()
func NewSub() *Tag {
	return &Tag{TagName: "sub"}
}

// NewSup represents a SUP tag
// Shortcut method exists: Sup()
func NewSup() *Tag {
	return &Tag{TagName: "sup"}
}

// NewSwal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
// Shortcut method exists: Swal()
func NewSwal(options SwalOptions) *Tag {
	return NewScript(swalToJS(options))
}

// NewTable represents a TABLE tag
// Shortcut method exists: Table()
func NewTable() *Tag {
	return &Tag{TagName: "table"}
}

// NewTag creates a tag, with the specified name
// useful for custom tags or ones that are not yet
// added to the hb library
func NewTag(tagName string) *Tag {
	return &Tag{
		TagName: tagName,
	}
}

// NewTD represents a TD tag
// Shortcut method exists: TD()
func NewTD() *Tag {
	return &Tag{TagName: "td"}
}

// NewTbody represents a TBODY tag
// Shortcut method exists: Tbody()
func NewTbody() *Tag {
	return &Tag{TagName: "tbody"}
}

// NewTemplate represents a TEMPLATE tag
// Shortcut method exists: Template()
func NewTemplate() *Tag {
	return &Tag{TagName: "template"}
}

// NewText creates pure escaped text without surrounding tags
// Shortcut method exists: Text()
func NewText(text string) *Tag {
	escapedText := template.HTMLEscapeString(text)

	return &Tag{
		TagName:    "",
		TagContent: escapedText,
	}
}

// NewTextArea represents a FORM tag
// Shortcut method exists: TextArea()
func NewTextArea() *Tag {
	return &Tag{
		TagName: "textarea",
	}
}

// NewTH represents a TH tag
// Shortcut method exists: TH()
func NewTH() *Tag {
	return &Tag{TagName: "th"}
}

// NewThead represents a THEAD tag
// Shortcut method exists: Thead()
func NewThead() *Tag {
	return &Tag{TagName: "thead"}
}

// NewTfoot represents a TFOOT tag
// Shortcut method exists: Tfoot()
func NewTfoot() *Tag {
	return &Tag{TagName: "tfoot"}
}

// NewTitle represents a TITLE tag
// Shortcut method exists: Title()
func NewTitle() *Tag {
	return &Tag{TagName: "title"}
}

// NewTR represents a TR tag
// Shortcut method exists: TR()
func NewTR() *Tag {
	return &Tag{TagName: "tr"}
}

// NewUL represents a UL tag
// Shortcut method exists: UL()
func NewUL() *Tag {
	return &Tag{
		TagName: "ul",
	}
}

// NewVideo represents a VIDEO tag
// Shortcut method exists: Video()
func NewVideo() *Tag {
	return &Tag{TagName: "video"}
}

// NewWebpage returns a webpage instance
// Shortcut method exists: Webpage()
func NewWebpage() *HtmlWebpage {
	return &HtmlWebpage{
		charset: "utf-8",
		head:    &Tag{TagName: "head"},
		body:    &Tag{TagName: "body"},
	}
}

// NewWrap is a convenience tagless container to wrap multiple
// elements together. Any attributes added to the wrap tag will
// be lost. If you need to keep these better use a DIV tag
// Shortcut method exists: Wrap()
func NewWrap() *Tag {
	return &Tag{TagName: ""}
}

func TestTagPRE(t *testing.T) {
	tag := NewPRE()
	h := tag.ToHTML()
	if !strings.Contains(h, "<pre></pre") {
		t.Error("Does not contain '<pre></pre>'", "Output:"+h)
	}
}

func TestTagOption(t *testing.T) {
	tag := NewSelect()
	option1 := NewOption().Attr("value", "key1").HTML("value1")
	option2 := NewOption().Attr("value", "key2").Attr("selected", "selected").HTML("value2")
	h := tag.AddChild(option1).AddChild(option2).ToHTML()
	if !strings.Contains(h, "<select><option value=\"key1\">value1</option><option selected=\"selected\" value=\"key2\">value2</option></select>") {
		t.Error("Does not contain '<select><option value=\"key1\">value1</option><option selected=\"selected\" value=\"key2\">value2</option></select>'", "Output:"+h)
	}
}

func TestTagSelect(t *testing.T) {
	tag := NewSelect()
	h := tag.ToHTML()
	if !strings.Contains(h, "<select></select>") {
		t.Error("Does not contain '<select></select>'", "Output:"+h)
	}
}
