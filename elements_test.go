package hb

import (
	"strings"
	"testing"
)

func TestTagPRE(t *testing.T) {
	tag := NewPRE()
	h := tag.ToHTML()
	if strings.Contains(h, "<pre></pre>") == false {
		t.Error("Does not contain '<pre></pre>'", "Output:"+h)
	}
}

func TestTagOption(t *testing.T) {
	tag := NewSelect()
	option1 := NewOption().Attr("value", "key1").HTML("value1")
	option2 := NewOption().Attr("value", "key2").Attr("selected", "selected").HTML("value2")
	h := tag.AddChild(option1).AddChild(option2).ToHTML()
	if strings.Contains(h, "<select><option value=\"key1\">value1</option><option selected=\"selected\" value=\"key2\">value2</option></select>") == false {
		t.Error("Does not contain '<select><option value=\"key1\">value1</option><option selected=\"selected\" value=\"key2\">value2</option></select>'", "Output:"+h)
	}
}

func TestTagSelect(t *testing.T) {
	tag := NewSelect()
	h := tag.ToHTML()
	if strings.Contains(h, "<select></select>") == false {
		t.Error("Does not contain '<select></select>'", "Output:"+h)
	}
}
