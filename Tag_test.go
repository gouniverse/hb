package hb

import (
	"strings"
	"testing"
)

func TestAttr(t *testing.T) {
	img := NewImage().Attr("width", "100")
	imgHtml := img.ToHTML()
	if strings.Contains(imgHtml, "width=\"100\"") == false {
		t.Error("Does not contain 'width=\"100\"", "Output:"+imgHtml)
	}
}

func TestAddClass(t *testing.T) {
	img := NewImage().Attr("class", "one")
	imgHtml := img.AddClass("two").ToHTML()
	if strings.Contains(imgHtml, "class=\"one two\"") == false {
		t.Error("Does not contain 'class=\"one two\", Output:" + imgHtml)
	}
}

func TestChild(t *testing.T) {
	img := NewImage().Attr("width", "100")
	div := NewDiv().Child(img)
	divHtml := div.ToHTML()
	if strings.Contains(divHtml, "<div><img width=\"100\"></div>") == false {
		t.Error("Does not contain '<div><img width=\"100\"></div>'", "Output:"+divHtml)
	}
}

func TestChildren(t *testing.T) {
	img := NewImage().Attr("width", "100")
	div := NewDiv().Children([]*Tag{
		img,
		img,
	})
	divHtml := div.ToHTML()
	if strings.Contains(divHtml, "<div><img width=\"100\"><img width=\"100\"></div>") == false {
		t.Error("Does not contain '<div><img width=\"100\"><img width=\"100\"></div>'", "Output:"+divHtml)
	}
}

func TestHasClass(t *testing.T) {
	img := NewImage().Attr("class", "one").AddClass("two").AddClass("three")
	if img.HasClass("two") == false {
		t.Error("Does not contain class \"two\", Output:" + img.ToHTML())
	}
}

func TestID(t *testing.T) {
	input := NewInput().ID("first").ToHTML()
	if strings.Contains(input, "id=\"first\"") == false {
		t.Error("Does not contain 'id=\"first\", Output:" + input)
	}
}

func TestEscapeAttributes(t *testing.T) {
	tag := &Tag{
		TagName: "div",
	}
	tag.Attr("onclick", "page('PAGE_ID')")
	h := tag.ToHTML()
	if strings.Contains(h, "onclick=\"page('PAGE_ID')\"") == false {
		t.Error("Does not contain onclick=\"page('PAGE_ID')\"", "Output:"+h)
	}
}

func TestAttrs(t *testing.T) {
	img := NewImage().Attrs(map[string]string{
		"width":  "100",
		"height": "40",
	})
	imgHtml := img.ToHTML()
	if strings.Contains(imgHtml, "width=\"100\"") == false {
		t.Error("Does not contain 'width=\"100\"", "Output:"+imgHtml)
	}
	if strings.Contains(imgHtml, "height=\"40\"") == false {
		t.Error("Does not contain 'height=\"40\"", "Output:"+imgHtml)
	}
}

func TestOnClick(t *testing.T) {
	input := NewButton().OnClick("alert('Clicked')").ToHTML()
	if strings.Contains(input, "onclick=\"alert('Clicked')\"") == false {
		t.Error("Does not contain 'onclick=\"alert('Clicked')\", Output:" + input)
	}
}

func TestStyle(t *testing.T) {
	input := NewInput().Style("text-align:center;background:green;").ToHTML()
	if strings.Contains(input, "style=\"text-align:center;background:green;\"") == false {
		t.Error("Does not contain 'style=\"text-align:center;background:green;\", Output:" + input)
	}
}
