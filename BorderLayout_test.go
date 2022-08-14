package hb

import (
	"log"
	"strings"
	"testing"
)

func TestBorderLayoutAttr(t *testing.T) {
	bl := NewBorderLayout()
	blHtml := bl.ToHTML()
	log.Println("WHAT TO EXPECT")
	log.Println(bl.ToHTML())
	if strings.Contains(blHtml, " style=\"width:100%;height:100%;\"") == false {
		t.Error("Does not contain ' style=\"width:100%;height:100%;\"', Output:" + blHtml)
	}
}

func TestBorderLayoutTopCenterBottomLeftRight(t *testing.T) {
	bl := NewBorderLayout()
	bl.AddTop(NewSpan().HTML("TOP"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddCenter(NewSpan().HTML("CENTER"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddBottom(NewSpan().HTML("BOTTOM"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddLeft(NewSpan().HTML("LEFT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddRight(NewSpan().HTML("RIGHT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	blHtml := bl.ToHTML()
	if strings.Contains(blHtml, " style=\"width:100%;height:100%;\"") == false {
		t.Error("Does not contain ' style=\"width:100%;height:100%;\"', Output:" + blHtml)
	}
	if strings.Contains(blHtml, "<span>TOP</span>") == false {
		t.Error("Does not contain '<span>TOP</span>', Output:" + blHtml)
	}
	if strings.Contains(blHtml, "<span>BOTTOM</span>") == false {
		t.Error("Does not contain '<span>BOTTOM</span>', Output:" + blHtml)
	}
	if strings.Contains(blHtml, "<span>LEFT</span>") == false {
		t.Error("Does not contain '<span>LEFT</span>', Output:" + blHtml)
	}
	if strings.Contains(blHtml, "<span>RIGHT</span>") == false {
		t.Error("Does not contain '<span>RIGHT</span>', Output:" + blHtml)
	}
	if strings.Contains(blHtml, "<span>CENTER</span>") == false {
		t.Error("Does not contain '<span>CENTER</span>', Output:" + blHtml)
	}
	if strings.Contains(blHtml, " class=\"Top\"") == false {
		t.Error("Does not contain ' class=\"Top\"', Output:" + blHtml)
	}
	if strings.Contains(blHtml, " class=\"Bottom\"") == false {
		t.Error("Does not contain ' class=\"Bottom\"', Output:" + blHtml)
	}
	if strings.Contains(blHtml, " class=\"Left\"") == false {
		t.Error("Does not contain ' class=\"Left\"', Output:" + blHtml)
	}
	if strings.Contains(blHtml, " class=\"Right\"") == false {
		t.Error("Does not contain ' class=\"Right\"', Output:" + blHtml)
	}
	if strings.Contains(blHtml, " class=\"Center\"") == false {
		t.Error("Does not contain ' class=\"Center\"', Output:" + blHtml)
	}
}
