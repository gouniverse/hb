package hb

/// This is an extension for HTMX attributes
/// https://htmx.org/reference/

// Hx shortcut for setting an HTMX attribute
// Htmx attributes have the format of hx-{NAME}="{VALUE}"
func (t *Tag) Hx(name string, value string) *Tag {
	return t.SetAttribute("hx-"+name, value)
}
