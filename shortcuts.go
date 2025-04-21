package hb

// This file contains short functions for creating HTML tags
// Idiomatic functions which start with the New prefix are
// defined in idiomatic.go

import (
	"text/template"
)

// A is a shortcut to create a new A tag
func A(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHyperlink().Children(children)
}

func Abbr(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewAbbr().Children(children)
}

func Address(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewAddress().Children(children)
}

func Article(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewArticle().Children(children)
}

func Aside(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewAside().Children(children)
}

// B is a shortcut to create a new B tag
func B(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewB().Children(children)
}

// Bold is a shortcut to create a new B tag
func Bold(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewB().Children(children)
}

// BR is a shortcut to create a new BR tag
func BR() *Tag {
	return NewBR()
}

// Br is a lowercase alias for BR
func Br() *Tag {
	return BR()
}

// Button is a shortcut to create a new BUTTON tag
func Button(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewButton().Children(children)
}

// Canvas is a shortcut to create a new CANVAS tag
func Canvas(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewCanvas().Children(children)
}

// Caption is a shortcut to create a new CAPTION tag
func Caption(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewCaption().Children(children)
}

// Code is a shortcut to create a new CODE tag
func Code(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewCode().Children(children)
}

func Dialog(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewDialog().Children(children)
}

// Div is a shortcut to create a new DIV tag
func Div(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewDiv().Children(children)
}

func FieldSet(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewFieldSet().Children(children)
}

// Footer is a shortcut to create a new FOOTER tag
func Footer(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewFooter().Children(children)
}

// Form is a shortcut to create a new FORM tag
func Form(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewForm().Children(children)
}

// Header is a shortcut to create a new HEADER tag
func Header(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeader().Children(children)
}

// H1 is a shortcut to create a new H1 tag
func H1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH1().Children(children)
}

// H2 is a shortcut to create a new H2 tag
func H2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH2().Children(children)
}

// H3 is a shortcut to create a new H3 tag
func H3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH3().Children(children)
}

// H4 is a shortcut to create a new H4 tag
func H4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH4().Children(children)
}

// H5 is a shortcut to create a new H5 tag
func H5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH5().Children(children)
}

// H6 is a shortcut to create a new H6 tag
func H6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewH6().Children(children)
}

// Heading1 is a shortcut to create a new H1 tag
func Heading1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading1().Children(children)
}

// Heading2 is a shortcut to create a new H2 tag
func Heading2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading2().Children(children)
}

// Heading3 is a shortcut to create a new H3 tag
func Heading3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading3().Children(children)
}

// Heading4 is a shortcut to create a new H4 tag
func Heading4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading4().Children(children)
}

// Heading5 is a shortcut to create a new H5 tag
func Heading5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading5().Children(children)
}

// Heading6 is a shortcut to create a new H6 tag
func Heading6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading6().Children(children)
}

// HR is a shortcut to create a new HR tag
func HR() *Tag {
	return NewHR()
}

// Hr is a lowercase alias for HR
func Hr() *Tag {
	return HR()
}

// Hyperlink is a shortcut to create a new A tag
func Hyperlink(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHyperlink().Children(children)
}

// I is a shortcut to create a new I tag
func I(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewI().Children(children)
}

// Img is a shortcut to create a new IMG tag
func Iframe(src string) *Tag {
	return NewIframe().Src(src)
}

// Img is a shortcut to create a new IMG tag
func Img(src string) *Tag {
	return NewImg().Src(src)
}

// Image is a shortcut to create a new IMG tag
func Image(src string) *Tag {
	return NewImage().Src(src)
}

// Input is a shortcut to create a new INPUT tag
func Input() *Tag {
	return NewInput()
}

// Label is a shortcut to create a new LABEL tag
func Label(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewLabel().Children(children)
}

// LI is a shortcut to create a new LI tag
func LI(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewLI().Children(children)
}

// Li is a lowercase alias for LI
func Li(children ...TagInterface) *Tag {
	return LI(children...)
}

// Link is a shortcut to create a new LINK tag
func Link() *Tag {
	return NewLink()
}

// Main is a shortcut to create a new MAIN tag
func Main(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewMain().Children(children)
}

