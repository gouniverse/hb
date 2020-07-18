# HTML

Unpretentious short and sweet HTML Builder

## Example
```
msg := html.NewHTML("This is in development. Please come back later")
div := html.NewDiv().Attr("class", "container").AddChild(msg)
style := html.NewStyle("section{padding:40px;}")
section := html.NewSection().AddChild(style).AddChild(div)
section.ToHTML()
````

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
- NewParagraph()
- NewScript()
- NewSpan()
- NewStyle()
- NewSection()

## Tag Methods

- Attr (shortcut for SetAttribute)
- AddChild(tag Tag)
- AddChildren(tag []Tag)
- GetAttribute(key string)
- SetAttribute(key, value string)
- ToHTML

## Working with Raw Tags

```
tag := &Tag{
		TagName: "custom-element",
}
tag.toHTML()
```
