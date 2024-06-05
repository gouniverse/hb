package hb

import (
	"strings"
)

var _ TagInterface = (*Webpage)(nil)

// Webpage represents a web page
type Webpage struct {
	Tag
	head        *Tag
	body        *Tag
	charset     string
	title       string
	favicon     string
	language    string
	keywords    string
	description string
	attributes  map[string]string
	scriptURLs  []string
	styleURLs   []string
	scripts     []string
	styles      []string
	metas       []TagInterface
}

// ToHTML returns HTML representation of the webpage
func (w *Webpage) ToHTML() string {
	preaddedChildren := w.head.TagChildren

	w.head.TagChildren = []TagInterface{}

	if w.charset == "" {
		w.charset = `utf-8`
	}

	metaCharset := NewMeta().Attr("charset", w.charset)
	w.head.AddChild(metaCharset)
	if w.title != "" {
		titleTag := &Tag{TagName: "title"}
		titleTag.AddChild(NewHTML(w.title))
		w.head.AddChild(titleTag)
	}

	if w.favicon == "" {
		// Set default
		w.favicon = "data:image/x-icon;base64,AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACoqKgNTV1Vg3+De/jutb/4/uHb/QLJ0+FRUVFQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVCw/uHb/i8qn/+zr6/7p6Of+2dzZ/j63df4/uHb/P7h2/1RUVCUAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVC8/uHb/Qrh4//Tz8/7x8fD+7u7t/uvr6v7o5+f+P7h2/z+4dv8/uHb/VFRUJQAAAAAAAAAAAAAAAFJcVwJGu33/Tr+E//r6+v/4+Pj/9vX1/vPz8v/w8O//ZcCP/z+4dv8/uHb/P7h2/z+4dv8AAAAAAAAAAAAAAABKtn3uVMKJ/13Gkv/9/f3//Pv7/vr5+f749/f/W8WQ/lHBh/5Gu33+P7h2/z+4dv8/uHb/VFRUVAAAAAAAAAAAVcKK/2HIlf9ox5r/n93B/ovTs/910qn+cdCk/mjMnP5dxpL+UcCH/kS6ev4/uHb/PrV0/m6tivYAAAAAAAAAAF7Hk/9szaD/ydXP/3/Wsf95x6f+g9e0/n7WsP510qj+aMyc/lvFkP5MvoL+8/Ly/+/v7v7s6+r+AAAAAAAAAABmypr/zNrT/4DWsv+K2br/kNy//4/cvv6H2bj+ftWw/nHQpP5hyJb+UsGH/vj39/719PT+8fHw/gAAAAAAAAAAac2e//X09P52vqH+fr+m/5vgyf+Z4Mf+j9y+/oPXtP510qn+Zcqa/lXDiv7r9e/++fn5/vb29f4AAAAAAAAAAGrNnv75+Pj/9/b2//T08/6d38n+m+DJ/5Dcv/+D17X+dtOp/mbLmv9Vw4v+Rbp7/v39/f9NT05zAAAAAAAAAABcgG8E/Pz8/vr6+v/4+Pj+9vb1/uLm5P6Ayqz+f9ax/3LRpv9jyZf/U8GI/+no5/4+t3X/KSsqAQAAAAAAAAAAAAAAAG7MoPH9/f3/+/v7/4G7pP749/f+esmo/3jTq/9rzZ//XcaS/0ihc//w7+/+U1NTLwAAAAAAAAAAAAAAAAAAAAAAAAAAbsyg8f7+/v/g5+P+gLGc/23DnP739/b+brOR/1TCif/F1sz+UlJSNQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB4jIIEaMKX/v7+/v78/Pz/+/r6/1XCiv+93szuaWpqAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAPw/AADwDwAA4AcAAMADAACAAwAAgAEAAIABAACAAQAAgAEAAIADAADAAwAAwAcAAOAPAAD4HwAA//8AAA=="
	}

	if w.favicon != "" {
		faviconTag := NewLink()
		faviconTag.SetAttribute("href", w.favicon)
		faviconTag.SetAttribute("rel", "icon")
		faviconTag.SetAttribute("type", "image/x-icon")
		w.head.AddChild(faviconTag)
	}

	if w.keywords != "" {
		metaKeywords := NewMeta()
		metaKeywords.SetAttribute("name", "keywords")
		metaKeywords.SetAttribute("content", w.keywords)
		w.head.AddChild(metaKeywords)
	}

	if w.description != "" {
		metaDescription := NewMeta()
		metaDescription.SetAttribute("name", "keywords")
		metaDescription.SetAttribute("content", w.description)
		w.head.AddChild(metaDescription)
	}

	for _, meta := range w.metas {
		w.head.AddChild(meta)
	}

	for _, styleURL := range w.styleURLs {
		if strings.HasPrefix(styleURL, "http") || strings.HasPrefix(styleURL, "//") {
			w.head.AddChild(NewStyleURL(styleURL))
		} else {
			w.head.AddChild(NewHTML(styleURL))
		}
	}
	for _, style := range w.styles {
		if !strings.HasPrefix(style, "<style") {
			w.head.AddChild(NewStyle(style))
		} else {
			w.head.AddChild(NewHTML(style))
		}
	}

	w.head.AddChildren(preaddedChildren)

	for _, scriptURL := range w.scriptURLs {
		if strings.HasPrefix(scriptURL, "http") || strings.HasPrefix(scriptURL, "//") {
			w.body.AddChild(NewScriptURL(scriptURL))
		} else {
			w.body.AddChild(NewHTML(scriptURL))
		}
	}
	for _, script := range w.scripts {
		if !strings.HasPrefix(script, "<script") {
			w.body.AddChild(NewScript(script))
		} else {
			w.body.AddChild(NewHTML(script))
		}
	}

	htmlTag := NewTag("html")
	if w.language != "" {
		w.Attr("lang", w.language)
	}

	for k, v := range w.attributes {
		htmlTag.Attr(k, v)
	}

	return NewHTML("<!DOCTYPE html>").
		Child(
			htmlTag.
				Child(w.head).
				Child(w.body),
		).
		ToHTML()
}

