package hb

import (
	"encoding/json"
	"text/template"
)

// A is a shortcut to create a new A tag
func A(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHyperlink().Children(children)
}

func Abbr(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "abbr"}).Children(children)
}

func Address(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "address"}).Children(children)
}

func Article(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "article"}).Children(children)
}

func Aside(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "aside"}).Children(children)
}

// B is a shortcut to create a new B tag
func B(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "b"}).Children(children)
}

// Bold is a shortcut to create a new B tag
func Bold(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "b"}).Children(children)
}

// BR is a shortcut to create a new BR tag
func BR() *Tag {
	return &Tag{TagName: "br"}
}

// Button is a shortcut to create a new BUTTON tag
func Button(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "button"}).Children(children)
}

// Canvas is a shortcut to create a new CANVAS tag
func Canvas(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "canvas"}).Children(children)
}

// Caption is a shortcut to create a new CAPTION tag
func Caption(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "caption"}).Children(children)
}

// Code is a shortcut to create a new CODE tag
func Code(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "code"}).Children(children)
}

func Dialog(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "dialog"}).Children(children)
}

// Div is a shortcut to create a new DIV tag
func Div(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "div"}).Children(children)
}

func FieldSet(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "fieldset"}).Children(children)
}

// Footer is a shortcut to create a new FOOTER tag
func Footer(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "footer"}).Children(children)
}

// Form is a shortcut to create a new FORM tag
func Form(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "form"}).Children(children)
}

// Header is a shortcut to create a new HEADER tag
func Header(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "header"}).Children(children)
}

// H1 is a shortcut to create a new H1 tag
func H1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h1"}).Children(children)
}

// H2 is a shortcut to create a new H2 tag
func H2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h2"}).Children(children)
}

// H3 is a shortcut to create a new H3 tag
func H3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h3"}).Children(children)
}

// H4 is a shortcut to create a new H4 tag
func H4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h4"}).Children(children)
}

// H5 is a shortcut to create a new H5 tag
func H5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h5"}).Children(children)
}

// H6 is a shortcut to create a new H6 tag
func H6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h6"}).Children(children)
}

// Heading1 is a shortcut to create a new H1 tag
func Heading1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h1"}).Children(children)
}

// Heading2 is a shortcut to create a new H2 tag
func Heading2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h2"}).Children(children)
}

// Heading3 is a shortcut to create a new H3 tag
func Heading3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h3"}).Children(children)
}

// Heading4 is a shortcut to create a new H4 tag
func Heading4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h4"}).Children(children)
}

// Heading5 is a shortcut to create a new H5 tag
func Heading5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h5"}).Children(children)
}

// Heading6 is a shortcut to create a new H6 tag
func Heading6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "h6"}).Children(children)
}

// HR is a shortcut to create a new HR tag
func HR() *Tag {
	return &Tag{TagName: "hr"}
}

// Hyperlink is a shortcut to create a new A tag
func Hyperlink(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "a"}).Children(children)
}

// I is a shortcut to create a new I tag
func I(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "i"}).Children(children)
}

// Img is a shortcut to create a new IMG tag
func Iframe(src string) *Tag {
	return (&Tag{TagName: `iframe`}).Src(src)
}

// Img is a shortcut to create a new IMG tag
func Img(src string) *Tag {
	return (&Tag{TagName: "img"}).Src(src)
}

// Image is a shortcut to create a new IMG tag
func Image(src string) *Tag {
	return (&Tag{TagName: "img"}).Src(src)
}

// Input is a shortcut to create a new INPUT tag
func Input() *Tag {
	return (&Tag{TagName: "input"})
}

// Label is a shortcut to create a new LABEL tag
func Label(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "label"}).Children(children)
}

// LI is a shortcut to create a new LI tag
func LI(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "li"}).Children(children)
}

// Link is a shortcut to create a new LINK tag
func Link() *Tag {
	return (&Tag{TagName: "link"})
}

// Main is a shortcut to create a new MAIN tag
func Main(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "main"}).Children(children)
}

// Meta is a shortcut to create a new META tag
func Meta() *Tag {
	return (&Tag{TagName: "meta"})
}

// Nav is a shortcut to create a new NAV tag
func Nav(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "nav"}).Children(children)
}

// Navbar is a shortcut to create a new NAVBAR tag
func Navbar(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "navbar"}).Children(children)
}

// OL is a shortcut to create a new OL tag
func OL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "ol"}).Children(children)
}

