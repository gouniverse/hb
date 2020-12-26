package html

import (
	"strings"
	"testing"
)

func TestTagSelect(t *testing.T) {
	tag := NewSelect()
	h := tag.ToHTML()
	if strings.Contains(h, "<select></select>") == false {
		t.Error("Does not contain <select></select>", "Output:"+h)
	}
}

func TestTagOption(t *testing.T) {
	tag := NewSelect()
	option1 := NewOption().Attr("value", "key1").HTML("value1")
	option2 := NewOption().Attr("value", "key2").Attr("selected", "selected").HTML("value2")
	h := tag.AddChild(option1).AddChild(option2).ToHTML()
	if strings.Contains(h, "<select><option value=\"key1\">value1</option><option value=\"key2\" selected=\"selected\">value2</option></select>") == false {
		t.Error("Does not contain '<select><option value=\"key1\">value1</option><option value=\"key2\" selected=\"selected\">value2</option></select>'", "Output:"+h)
	}
}
