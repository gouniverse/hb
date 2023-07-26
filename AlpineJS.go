package hb

/// This is an extension for Alpine JS attributes
/// https://alpinejs.dev/

// X shortcut for setting an Alpine attribute
// AlpineJS attributes have the format of x-{NAME}="{VALUE}"
func (t *Tag) X(name string, value string) *Tag {
	return t.SetAttribute("x-"+name, value)
}

func (t *Tag) XBind(name string, value string) *Tag {
	return t.SetAttribute("x-bind:"+name, value)
}

func (t *Tag) XOn(name string, value string) *Tag {
	return t.SetAttribute("x-on:"+name, value)
}
