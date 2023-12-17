package hb

import (
	"strings"
	"testing"
)

func TestTernary(t *testing.T) {
	inputTrue := Ternary(true, NewDiv().Text("true"), NewDiv().Text("false"))

	if strings.Contains(inputTrue.ToHTML(), `<div>true</div>`) == false {
		t.Error(`Does not contain '<div>true</div>', Output:` + inputTrue.ToHTML())
	}

	inputFalse := Ternary(true, NewDiv().HTML("true"), NewDiv().HTML("false"))

	if strings.Contains(inputFalse.ToHTML(), `<div>true</div>`) == false {
		t.Error(`Does not contain '<div>true</div>', Output:` + inputFalse.ToHTML())
	}
}
