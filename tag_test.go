package hb

import (
	"strings"
	"testing"
)

func TestEscapeAttributes(t *testing.T) {
    tag := &Tag{
		TagName: "div",
	}
	tag.Attr("onclick","page('PAGE_ID')")
	h := tag.ToHTML()
	if strings.Contains(h, "onclick=\"page('PAGE_ID')\"") == false{
		t.Error("Does not contain onclick=\"page('PAGE_ID')\"", "Output:" + h)
	}
}
