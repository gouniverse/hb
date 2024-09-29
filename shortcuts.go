package hb

// Shortcuts. These are still experimental.

func A(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHyperlink().Children(children)
}

func B(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "b"}).Children(children)
}

func Bold(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "b"}).Children(children)
}

func BR() *Tag {
	return NewBR()
}

func Button(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewButton().Children(children)
}

func Code(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewCode().Children(children)
}

func Div(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewDiv().Children(children)
}

func Footer(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewFooter().Children(children)
}

func Form(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewForm().Children(children)
}

func Header(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeader().Children(children)
}

func H1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading1().Children(children)
}

func H2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading2().Children(children)
}

func H3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading3().Children(children)
}

func H4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading4().Children(children)
}

func H5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading5().Children(children)
}

func H6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading6().Children(children)
}

func Heading1(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading1().Children(children)
}

func Heading2(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading2().Children(children)
}

func Heading3(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading3().Children(children)
}

func Heading4(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading4().Children(children)
}

func Heading5(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading5().Children(children)
}

func Heading6(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHeading6().Children(children)
}

func HR() *Tag {
	return NewHR()
}

func Hyperlink(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewHyperlink().Children(children)
}

func I(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewI().Children(children)
}

func Img(src string) *Tag {
	return NewImage().Src(src)
}

func Image(src string) *Tag {
	return NewImage().Src(src)
}

func Input() *Tag {
	return NewInput()
}

func Label(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewLabel().Children(children)
}

func LI(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewLI().Children(children)
}

func Link() *Tag {
	return NewLink()
}

func Main(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "main"}).Children(children)
}

func Meta() *Tag {
	return NewMeta()
}

func Nav(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewNav().Children(children)
}

func Navbar(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewNavbar().Children(children)
}

func OL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewOL().Children(children)
}

func Option(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewOption().Children(children)
}

func P(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewParagraph().Children(children)
}

func Paragraph(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewParagraph().Children(children)
}

func PRE(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "pre"}).Children(children)
}

func Raw(html string) *Tag {
	return NewHTML(html)
}

func Script(javascript string) *Tag {
	return NewScript(javascript)
}

func ScriptURL(javascriptURL string) *Tag {
	return NewScriptURL(javascriptURL)
}

func Section(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSection().Children(children)
}

func Select(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSelect().Children(children)
}

func Span(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewSpan().Children(children)
}

func Strong(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "strong"}).Children(children)
}

func Style(style string) *Tag {
	return NewStyle(style)
}

func StyleURL(styleURL string) *Tag {
	return NewStyleURL(styleURL)
}

func Sub(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "sub"}).Children(children)
}

func Sup(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return (&Tag{TagName: "sup"}).Children(children)
}

func Swal(options SwalOptions) *Tag {
	return NewSwal(options)
}

func Text(text string) *Tag {
	return NewText(text)
}

func TextArea() *Tag {
	return NewTextArea()
}

func Table(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTable().Children(children)
}

func Tbody(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTbody().Children(children)
}

func TD(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTD().Children(children)
}

func TH(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTH().Children(children)
}

func Thead(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewThead().Children(children)
}

func TR(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewTR().Children(children)
}

func UL(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewUL().Children(children)
}

func Webpage(children ...TagInterface) *HtmlWebpage {
	children = append([]TagInterface{}, children...)
	return NewWebpage().Children(children)
}

func Wrap(children ...TagInterface) *Tag {
	children = append([]TagInterface{}, children...)
	return NewWrap().Children(children)
}
