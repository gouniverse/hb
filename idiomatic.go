package hb

// This file contains idiomatic functions for creating HTML tags
// These functions all start with the New prefix
// Shortcuts are defined in shortcuts.go

import (
	"encoding/json"
	"text/template"
)

func NewA() *Tag {
	return NewHyperlink()
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

type SwalOptions struct {
	Background         string `json:"background,omitempty"`
	Backdrop           string `json:"backdrop,omitempty"`
	CancelButtonColor  string `json:"cancelButtonColor,omitempty"`
	CancelButtonText   string `json:"cancelButtonText,omitempty"`
	Color              string `json:"color,omitempty"`
	ConfirmButtonText  string `json:"confirmButtonText,omitempty"`
	ConfirmButtonColor string `json:"confirmButtonColor,omitempty"`
	ConfirmCallback    string `json:"-"`
	CustomClass        string `json:"customClass,omitempty"`
	Grow               string `json:"grow,omitempty"`
	HeightAuto         bool   `json:"heightAuto,omitempty"`
	HTML               string `json:"html,omitempty"`
	Icon               string `json:"icon,omitempty"`
	IconColor          string `json:"iconColor,omitempty"`
	IconHtml           string `json:"iconHtml,omitempty"`
	ImageURL           string `json:"imageUrl,omitempty"`
	ImageWidth         string `json:"imageWidth,omitempty"`
	ImageHeight        string `json:"imageHeight,omitempty"`
	ImageAlt           string `json:"imageAlt,omitempty"`
	Padding            string `json:"padding,omitempty"`
	Position           string `json:"position,omitempty"`
	ShowCancelButton   bool   `json:"showCancelButton,omitempty"`
	ShowConfirmButton  bool   `json:"showConfirmButton,omitempty"`
	Text               string `json:"text,omitempty"`
	TimerProgressBar   bool   `json:"timerProgressBar,omitempty"`
	Title              string `json:"title,omitempty"`
	Timer              int    `json:"timer,omitempty"`
	Toast              string `json:"toast,omitempty"`
	Width              string `json:"width,omitempty"`
}

// NewHTML creates pure HTML without surrounding tags
// Shortcut method exists: Raw()
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

// NewSwal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
// Shortcut method exists: Swal()
func NewSwal(options SwalOptions) *Tag {
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

	return NewScript(swal)
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

// NewOption represents an OPTION tag
// Shortcut method exists: Option()
func NewOption() *Tag {
	return &Tag{TagName: "option"}
}

// NewPod represents a wrapper tag
// It serves as a container for other tags,
// and is used to wrap other tags together.
// Deprecated: replaced by the new method Wrap()
func NewPod() *Tag {
	return &Tag{TagName: ""}
}

// NewScript represents a SCRIPT tag
// Shortcut method exists: Script()
func NewScript(javascript string) *Tag {
	return &Tag{
		TagName:     "script",
		TagChildren: []TagInterface{NewHTML(javascript)},
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

// NewTable represents a TABLE tag
// Shortcut method exists: Table()
func NewTable() *Tag {
	return &Tag{TagName: "table"}
}

// NewTR represents a TR tag
// Shortcut method exists: TR()
func NewTR() *Tag {
	return &Tag{TagName: "tr"}
}

// NewTD represents a TD tag
// Shortcut method exists: TD()
func NewTD() *Tag {
	return &Tag{TagName: "td"}
}

// NewTH represents a TH tag
// Shortcut method exists: TH()
func NewTH() *Tag {
	return &Tag{TagName: "th"}
}

// NewTbody represents a TBODY tag
// Shortcut method exists: Tbody()
func NewTbody() *Tag {
	return &Tag{TagName: "tbody"}
}

// NewThead represents a THEAD tag
// Shortcut method exists: Thead()
func NewThead() *Tag {
	return &Tag{TagName: "thead"}
}

// NewTag creates a tag, with the specified name
// useful for custom tags or ones that are not yet
// added to the hb library
func NewTag(tagName string) *Tag {
	return &Tag{
		TagName: tagName,
	}
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

// NewUL represents a UL tag
// Shortcut method exists: UL()
func NewUL() *Tag {
	return &Tag{
		TagName: "ul",
	}
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
