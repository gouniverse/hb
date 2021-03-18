package html

import (
	"strings"
	"testing"
)

func TestTagCode(t *testing.T) {
	tag := NewCode()
	h := tag.ToHTML()
	if strings.Contains(h, "<code></code>") == false {
		t.Error("Does not contain '<code></code>'", "Output:"+h)
	}
}

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

func TestTagTable(t *testing.T) {
	tag := NewTable()
	h := tag.ToHTML()
	if strings.Contains(h, "<table></table>") == false {
		t.Error("Does not contain '<table></table>'", "Output:"+h)
	}
}

func TestTagSelect(t *testing.T) {
	tag := NewSelect()
	h := tag.ToHTML()
	if strings.Contains(h, "<select></select>") == false {
		t.Error("Does not contain '<select></select>'", "Output:"+h)
	}
}

func TestTagTemplate(t *testing.T) {
	tag := NewTemplate()
	h := tag.ToHTML()
	if strings.Contains(h, "<template></template>") == false {
		t.Error("Does not contain '<template></template>'", "Output:"+h)
	}
}
