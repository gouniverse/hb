/*
Package hb provides HTMX integration for HTML builder
HTMX is a modern library that allows AJAX requests directly from HTML
See https://htmx.org for complete documentation
*/
package hb

// Hx is a generic method for setting any HTMX attribute
// Attributes follow the format hx-{NAME}="{VALUE}"
// Returns the tag for method chaining
func (t *Tag) Hx(name, value string) *Tag {
	if name == "" || value == "" {
		return t
	}
	return t.SetAttribute("hx-"+name, value)
}

// HxConfirm sets the hx-confirm attribute for confirmation dialogs
// Returns the tag for method chaining
func (t *Tag) HxConfirm(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-confirm", value)
}

// HxDelete sets the hx-delete attribute for AJAX DELETE requests
// Returns the tag for method chaining
func (t *Tag) HxDelete(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-delete", value)
}

// HxGet sets the hx-get attribute for AJAX GET requests
// Returns the tag for method chaining
func (t *Tag) HxGet(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-get", value)
}

// HxInclude sets the hx-include attribute for element inclusion
// Returns the tag for method chaining
func (t *Tag) HxInclude(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-include", value)
}

// HxIndicator sets the hx-indicator attribute for loading indicators
// Returns the tag for method chaining
func (t *Tag) HxIndicator(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-indicator", value)
}

// HxOn sets event handlers using the hx-on:{event} attribute
// Returns the tag for method chaining
func (t *Tag) HxOn(name, value string) *Tag {
	if name == "" || value == "" {
		return t
	}
	return t.SetAttribute("hx-on:"+name, value)
}

// HxPatch sets the hx-patch attribute for AJAX PATCH requests
// Returns the tag for method chaining
func (t *Tag) HxPatch(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-patch", value)
}

// HxPost sets the hx-post attribute for AJAX POST requests
// Returns the tag for method chaining
func (t *Tag) HxPost(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-post", value)
}

// HxPut sets the hx-put attribute for AJAX PUT requests
// Returns the tag for method chaining
func (t *Tag) HxPut(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-put", value)
}

// HxSelect sets the hx-select attribute for element selection
// Returns the tag for method chaining
func (t *Tag) HxSelect(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-select", value)
}

// HxSelectOob sets the hx-select-oob attribute for out-of-band selection
// Returns the tag for method chaining
func (t *Tag) HxSelectOob(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-select-oob", value)
}

// HxSync sets the hx-sync attribute for request synchronization
// Returns the tag for method chaining
func (t *Tag) HxSync(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-sync", value)
}

type SwapMethod string

const (
	SwapInnerHTML   SwapMethod = "innerHTML"
	SwapOuterHTML   SwapMethod = "outerHTML"
	SwapBeforeEnd   SwapMethod = "beforeend"
	SwapAfterEnd    SwapMethod = "afterend"
	SwapBeforeBegin SwapMethod = "beforebegin"
	SwapAfterBegin  SwapMethod = "afterbegin"
	SwapNone        SwapMethod = "none"
)

// HxSwap sets the hx-swap attribute for content swapping
// Returns the tag for method chaining
func (t *Tag) HxSwap(method SwapMethod) *Tag {
	return t.SetAttribute("hx-swap", string(method))
}

// HxSwapOob sets the hx-swap-oob attribute for out-of-band swaps
// Returns the tag for method chaining
func (t *Tag) HxSwapOob(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-swap-oob", value)
}

// HxTarget sets the hx-target attribute for request targeting
// Returns the tag for method chaining
func (t *Tag) HxTarget(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-target", value)
}

// HxTrigger sets the hx-trigger attribute for event triggering
// Returns the tag for method chaining
func (t *Tag) HxTrigger(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-trigger", value)
}

// HxVals sets the hx-vals attribute for request values
// Returns the tag for method chaining
func (t *Tag) HxVals(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-vals", value)
}

// HxVars sets the hx-vars attribute for request variables
// Returns the tag for method chaining
func (t *Tag) HxVars(value string) *Tag {
	if value == "" {
		return t
	}
	return t.SetAttribute("hx-vars", value)
}
