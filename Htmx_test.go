package hb

import (
	"strings"
	"testing"
)

func TestHx(t *testing.T) {
	input := Button().
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

func TestHxShorts(t *testing.T) {
	input := Button().
		HxPost("http://test.com").
		HxInclude("#DivID").
		HxIndicator("#ButtonIndicator").
		HxTarget("#PageID").
		HxSwap("outerHTML").
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

	if strings.Contains(input, ` hx-indicator="#ButtonIndicator"`) == false {
		t.Error(`Does not contain ' hx-indicator="#ButtonIndicator"', Output:` + input)
	}

	if strings.Contains(input, ` hx-swap="outerHTML"`) == false {
		t.Error(`Does not contain ' hx-swap="outerHTML"', Output:` + input)
	}

	expected := `<button hx-include="#DivID" hx-indicator="#ButtonIndicator" hx-post="http://test.com" hx-swap="outerHTML" hx-target="#PageID">Submit</button>`
	if input != expected {
		t.Error(`Expected: `, expected, `, Output:`, input)
	}
}
