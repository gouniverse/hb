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

func TestSwapMethod(t *testing.T) {
	tests := []struct {
		method   SwapMethod
		expected string
	}{
		{SwapInnerHTML, "innerHTML"},
		{SwapOuterHTML, "outerHTML"},
		{SwapBeforeEnd, "beforeend"},
		{SwapAfterEnd, "afterend"},
		{SwapBeforeBegin, "beforebegin"},
		{SwapAfterBegin, "afterbegin"},
		{SwapNone, "none"},
	}

	for _, tt := range tests {
		t.Run(string(tt.method), func(t *testing.T) {
			btn := Button().HxSwap(tt.method).ToHTML()
			if !strings.Contains(btn, `hx-swap="`+tt.expected+`"`) {
				t.Errorf("Expected hx-swap=%q, got: %s", tt.expected, btn)
			}
		})
	}
}

func TestHxValidation(t *testing.T) {
	t.Run("EmptyValue", func(t *testing.T) {
		btn := Button().HxPost("").ToHTML()
		if btn != "<button></button>" {
			t.Errorf("Expected empty button, got: %s", btn)
		}
	})

	t.Run("EmptyName", func(t *testing.T) {
		btn := Button().Hx("", "value").ToHTML()
		if btn != "<button></button>" {
			t.Errorf("Expected empty button, got: %s", btn)
		}
	})
}

func BenchmarkHxMethods(b *testing.B) {
	b.Run("HxPost", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Button().HxPost("http://test.com").ToHTML()
		}
	})

	b.Run("HxSwap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Button().HxSwap(SwapOuterHTML).ToHTML()
		}
	})
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
