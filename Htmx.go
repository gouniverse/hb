package hb

/// This is an extension for HTMX attributes
/// https://htmx.org/reference/

// Hx shortcut for setting an HTMX attribute
// Htmx attributes have the format of hx-{NAME}="{VALUE}"
func (t *Tag) Hx(name string, value string) *Tag {
	return t.SetAttribute("hx-"+name, value)
}

func (t *Tag) HxConfirm(value string) *Tag {
	return t.SetAttribute("hx-confirm", value)
}

func (t *Tag) HxGet(value string) *Tag {
	return t.SetAttribute("hx-get", value)
}

func (t *Tag) HxInclude(value string) *Tag {
	return t.SetAttribute("hx-include", value)
}

func (t *Tag) HxOn(name string, value string) *Tag {
	return t.SetAttribute("hx-on:"+name, value)
}

func (t *Tag) HxPost(value string) *Tag {
	return t.SetAttribute("hx-post", value)
}

func (t *Tag) HxSelect(value string) *Tag {
	return t.SetAttribute("hx-select", value)
}

func (t *Tag) HxSelectOob(value string) *Tag {
	return t.SetAttribute("hx-select-oob", value)
}

func (t *Tag) HxSync(value string) *Tag {
	return t.SetAttribute("hx-post", value)
}

func (t *Tag) HxSwap(value string) *Tag {
	return t.SetAttribute("hx-swap", value)
}

func (t *Tag) HxSwapOob(value string) *Tag {
	return t.SetAttribute("hx-swap-oob", value)
}

func (t *Tag) HxTarget(value string) *Tag {
	return t.SetAttribute("hx-target", value)
}

func (t *Tag) HxTrigger(value string) *Tag {
	return t.SetAttribute("hx-trigger", value)
}

func (t *Tag) HxVals(value string) *Tag {
	return t.SetAttribute("hx-vals", value)
}

func (t *Tag) HxVars(value string) *Tag {
	return t.SetAttribute("hx-vars", value)
}