// AddChild adds a tag to the webpage
func (w *Webpage) AddChild(child TagInterface) *Webpage {
	w.body.AddChild(child)
	return w
}

// AddChildren adds tags to the webpage
func (w *Webpage) AddChildren(children []TagInterface) *Webpage {
	for _, child := range children {
		w.AddChild(child)
	}
	return w
}

// AddChild shortcut for AddChild
func (w *Webpage) Child(child TagInterface) *Webpage {
	return w.AddChild(child)
}

// Children shortcut for AddChildren
func (w *Webpage) Children(children []TagInterface) *Webpage {
	return w.AddChildren(children)
}

// SetCharset sets the charset of the webpage
func (w *Webpage) SetCharset(charset string) *Webpage {
	w.charset = charset
	return w
}

// SetFavicon sets the favicon of the webpage
func (w *Webpage) SetFavicon(favicon string) *Webpage {
	w.favicon = favicon
	return w
}

// SetLanguage sets the language of the webpage
func (w *Webpage) SetLanguage(language string) *Webpage {
	w.language = language
	return w
}

// SetTitle sets the title of the webpage
func (w *Webpage) SetTitle(title string) *Webpage {
	w.title = title
	return w
}

// Attr shortcut for SetAttribute
func (w *Webpage) Attr(key, value string) *Webpage {
	return w.SetAttribute(key, value)
}

// Attrs shortcut for setting multiple attributes
func (w *Webpage) Attrs(attrs map[string]string) *Webpage {
	for key, value := range attrs {
		w.SetAttribute(key, value)
	}
	return w
}

