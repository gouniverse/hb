# HTML Builder

![tests](https://github.com/gouniverse/html/workflows/tests/badge.svg)

Unpretentious short and sweet declarative HTML Builder.

## Benefits

- Valid (X)HTML
- Programatically generate (X)HTML
- Pure GO code
- No need to transfer HTML files
- No need to embed HTML files
- No need for using template files
- Full reuse of the code
- Filly eliminates the quirks of the template package
- Great for email templates
- Easily composable

## Example

- Section containing div container (Bootstrap) with a message "Hello world"

```go
import "github.com/gouniverse/hb"
	
// 1. Create a div container with "Hello world" message
div := hb.NewDiv().Class("container").HTML("Hello world")

// 2. Create a section with padding of 40px containing the container
section := hb.NewSection().Style("padding:40px;").Child(div)

// 3. Render to HTML to display
html := section.ToHTML()
```

!!! For more examples look below in the README

## Installation

```ssh
go get -u github.com/gouniverse/hb@v2.0.0
```

## Implemented Tag Shortcuts

- <b>NewBR()</b> - shortcut for &lt;br> tag
- <b>NewBorderLayout()</b> - border layout with top, bottom, left, right and center slots (see example below how to use it)
- <b>NewButton()</b> - shortcut for &lt;button> tag
- <b>NewCode()</b> - shortcut for &lt;code> tag
- <b>NewDiv()</b> - shortcut for &lt;div> tag
- <b>NewForm()</b> - shortcut for &lt;form> tag
- <b>NewI()</b> - shortcut for &lt;i> tag
- <b>NewHTML(html string)</b> - creates empty tag with the HTML content
- <b>NewHR()</b> - shortcut for &lt;hr> tag
- <b>NewHeading1()</b> - shortcut for &lt;h1> tag
- <b>NewHeading2()</b> - shortcut for &lt;h2> tag
- <b>NewHeading3()</b> - shortcut for &lt;h3> tag
- <b>NewHeading4()</b> - shortcut for &lt;h4> tag
- <b>NewHeading5()</b> - shortcut for &lt;h5> tag
- <b>NewHeading6()</b> - shortcut for &lt;h6> tag
- <b>NewHyperlink()</b> - shortcut for &lt;a> tag
- <b>NewImage()</b> - shortcut for &lt;img> tag
- <b>NewInput()</b> - shortcut for &lt;input> tag
- <b>NewLI()</b> - shortcut for &lt;li> tag
- <b>NewLabel()</b> - shortcut for &lt;label> tag
- <b>NewNav()</b> - shortcut for &lt;nav> tag
- <b>NewNavbar()</b> - shortcut for &lt;navbar> tag
- <b>NewOL()</b> - shortcut for &lt;ol> tag
- <b>NewOption()</b> - shortcut for &lt;option> tag
- <b>NewParagraph()</b> - shortcut for &lt;p> tag
- <b>NewPRE()</b> - shortcut for &lt;pre> tag
- <b>NewScript()</b> - shortcut for &lt;script> tag
- <b>NewScriptURL()</b> - shortcut for &lt;script src="{SRC}"> tag
- <b>NewSelect()</b> - shortcut for &lt;select> tag
- <b>NewSpan()</b> - shortcut for &lt;span> tag
- <b>NewStyle()</b> - shortcut for &lt;style> tag
- <b>NewStyleURL()</b> - shortcut for &lt;link> tag
- <b>NewSection()</b> - shortcut for &lt;section> tag
- <b>NewSub()</b> - shortcut for &lt;sub> tag
- <b>NewSup()</b> - shortcut for &lt;sup> tag
- <b>NewTag(tagName string)</b> - for custom tags
- <b>NewTable()</b> - shortcut for &lt;table> tag
- <b>NewTBody()</b> - shortcut for &lt;tbody> tag
- <b>NewTD()</b> - shortcut for &lt;td> tag
- <b>NewTemplate()</b> - shortcut for &lt;template> tag
- <b>NewTextArea()</b> - shortcut for &lt;textarea> tag
- <b>NewTH()</b> - shortcut for &lt;th> tag
- <b>NewThead()</b> - shortcut for &lt;thead> tag
- <b>NewTR()</b> - shortcut for &lt;tr> tag
- <b>NewUL()</b> - shortcut for &lt;ul> tag
- <b>NewWebpage()</b> - full HTML page withe head, body, meta, styles and scripts

## Tag Methods

- <b>Attr(key, value string)</b> - shortcut for SetAttribute
- <b>Attrs(map[string]string)</b> - shortcut for setting multiple attributes
- <b>HTML(html string)</b> - shortcut for AddHTML
- <b>AddChild(tag Tag)</b> - adds a child element
- <b>AddChildren(tag []Tag)</b> - adds an array of child elements
- <b>AddClass(className string)</b> - adds a class name to the "class" attribute
- <b>AddHTML(html string)</b> - adds HTML content to the element
- <b>Child(tag Tag)</b> - shortcut for AddChild
- <b>Children(tag []Tag)</b> - shortcut for AddChildren
- <b>Class(className string)</b> - shortcut for AddClass
- <b>HasClass(className string)</b> - checks if the class is available
- <b>ID(className string)</b> - shortcut to add an "id" attribute
- <b>GetAttribute(key string) string</b>
- <b>OnClick(js string)</b> - shortcut to add an "onclick" attribute
- <b>SetAttribute(key, value string)</b> - sets an attribute (i.e. id, class, etc)
- <b>Style(style string)</b> - shortcut to add a "style" attribute
- <b>ToHTML() string</b> - outputs HTML code

## Webpage Methods
- <b>AddChild(child *Tag)</b>
- <b>SetFavicon(favicon string)</b>
- <b>SetTitle(title string)</b>
- <b>AddScripts(scripts []string)</b>
- <b>AddScript(script string)</b>
- <b>AddScriptURLs(scriptURLs []string)</b>
- <b>AddScriptURL(scriptURL string)</b>
- <b>AddStyle(style string)</b>
- <b>AddStyles(styles []string)</b>
- <b>AddStyleURL(styleURL string)</b>
- <b>AddStyleURLs(styleURLs []string)</b>

## Border Layout Methods
- <b>AddTop(tag *Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddBottom(tag *Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddLeft(tag *Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddRight(tag *Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddCenter(tag *Tag, alignHorizontal string, alignVertical string)</b>

## Working with Raw Tags

```go
tag := &Tag{
	TagName: "custom-element",
}
tag.toHTML()
```

## Escaping HTML
For safeguarding HTML use the EscapeString method from the standard HTML library

Link with example: https://golang.org/pkg/html/#EscapeString

## Examples

- Border Layout
```go
bl := NewBorderLayout()
	bl.AddTop(NewSpan().HTML("TOP"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddCenter(NewSpan().HTML("CENTER"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddBottom(NewSpan().HTML("BOTTOM"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddLeft(NewSpan().HTML("LEFT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddRight(NewSpan().HTML("RIGHT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	blHtml := bl.ToHTML()
```

- Bootstrap login form

```go
// Elements for the form
header := hb.NewHeading3().HTML("Please sign in").Attr("style", "margin:0px;")
emailLabel := hb.NewLabel().HTML("E-mail Address")
emailInput := hb.NewInput().Class("form-control").Attr("name", "email").Attr("placeholder", "Enter e-mail address")
emailFormGroup := hb.NewDiv().Class("form-group").AddChild(emailLabel).AddChild(emailInput)
passwordLabel := hb.NewLabel().AddChild(hb.NewHTML("Password"))
passwordInput := hb.NewInput().Class("form-control").Attr("name", "password").Attr("type", "password").Attr("placeholder", "Enter password")
passwordFormGroup := hb.NewDiv().Class("form-group").AddChild(passwordLabel).AddChild(passwordInput)
buttonLogin := hb.NewButton().Class("btn btn-lg btn-success btn-block").HTML("Login")
buttonRegister := hb.NewHyperlink().Class("btn btn-lg btn-info float-left").HTML("Register").Attr("href", "auth/register")
buttonForgotPassword := hb.NewHyperlink().Class("btn btn-lg btn-warning float-right").HTML("Forgot password?").Attr("href", "auth/password-restore")

// Add elements in a card
cardHeader := hb.NewDiv().Class("card-header").AddChild(header)
cardBody := hb.NewDiv().Class("card-body").AddChildren([]*hb.Tag{
	emailFormGroup,
	passwordFormGroup,
	buttonLogin,
})
cardFooter := hb.NewDiv().Class("card-footer").AddChildren([]*hb.Tag{
	buttonRegister,
	buttonForgotPassword,
})
card := hb.NewDiv().Class("card card-default").Style("margin:0 auto;max-width: 360px;")
card.AddChild(cardHeader).AddChild(cardBody).AddChild(cardFooter)

// Convert to HTML to display
html := card.ToHTML()
```

Result:

<img src="https://github.com/gouniverse/html/blob/master/screnshots/html-builder-form-login.png" />

- Webpage with title, favicon, Font Awesome icons 4.7.0, jQuery 3.2.1, and Bootstrap 4.5

```go
// 1. Webpage Title
title := "Title"

// 2. Webpage Favicon
favicon := "data:image/x-icon;base64,AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAABNTU0AVKH/AOPj4wDExMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIiAREQEREAIiIBERAREQAiIgIiICIiACIiAiIgIiIAMzMDMzAzMwAzMwMzMDMzACIiAiIgIiIAIiICIiAiIgAzMwMzMDMzADMzAzMwMzMAIiICIiAiIgAiIgIiICIiAAAAAAAAAAAAIiICIiAiIgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

// 3. Webpage
webpage := NewWebpage().SetTitle(title).SetFavicon(favicon).AddStyleURLs([]string{
		"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css",
		"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css",
	}).AddScriptURLs([]string{
		"https://code.jquery.com/jquery-3.2.1.min.js",
		"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.bundle.min.js",
	}).AddStyle(`html,body{height:100%;font-family: Ubuntu, sans-serif;}`).AddChild(NewDiv().HTML("Hello"))
```

- Wrap webpage in a function to be reused as a master template

```go
// Webpage returns the webpage template for the website
func Webpage(title, content string) *hb.Webpage {
	faviconImgCms := `data:image/x-icon;base64,AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAmzKzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEQEAAQERAAEAAQABAAEAAQABAQEBEQABAAEREQEAAAERARARAREAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//wAA//8AAP//AAD//wAA//8AAP//AAD//wAAi6MAALu7AAC6owAAuC8AAIkjAAD//wAA//8AAP//AAD//wAA`
	webpage := hb.NewWebpage()
	webpage.SetTitle(title)
	webpage.SetFavicon(faviconImgCms)

	webpage.AddStyleURLs([]string{
		"https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css",
	})
	webpage.AddScriptURLs([]string{
		"https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js",
		"http://code.jquery.com/jquery-3.5.1.min.js",
		"https://unpkg.com/vue@next",
		"https://cdn.jsdelivr.net/npm/sweetalert2@9",
	})
	webpage.AddScripts([]string{})
	webpage.AddStyle(`html,body{height:100%;font-family: Ubuntu, sans-serif;}`)
	webpage.AddStyle(`body {
		font-family: "Nunito", sans-serif;
		font-size: 0.9rem;
		font-weight: 400;
		line-height: 1.6;
		color: #212529;
		text-align: left;
		background-color: #f8fafc;
	}`)
	webpage.AddChild(hb.NewHTML(content))
	return webpage
}

// Generate HTML
html := webpage("Home", "Hello world").ToHTML()
```

## How to Add a Redirection?

```go
webpage.Head.AddChild(hb.NewMeta().Attr("http-equiv", "refresh").Attr("content", "2; url = https://www.yahoo.com"))
```

## Changelog

2023.01.22 - Added shortcut for &lt;footer> tag

2023.01.14 - Flow pattern applied to BorderLayout methods

2022.09.07 - Added Child and Children shortcuts

2022.08.29 - Added default favicon to Webpage to fix 404 if missing

2022.01.07 - Added Attrs shortcut for setting multiple attributes

2021.07.30 - Added shortcut for &lt;hr> tag

2021.03.20 - Renamed package name to hb to not conflict/confuse with the standard html package

2021.03.18 - Added shortcuts for &lt;template> tag

2021.02.26 - Added shortcuts for &lt;sub>, &lt;sup> tags

2021.01.03 - Added example for webpage layout, and screenshot

2020.12.28 - Added shortcuts for &lt;code>, &lt;pre> tags, shortcuts sorted alphabetically

2020.12.27 - Added shortcuts for &lt;table>, &lt;thead>, &lt;tbody>, &lt;tr>, &lt;th>, &lt;td> tags

2020.12.26 - Fix for attribute escapes, added tests

2020.09.16 - Added shortcuts for &lt;nav>, &lt;navbar>, &lt;ol>, &lt;ul>, &lt;li> tags

2020.06.16 - Initial commit


## Aternatives

- Stevelacy Daz (https://github.com/stevelacy/daz) | Last update: 24 Jun 2021

- Extemplate (https://github.com/dannyvankooten/extemplate) | Last update: 15 Jun 2021

- ThePlant HTMLGO (https://github.com/theplant/htmlgo) | Last update: 18 Sep 2021

- Julvo HTMLGO (https://github.com/julvo/htmlgo) | Last update: 5 May 2020

- Go Components (https://github.com/maragudk/gomponents)

- Uberswe HTML (https://github.com/uberswe/html) | Last update: 24 Feb 2021

- Forms from Structs (https://github.com/joncalhoun/form)

- HTMLX (https://github.com/mdigger/htmlx) | Last update: 20 Jun 2021


