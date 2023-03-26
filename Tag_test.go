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
	if strings.Contains(divHtml, "<div><img width=\"100\" /></div>") == false {
		t.Error("Does not contain '<div><img width=\"100\" /></div>'", "Output:"+divHtml)
	}
}

func TestChildren(t *testing.T) {
	img := NewImage().Attr("width", "100")
	div := NewDiv().Children([]*Tag{
		img,
		img,
	})
	divHtml := div.ToHTML()
	if strings.Contains(divHtml, "<div><img width=\"100\" /><img width=\"100\" /></div>") == false {
		t.Error("Does not contain '<div><img width=\"100\" /><img width=\"100\" /></div>'", "Output:"+divHtml)
	}
}

func TestData(t *testing.T) {
	input := NewDiv().Data("id", "TestID").Data("name", "TestName").ToHTML()

	if strings.Contains(input, `data-id="TestID"`) == false {
		t.Error(`Does not contain 'data-id="TestID"', Output:` + input)
	}

	if strings.Contains(input, `data-name="TestName"`) == false {
		t.Error(`Does not contain 'data-name="TestName"', Output:` + input)
	}
}

func TestFormActionMethodEnctype(t *testing.T) {
	form := NewForm().Method("post").Action("http://test.com/form-post").Enctype(ENCTYPE_FORM_MULTIPART).ToHTML()

	if strings.Contains(form, `method="post"`) == false {
		t.Error(`Does not contain 'method="post"', Output:` + form)
	}

	if strings.Contains(form, `action="http://test.com/form-post"`) == false {
		t.Error(`Does not contain 'action="http://test.com/form-post"', Output:` + form)
	}

	if strings.Contains(form, `enctype="multipart/form-data"`) == false {
		t.Error(`Does not contain 'enctype="multipart/form-data"', Output:` + form)
	}

	if form != `<form action="http://test.com/form-post" enctype="multipart/form-data" method="post"></form>` {
		t.Error(`Does not match '<form action="http://test.com/form-post" enctype="multipart/form-data" method="post"></form>', Output:` + form)
	}
}

func TestHasClass(t *testing.T) {
	img := NewImage().Attr("class", "one").AddClass("two").AddClass("three")
	if img.HasClass("two") == false {
		t.Error("Does not contain class \"two\", Output:" + img.ToHTML())
	}
}

func TestIDNameAndValue(t *testing.T) {
	input := NewInput().ID("first").Name("first_name").Value("John").ToHTML()
	if strings.Contains(input, "id=\"first\"") == false {
		t.Error("Does not contain 'id=\"first\", Output:" + input)
	}
	if strings.Contains(input, "name=\"first_name\"") == false {
		t.Error("Does not contain 'name=\"first_name\", Output:" + input)
	}
	if strings.Contains(input, "value=\"John\"") == false {
		t.Error("Does not contain 'value=\"John\", Output:" + input)
	}
	if input != `<input id="first" name="first_name" value="John" />` {
		t.Error(`Does not match '<input id="first" name="first_name" value="John" />', Output:` + input)
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

func TestHrefTarget(t *testing.T) {
	link := NewHyperlink().Href("http://test.com").Target("_blank").HTML("Test").ToHTML()
	if strings.Contains(link, `href="http://test.com"`) == false {
		t.Error(`Does not contain 'href="http://test.com"', Output:` + link)
	}
	if strings.Contains(link, `target="_blank"`) == false {
		t.Error(`Does not contain 'target="_blank"', Output:` + link)
	}
	if link != `<a href="http://test.com" target="_blank">Test</a>` {
		t.Error(`Does not match '<a href="http://test.com" target="_blank">Test</a>', Output:` + link)
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

func TestOnBlur(t *testing.T) {
	input := NewButton().OnBlur("alert('Focus Lost')").ToHTML()
	if strings.Contains(input, `onblur="alert('Focus Lost')"`) == false {
		t.Error(`Does not contain 'onblur="alert('Focus Lost')"', Output:` + input)
	}
}

func TestOnChange(t *testing.T) {
	input := NewButton().OnChange("alert('Changed')").ToHTML()
	if strings.Contains(input, `onchange="alert('Changed')"`) == false {
		t.Error(`Does not contain 'onchange="alert('Changed')"', Output:` + input)
	}
}

func TestOnClick(t *testing.T) {
	input := NewButton().OnClick("alert('Clicked')").ToHTML()
	if strings.Contains(input, "onclick=\"alert('Clicked')\"") == false {
		t.Error("Does not contain 'onclick=\"alert('Clicked')\", Output:" + input)
	}
}

func TestOnFocus(t *testing.T) {
	input := NewButton().OnFocus("alert('Focus Gained')").ToHTML()
	if strings.Contains(input, `onfocus="alert('Focus Gained')"`) == false {
		t.Error(`Does not contain 'onfocus="alert('Focus Gained')"', Output:` + input)
	}
}

func TestOnKeyDown(t *testing.T) {
	input := NewButton().OnKeyDown("alert('Key Down')").ToHTML()
	if strings.Contains(input, `onkeydown="alert('Key Down')"`) == false {
		t.Error(`Does not contain 'onkeydown="alert('Key Down')"', Output:` + input)
	}
}

func TestOnKeyUp(t *testing.T) {
	input := NewButton().OnKeyUp("alert('Key Up')").ToHTML()
	if strings.Contains(input, `onkeyup="alert('Key Up')"`) == false {
		t.Error(`Does not contain 'onkeyup="alert('Key Up')"', Output:` + input)
	}
}

func TestStyle(t *testing.T) {
	input := NewInput().Style("text-align:center;background:green;").ToHTML()
	if strings.Contains(input, "style=\"text-align:center;background:green;\"") == false {
		t.Error("Does not contain 'style=\"text-align:center;background:green;\", Output:" + input)
	}
}