// Option is a shortcut to create a new OPTION tag
func Option(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "option"}).Children(children)
}

// P is a shortcut to create a new P tag
func P(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "p"}).Children(children)
}

// Paragraph is a shortcut to create a new P tag
func Paragraph(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "p"}).Children(children)
}

// PRE is a shortcut to create a new PRE tag
func PRE(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "pre"}).Children(children)
}

// Raw is a shortcut to create a new HTML tag
func Raw(html string) *Tag {
	return &Tag{
		TagName:    "",
		TagContent: html,
	}
}

// Script is a shortcut to create a new SCRIPT tag
func Script(javascript string) *Tag {
	return &Tag{
		TagName:     "script",
		TagChildren: []TagInterface{NewHTML(javascript)},
	}
}

// ScriptURL is a shortcut to create a new SCRIPT tag
func ScriptURL(javascriptURL string) *Tag {
	return &Tag{
		TagName: "script",
		TagAttributes: map[string]string{
			"src": javascriptURL,
		},
	}
}

// Section is a shortcut to create a new SECTION tag
func Section(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "section"}).Children(children)
}

// Select is a shortcut to create a new SELECT tag
func Select(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "select"}).Children(children)
}

// Span is a shortcut to create a new SPAN tag
func Span(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "span"}).Children(children)
}

// Strong is a shortcut to create a new STRONG tag
func Strong(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "strong"}).Children(children)
}

// Style is a shortcut to create a new STYLE tag
func Style(css string) *Tag {
	return &Tag{
		TagName:     "style",
		TagChildren: []TagInterface{Raw(css)},
	}
}

// StyleURL is a shortcut to create a new STYLE tag
func StyleURL(styleURL string) *Tag {
	return &Tag{
		TagName: "link",
		TagAttributes: map[string]string{
			"href": styleURL,
			"rel":  "stylesheet",
		},
	}
}

// Sub is a shortcut to create a new SUB tag
func Sub(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "sub"}).Children(children)
}

// Sup is a shortcut to create a new SUP tag
func Sup(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "sup"}).Children(children)
}

// Swal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
func Swal(options SwalOptions) TagInterface {
	optionsBytes, err := json.Marshal(options)

	var optionsJSON string

	if err != nil {
		optionsJSON = ""
	} else {
		optionsJSON = string(optionsBytes)
	}

	swal := `Swal.fire(` + optionsJSON + `)`

	if options.ConfirmCallback != "" {
		swal += `.then((result) => {
			if (result.isConfirmed) {
				` + options.ConfirmCallback + `
			}
		});`
	}

	return Script(swal)
}

// Template is a shortcut to create a new TEMPLATE tag
func Template() *Tag {
	return &Tag{TagName: "template"}
}

// Text creates pure escaped text without surrounding tags
// to use a raw text use Raw
func Text(text string) *Tag {
	escapedText := template.HTMLEscapeString(text)

	return &Tag{
		TagName:    "",
		TagContent: escapedText,
	}
}

// TextArea is a shortcut to create a new TEXTAREA tag
func TextArea() *Tag {
	return (&Tag{TagName: "textarea"})
}

// Table is a shortcut to create a new TABLE tag
func Table(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "table"}).Children(children)
}

// Tbody is a shortcut to create a new TBODY tag
func Tbody(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "tbody"}).Children(children)
}

// TD is a shortcut to create a new TD tag
func TD(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "td"}).Children(children)
}

// TH is a shortcut to create a new TH tag
func TH(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "th"}).Children(children)
}

// Thead is a shortcut to create a new THEAD tag
func Thead(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "thead"}).Children(children)
}

// Tfoot is a shortcut to create a new TFOOT tag
func Tfoot(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "tfoot"}).Children(children)
}

// TR is a shortcut to create a new TR tag
func TR(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "tr"}).Children(children)
}

// Title is a shortcut to create a new title tag
func Title(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "title"}).Children(children)
}

// UL is a shortcut to create a new UL tag
func UL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "ul"}).Children(children)
}

// Webpage is a shortcut to create a new HTML page
func Webpage(children ...TagInterface) *HtmlWebpage {
	children = append([]TagInterface{}, children...)
	return NewWebpage().Children(children)
}

// Wrap is a convenience tagless container to wrap multiple
// elements together. Any attributes added to the wrap tag will
// be lost. If you need to keep these better use a DIV tag
func Wrap(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: ""}).Children(children)
}
