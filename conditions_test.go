package hb

import (
	"strings"
	"testing"
)

func TestIf(t *testing.T) {
	inputTrue := If(true, NewDiv().HTML("true"))

	if inputTrue == nil {
		t.Error(`True input MUST NOT be nil`)
	}

	if strings.Contains(inputTrue.ToHTML(), `<div>true</div>`) == false {
		t.Error(`Does not contain '<div>true</div>', Output:` + inputTrue.ToHTML())
	}

	inputFalse := If(false, NewDiv().HTML("true"))

	if inputFalse != nil {
		t.Error(`False input MUST be nil`)
	}
}

func TestIfElse(t *testing.T) {
	inputTrue := IfElse(true, NewDiv().HTML("true"), NewDiv().HTML("false"))

	if strings.Contains(inputTrue.ToHTML(), `<div>true</div>`) == false {
		t.Error(`Does not contain '<div>true</div>', Output:` + inputTrue.ToHTML())
	}

	inputFalse := IfElse(true, NewDiv().HTML("true"), NewDiv().HTML("false"))

	if strings.Contains(inputFalse.ToHTML(), `<div>true</div>`) == false {
		t.Error(`Does not contain '<div>true</div>', Output:` + inputFalse.ToHTML())
	}
}

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