// Meta is a shortcut to create a new META tag
func Meta() *Tag {
	return NewMeta()
}

// Nav is a shortcut to create a new NAV tag
func Nav(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewNav().Children(children)
}

// Navbar is a shortcut to create a new NAVBAR tag
func Navbar(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewNavbar().Children(children)
}

// OL is a shortcut to create a new OL tag
func OL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewOL().Children(children)
}

// Ol is a lowercase alias for OL
func Ol(children ...TagInterface) *Tag {
	return OL(children...)
}

// Option is a shortcut to create a new OPTION tag
func Option(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewOption().Children(children)
}

// P is a shortcut to create a new P tag
func P(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewP().Children(children)
}

// Paragraph is a shortcut to create a new P tag
func Paragraph(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewParagraph().Children(children)
}

// PRE is a shortcut to create a new PRE tag
func PRE(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewPRE().Children(children)
}

// Raw is a shortcut to create a new HTML tag
func Raw(html string) *Tag {
	return NewRaw(html)
}

// Script is a shortcut to create a new SCRIPT tag
func Script(javascript string) *Tag {
	return NewScript(javascript)
}

// ScriptURL is a shortcut to create a new SCRIPT tag
func ScriptURL(javascriptURL string) *Tag {
	return NewScriptURL(javascriptURL)
}

// Section is a shortcut to create a new SECTION tag
func Section(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSection().Children(children)
}

// Select is a shortcut to create a new SELECT tag
func Select(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSelect().Children(children)
}

// Span is a shortcut to create a new SPAN tag
func Span(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSpan().Children(children)
}

// Strong is a shortcut to create a new STRONG tag
func Strong(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewStrong().Children(children)
}

// Style is a shortcut to create a new STYLE tag
func Style(css string) *Tag {
	return NewStyle(css)
}

// StyleURL is a shortcut to create a new STYLE tag
func StyleURL(styleURL string) *Tag {
	return NewStyleURL(styleURL)
}

// Sub is a shortcut to create a new SUB tag
func Sub(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSub().Children(children)
}

// Sup is a shortcut to create a new SUP tag
func Sup(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSup().Children(children)
}

// Template is a shortcut to create a new TEMPLATE tag
func Template() *Tag {
	return NewTemplate()
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
	return NewTextArea()
}

// Textarea is a lowercase alias for TextArea
func Textarea() *Tag {
	return TextArea()
}

// Table is a shortcut to create a new TABLE tag
func Table(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTable().Children(children)
}

// Tbody is a shortcut to create a new TBODY tag
func Tbody(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTbody().Children(children)
}

// TBody is an alternative capitalization alias for Tbody
func TBody(children ...TagInterface) *Tag {
	return Tbody(children...)
}

// TD is a shortcut to create a new TD tag
func TD(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTD().Children(children)
}

// Td is a lowercase alias for TD
func Td(children ...TagInterface) *Tag {
	return TD(children...)
}

// TH is a shortcut to create a new TH tag
func TH(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTH().Children(children)
}

// Th is a lowercase alias for TH
func Th(children ...TagInterface) *Tag {
	return TH(children...)
}

// Thead is a shortcut to create a new THEAD tag
func Thead(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewThead().Children(children)
}

// THead is an alternative capitalization alias for Thead
func THead(children ...TagInterface) *Tag {
	return Thead(children...)
}

// Tfoot is a shortcut to create a new TFOOT tag
func Tfoot(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTfoot().Children(children)
}

// TFoot is an alternative capitalization alias for Tfoot
func TFoot(children ...TagInterface) *Tag {
	return Tfoot(children...)
}

// TR is a shortcut to create a new TR tag
func TR(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTR().Children(children)
}

// Tr is a lowercase alias for TR
func Tr(children ...TagInterface) *Tag {
	return TR(children...)
}

// Title is a shortcut to create a new title tag
func Title(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTitle().Children(children)
}

// UL is a shortcut to create a new UL tag
func UL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewUL().Children(children)
}

// Ul is a lowercase alias for UL
func Ul(children ...TagInterface) *Tag {
	return UL(children...)
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
	return NewWrap().Children(children)
}
