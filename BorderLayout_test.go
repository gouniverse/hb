package hb

import (
	"strings"
	"testing"
)

func TestAttr(t *testing.T) {
	bl := NewBorderLayout().Attr("width", "50%")
	blHtml := bl.ToHTML()
	if strings.Contains(blHtml, "width=\"50%\"") == false {
		t.Error("Does not contain 'width=\"50%\"", "Output:"+blHtml)
	}
}
