# HTML

Unpretentious short and sweet HTML Builder

```
msg := html.NewHTML("This is in development. Please come back later")
div := html.NewDiv().Attr("class", "container").AddChild(msg)
style := html.NewStyle("section{padding:40px;}")
section := html.NewSection().AddChild(style).AddChild(div)
section.ToHTML()
````
  

# Implemented Tag Shortcuts

- NewDiv()
- NewForm()
- NewHTML()
- NewHeading1()
- NewHeading2()
- NewHeading3()
- NewHeading4()
- NewHeading5()
- NewHeading6()
- NewHyperlink()
- NewImg()
- NewParagraph()
- NewScript()
- NewStyle()
- NewSection()
