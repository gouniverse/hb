package hb

import (
	"strings"
	"testing"
)

func TestSwalJs(t *testing.T) {
	opts := SwalOptions{
		Title: "Hello",
		Text:  "World",
		Toast: true,
	}
	js := swalToJS(opts)

	expecteds := []string{
		"Swal.fire({",
		"})",
		`"title":"Hello"`,
		`"text":"World"`,
		`"showCancelButton":false`,
		`"showConfirmButton":false`,
		`"showDenyButton":false`,
		`"toast":true`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(js, expected) {
			t.Error("Does not contain: ", expected, ", Output: ", js)
		}
	}
}

func TestSwalJsShowButtons(t *testing.T) {
	opts := SwalOptions{
		Title:             "Hello",
		Text:              "World",
		ShowCancelButton:  true,
		ShowConfirmButton: true,
		ShowDenyButton:    true,
		CancelButtonText:  "Cancel Button",
		DenyButtonText:    "Deny Button",
		ConfirmButtonText: "Confirm Button",
	}
	js := swalToJS(opts)

	expecteds := []string{
		"Swal.fire({",
		"})",
		`"title":"Hello"`,
		`"text":"World"`,
		`"showCancelButton":true`,
		`"showConfirmButton":true`,
		`"showDenyButton":true`,
		`"cancelButtonText":"Cancel Button"`,
		`"denyButtonText":"Deny Button"`,
		`"confirmButtonText":"Confirm Button"`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(js, expected) {
			t.Error("Does not contain: ", expected, ", Output: ", js)
		}
	}
}
