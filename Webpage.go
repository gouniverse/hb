package hb

import "strings"

// Webpage represents a web page
type Webpage struct {
	Tag
	Head        *Tag
	Body        *Tag
	Charset     string
	Title       string
	Favicon     string
	Keywords    string
	Description string
	ScriptURLs  []string
	StyleURLs   []string
	Scripts     []string
	Styles      []string
}

// ToHTML returns HTML representation of the webpage
func (w *Webpage) ToHTML() string {
	metaCharset := NewMeta().Attr("charset", "utf-8")
	w.Head.AddChild(metaCharset)
	if w.Title != "" {
		titleTag := &Tag{TagName: "title"}
		titleTag.AddChild(NewHTML(w.Title))
		w.Head.AddChild(titleTag)
	}
	
	if w.Favicon == "" {
		// Set default
		w.Favicon = "data:image/x-icon;base64,AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACoqKgNTV1Vg3+De/jutb/4/uHb/QLJ0+FRUVFQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVCw/uHb/i8qn/+zr6/7p6Of+2dzZ/j63df4/uHb/P7h2/1RUVCUAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVC8/uHb/Qrh4//Tz8/7x8fD+7u7t/uvr6v7o5+f+P7h2/z+4dv8/uHb/VFRUJQAAAAAAAAAAAAAAAFJcVwJGu33/Tr+E//r6+v/4+Pj/9vX1/vPz8v/w8O//ZcCP/z+4dv8/uHb/P7h2/z+4dv8AAAAAAAAAAAAAAABKtn3uVMKJ/13Gkv/9/f3//Pv7/vr5+f749/f/W8WQ/lHBh/5Gu33+P7h2/z+4dv8/uHb/VFRUVAAAAAAAAAAAVcKK/2HIlf9ox5r/n93B/ovTs/910qn+cdCk/mjMnP5dxpL+UcCH/kS6ev4/uHb/PrV0/m6tivYAAAAAAAAAAF7Hk/9szaD/ydXP/3/Wsf95x6f+g9e0/n7WsP510qj+aMyc/lvFkP5MvoL+8/Ly/+/v7v7s6+r+AAAAAAAAAABmypr/zNrT/4DWsv+K2br/kNy//4/cvv6H2bj+ftWw/nHQpP5hyJb+UsGH/vj39/719PT+8fHw/gAAAAAAAAAAac2e//X09P52vqH+fr+m/5vgyf+Z4Mf+j9y+/oPXtP510qn+Zcqa/lXDiv7r9e/++fn5/vb29f4AAAAAAAAAAGrNnv75+Pj/9/b2//T08/6d38n+m+DJ/5Dcv/+D17X+dtOp/mbLmv9Vw4v+Rbp7/v39/f9NT05zAAAAAAAAAABcgG8E/Pz8/vr6+v/4+Pj+9vb1/uLm5P6Ayqz+f9ax/3LRpv9jyZf/U8GI/+no5/4+t3X/KSsqAQAAAAAAAAAAAAAAAG7MoPH9/f3/+/v7/4G7pP749/f+esmo/3jTq/9rzZ//XcaS/0ihc//w7+/+U1NTLwAAAAAAAAAAAAAAAAAAAAAAAAAAbsyg8f7+/v/g5+P+gLGc/23DnP739/b+brOR/1TCif/F1sz+UlJSNQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB4jIIEaMKX/v7+/v78/Pz/+/r6/1XCiv+93szuaWpqAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAPw/AADwDwAA4AcAAMADAACAAwAAgAEAAIABAACAAQAAgAEAAIADAADAAwAAwAcAAOAPAAD4HwAA//8AAA=="
	}
	
	if w.Favicon != "" {
		faviconTag := NewLink()
		faviconTag.SetAttribute("href", w.Favicon)
		faviconTag.SetAttribute("rel", "icon")
		faviconTag.SetAttribute("type", "image/x-icon")
		w.Head.AddChild(faviconTag)
	}
	if w.Keywords != "" {
		metaKeywords := NewMeta()
		metaKeywords.SetAttribute("name", "keywords")
		metaKeywords.SetAttribute("content", w.Keywords)
		w.Head.AddChild(metaKeywords)
	}
	if w.Description != "" {
		metaDescription := NewMeta()
		metaDescription.SetAttribute("name", "keywords")
		metaDescription.SetAttribute("content", w.Description)
		w.Head.AddChild(metaDescription)
	}
	for _, styleURL := range w.StyleURLs {
		w.Head.AddChild(NewStyleURL(styleURL))
		if strings.HasPrefix(styleURL, "http") || strings.HasPrefix(styleURL, "//") {
			w.Head.AddChild(NewStyleURL(styleURL))
		} else {
			w.Head.AddChild(NewHTML(styleURL))
		}
	}
	for _, style := range w.Styles {
		if !strings.HasPrefix(style, "<style") {
			w.Head.AddChild(NewStyle(style))
		} else {
			w.Head.AddChild(NewHTML(style))
		}
	}
	for _, scriptURL := range w.ScriptURLs {
		if strings.HasPrefix(scriptURL, "http") || strings.HasPrefix(scriptURL, "//") {
			w.Body.AddChild(NewScriptURL(scriptURL))
		} else {
			w.Body.AddChild(NewHTML(scriptURL))
		}
	}
	for _, script := range w.Scripts {
		if !strings.HasPrefix(script, "<script") {
			w.Body.AddChild(NewScript(script))
		} else {
			w.Body.AddChild(NewHTML(script))
		}
	}
	h := NewHTML("<!DOCTYPE html>")
	h.AddChild(NewTag("html").AddChild(w.Head).AddChild(w.Body))
	return h.ToHTML()
}

// AddChild adds a tag to the webpage
func (w *Webpage) AddChild(child *Tag) *Webpage {
	w.Body.AddChild(child)
	return w
}

// SetFavicon sets the favicon of the webpage
func (w *Webpage) SetFavicon(favicon string) *Webpage {
	w.Favicon = favicon
	return w
}

// SetTitle sets the title of the webpage
func (w *Webpage) SetTitle(title string) *Webpage {
	w.Title = title
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
	w.Scripts = append(w.Scripts, script)
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
	w.ScriptURLs = append(w.ScriptURLs, scriptURL)
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
	w.Styles = append(w.Styles, style)
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

// AddStyleURL adds a style URL to the webpage
func (w *Webpage) AddStyleURL(styleURL string) *Webpage {
	if styleURL == "" {
		return w
	}
	w.StyleURLs = append(w.StyleURLs, styleURL)
	return w
}
