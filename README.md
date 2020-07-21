# HTML

Unpretentious short and sweet HTML Builder

## Example

- Section containing div container (Bootstrap) with a message "Hello world"

```
// 1. Create a div container with "Hello world" message
div := html.NewDiv().Attr("class", "container").HTML("Hello world")

// 2. Create a section with padding of 40px containing the container
section := html.NewSection().Attr("style","padding:40px;").AddChild(div)

// 3. Render to HTML to display
html := section.ToHTML()
````
!!! For more examples look below

## Installation

```
go get -u github.com/gouniverse/html@v1.9.0
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
- NewLabel()
- NewParagraph()
- NewScript()
- NewSpan()
- NewStyle()
- NewSection()
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

## Working with Raw Tags

```
tag := &Tag{
	TagName: "custom-element",
}
tag.toHTML()
```

## Examples

- Bootstrap login form

```
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
