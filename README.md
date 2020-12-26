# HTML

Unpretentious short and sweet HTML Builder

## Example

- Section containing div container (Bootstrap) with a message "Hello world"

```go
import hb "github.com/gouniverse/html"
	
// 1. Create a div container with "Hello world" message
div := hb.NewDiv().Attr("class", "container").HTML("Hello world")

// 2. Create a section with padding of 40px containing the container
section := hb.NewSection().Attr("style","padding:40px;").AddChild(div)

// 3. Render to HTML to display
html := section.ToHTML()
```

!!! For more examples look below in the README

## Installation

```ssh
go get -u github.com/gouniverse/html@v1.10.0
```

## Implemented Tag Shortcuts

- NewButton()
- NewDiv()
- NewForm()
- NewHTML(html string)
- NewHeading1()
- NewHeading2()
- NewHeading3()
- NewHeading4()
- NewHeading5()
- NewHeading6()
- NewHyperlink()
- NewImage()
- NewInput()
- NewLI()
- NewLabel()
- NewNav()
- NewNavbar()
- NewOL()
- NewParagraph()
- NewScript()
- NewScriptURL()
- NewSpan()
- NewStyle()
- NewStyleURL()
- NewSection()
- NewTextArea()
- NewUL()
- NewWebpage()

## Tag Methods

- Attr (shortcut for SetAttribute)
- HTML (shortcut for AddHTML)
- AddChild(tag Tag)
- AddChildren(tag []Tag)
- AddHTML(html string)
- GetAttribute(key string)
- SetAttribute(key, value string)
- ToHTML()

## Webpage Methods
- AddChild(child *Tag)
- SetFavicon(favicon string)
- SetTitle(title string)
- AddScripts(scripts []string)
- AddScript(script string)
- AddScriptURLs(scriptURLs []string)
- AddScriptURL(scriptURL string)
- AddStyle(style string)
- AddStyles(styles []string)
- AddStyleURL(styleURL string)
- AddStyleURLs(styleURLs []string)

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

- Bootstrap login form

```go
// Elements for the form
header := hb.NewHeading3().HTML("Please sign in").Attr("style", "margin:0px;")
emailLabel := hb.NewLabel().HTML("E-mail Address")
emailInput := hb.NewInput().Attr("class", "form-control").Attr("name", "email").Attr("placeholder", "Enter e-mail address")
emailFormGroup := hb.NewDiv().Attr("class", "form-group").AddChild(emailLabel).AddChild(emailInput)
passwordLabel := hb.NewLabel().AddChild(hb.NewHTML("Password"))
passwordInput := hb.NewInput().Attr("class", "form-control").Attr("name", "password").Attr("type", "password").Attr("placeholder", "Enter password")
passwordFormGroup := hb.NewDiv().Attr("class", "form-group").AddChild(passwordLabel).AddChild(passwordInput)
buttonLogin := hb.NewButton().Attr("class", "btn btn-lg btn-success btn-block").HTML("Login")
buttonRegister := hb.NewHyperlink().Attr("class", "btn btn-lg btn-info float-left").HTML("Register").Attr("href", "auth/register")
buttonForgotPassword := hb.NewHyperlink().Attr("class", "btn btn-lg btn-warning float-right").HTML("Forgot password?").Attr("href", "auth/password-restore")

// Add elements in a card
cardHeader := hb.NewDiv().Attr("class", "card-header").AddChild(header)
cardBody := hb.NewDiv().Attr("class", "card-body").AddChildren([]*hb.Tag{
	emailFormGroup,
	passwordFormGroup,
	buttonLogin,
})
cardFooter := hb.NewDiv().Attr("class", "card-footer").AddChildren([]*hb.Tag{
	buttonRegister,
	buttonForgotPassword,
})
card := hb.NewDiv().Attr("class", "card card-default").Attr("style", "margin:0 auto;max-width: 360px;")
card.AddChild(cardHeader).AddChild(cardBody).AddChild(cardFooter)

// Convert to HTML to display
html := card.ToHTML()
```

- Webpage with title, favicon, font-awesome icons, jQuery and Bootstrap

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

## Changelog
2020.12.26 - Fix for attribute escapes, added tests
