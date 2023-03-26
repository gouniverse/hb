package hb

import (
	"strings"
	"testing"
)

func TestHx(t *testing.T) {
	input := NewButton().
		Hx("post", "http://test.com").
		Hx("include", "#DivID").
		Hx("target", "#PageID").
		Hx("swap", "outerHTML").
		HTML("Submit").
		ToHTML()

	if strings.Contains(input, ` hx-post="http://test.com"`) == false {
		t.Error(`Does not contain ' hx-post="http://test.com"', Output:` + input)
	}

	if strings.Contains(input, ` hx-include="#DivID"`) == false {
		t.Error(`Does not contain ' hx-include="#DivID"', Output:` + input)
	}

	if strings.Contains(input, ` hx-target="#PageID"`) == false {
		t.Error(`Does not contain ' hx-target="#PageID"', Output:` + input)
	}

	if strings.Contains(input, ` hx-swap="outerHTML"`) == false {
		t.Error(`Does not contain ' hx-swap="outerHTML"', Output:` + input)
	}

	expected := `<button hx-include="#DivID" hx-post="http://test.com" hx-swap="outerHTML" hx-target="#PageID">Submit</button>`
	if input != expected {
		t.Error(`Expected: `, expected, `, Output:`, input)
	}
}
