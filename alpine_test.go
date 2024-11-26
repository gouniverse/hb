package hb

import (
	"strings"
	"testing"
)

func TestX(t *testing.T) {
	input := NewDiv().
		X("data", "{open=false}").
		Child(
			Button().
				XOn("click", "open = true").
				XBind("class", "! open ? 'hidden' : ''").
				HTML("Submit"),
		).
		Child(
			Span().
				X("show", "open").
				HTML("Content..."),
		).
		ToHTML()

	if strings.Contains(input, ` x-data="{open=false}"`) == false {
		t.Error(`Does not contain ' x-data="{open=false}"', Output:` + input)
	}

	if strings.Contains(input, ` x-show="open"`) == false {
		t.Error(`Does not contain ' x-show="open"', Output:` + input)
	}

	if strings.Contains(input, ` x-on:click="open = true"`) == false {
		t.Error(`Does not contain ' x-on:click="open = true"', Output:` + input)
	}

	if strings.Contains(input, ` x-bind:class="! open ? &#39;hidden&#39; : &#39;&#39;"`) == false {
		t.Error(`Does not contain ' x-bind:class="! open ? &#39;hidden&#39; : &#39;&#39;"', Output:` + input)
	}

	expected := `<div x-data="{open=false}"><button x-bind:class="! open ? &#39;hidden&#39; : &#39;&#39;" x-on:click="open = true">Submit</button><span x-show="open">Content...</span></div>`
	if input != expected {
		t.Error(`Expected: `, expected, `, Output:`, input)
	}
}
