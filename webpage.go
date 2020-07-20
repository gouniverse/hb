package html

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
	}
	for _, style := range w.Styles {
		w.Head.AddChild(NewStyle(style))
	}
	for _, scriptURL := range w.ScriptURLs {
		w.Body.AddChild(NewScriptURL(scriptURL))
	}
	for _, script := range w.Scripts {
		w.Body.AddChild(NewScript(script))
	}
	h := NewHTML("<!DOCTYPE html>")
	h.AddChild(w.Head)
	h.AddChild(w.Body)
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
		w.AddScript(script)
	}
	return w
}

// AddScript adds a script to the webpage
func (w *Webpage) AddScript(script string) *Webpage {
	w.Scripts = append(w.Scripts, script)
	return w
}

// AddScriptURLs adds style URLs to the webpage
func (w *Webpage) AddScriptURLs(scriptURLs []string) *Webpage {
	for _, scriptURL := range scriptURLs {
		w.AddScriptURL(scriptURL)
	}
	return w
}

// AddScriptURL adds a style URL to the webpage
func (w *Webpage) AddScriptURL(scriptURL string) *Webpage {
	w.ScriptURLs = append(w.ScriptURLs, scriptURL)
	return w
}

// AddStyles adds styles to the webpage
func (w *Webpage) AddStyles(styles []string) *Webpage {
	for _, style := range styles {
		w.AddStyle(style)
	}
	return w
}

// AddStyle adds a style to the webpage
func (w *Webpage) AddStyle(style string) *Webpage {
	w.Styles = append(w.Styles, style)
	return w
}

// AddStyleURLs adds style URLs to the webpage
func (w *Webpage) AddStyleURLs(styleURLs []string) *Webpage {
	for _, styleURL := range styleURLs {
		w.AddStyleURL(styleURL)
	}
	return w
}

// AddStyleURL adds a style URL to the webpage
func (w *Webpage) AddStyleURL(styleURL string) *Webpage {
	w.StyleURLs = append(w.StyleURLs, styleURL)
	return w
}