// SetAttribute adds a style to the webpage
func (w *Webpage) SetAttribute(key string, value string) *Webpage {
	if value == "" {
		return w
	}

	if w.attributes == nil {
		w.attributes = map[string]string{}
	}

	w.attributes[key] = value
	return w
}

// AddHTML adds an HTML to the body of the webpage
func (w *Webpage) AddHTML(html string) *Webpage {
	w.body.HTML(html)
	return w
}

func (w *Webpage) AddMeta(meta TagInterface) *Webpage {
	w.metas = append(w.metas, meta)
	return w
}

// AddScripts adds scripts to the webpage
func (w *Webpage) AddScripts(scripts []string) *Webpage {
	for _, script := range scripts {
		if script == "" {
			continue
		}
		w.AddScript(script)
	}
	return w
}

// AddScript adds a script to the webpage
func (w *Webpage) AddScript(script string) *Webpage {
	if script == "" {
		return w
	}
	w.scripts = append(w.scripts, script)
	return w
}

// AddScriptURLs adds style URLs to the webpage
func (w *Webpage) AddScriptURLs(scriptURLs []string) *Webpage {
	for _, scriptURL := range scriptURLs {
		if scriptURL == "" {
			continue
		}
		w.AddScriptURL(scriptURL)
	}
	return w
}

// AddScriptURL adds a style URL to the webpage
func (w *Webpage) AddScriptURL(scriptURL string) *Webpage {
	if scriptURL == "" {
		return w
	}
	w.scriptURLs = append(w.scriptURLs, scriptURL)
	return w
}

// AddStyles adds styles to the webpage
func (w *Webpage) AddStyles(styles []string) *Webpage {
	for _, style := range styles {
		if style == "" {
			continue
		}
		w.AddStyle(style)
	}
	return w
}

// AddStyle adds a style to the webpage
func (w *Webpage) AddStyle(style string) *Webpage {
	if style == "" {
		return w
	}
	w.styles = append(w.styles, style)
	return w
}

// AddStyleURL adds a style URL to the webpage
func (w *Webpage) AddStyleURL(styleURL string) *Webpage {
	if styleURL == "" {
		return w
	}
	w.styleURLs = append(w.styleURLs, styleURL)
	return w
}

// AddStyleURLs adds style URLs to the webpage
func (w *Webpage) AddStyleURLs(styleURLs []string) *Webpage {
	for _, styleURL := range styleURLs {
		if styleURL == "" {
			continue
		}
		w.AddStyleURL(styleURL)
	}
	return w
}

// Body returns a pointer to the body tag
func (w *Webpage) Body() *Tag {
	return w.body
}

// Head returns a pointer to the head tag
func (w *Webpage) Head() *Tag {
	return w.head
}

// HTML shortcut for adding HTML to the body
func (w *Webpage) HTML(html string) *Webpage {
	return w.AddHTML(html)
}

// Meta shortcut for adding a meta
func (w *Webpage) Meta(meta *Tag) *Webpage {
	return w.AddMeta(meta)
}

// Script shortcut for adding a script
func (w *Webpage) Script(script string) *Webpage {
	return w.AddScript(script)
}

// ScriptURL shortcut for adding a script URL
func (w *Webpage) ScriptURL(scriptURL string) *Webpage {
	return w.AddScriptURL(scriptURL)
}

// ScriptURLs shortcut for adding script URLs
func (w *Webpage) ScriptURLs(scriptURLs []string) *Webpage {
	return w.AddScriptURLs(scriptURLs)
}

// Style shortcut for adding a style
func (w *Webpage) Style(style string) *Webpage {
	return w.AddStyle(style)
}

// StyleURL shortcut for adding a style URL to the webpage
func (w *Webpage) StyleURL(styleURL string) *Webpage {
	return w.AddStyleURL(styleURL)
}

// StyleURLs shortcut for adding style URLs to the webpage
func (w *Webpage) StyleURLs(styleURLs []string) *Webpage {
	return w.AddStyleURLs(styleURLs)
}
