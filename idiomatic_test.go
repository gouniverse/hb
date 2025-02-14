package hb

import (
	"strings"
	"testing"
)

// ... [previous test functions remain unchanged] ...

func TestNewAddress(t *testing.T) {
	address := NewAddress()
	html := address.ToHTML()

	if strings.Contains(html, "<address>") == false {
		t.Error("Does not contain '<address>', Output:" + html)
	}

	if strings.Contains(html, "</address>") == false {
		t.Error("Does not contain '</address>', Output:" + html)
	}
}

func TestNewB(t *testing.T) {
	b := NewB()
	html := b.ToHTML()

	if strings.Contains(html, "<b>") == false {
		t.Error("Does not contain '<b>', Output:" + html)
	}

	if strings.Contains(html, "</b>") == false {
		t.Error("Does not contain '</b>', Output:" + html)
	}
}

func TestNewBold(t *testing.T) {
	bold := NewBold()
	html := bold.ToHTML()

	if strings.Contains(html, "<b>") == false {
		t.Error("Does not contain '<b>', Output:" + html)
	}

	if strings.Contains(html, "</b>") == false {
		t.Error("Does not contain '</b>', Output:" + html)
	}
}

func TestNewCanvas(t *testing.T) {
	canvas := NewCanvas()
	html := canvas.ToHTML()

	if strings.Contains(html, "<canvas>") == false {
		t.Error("Does not contain '<canvas>', Output:" + html)
	}

	if strings.Contains(html, "</canvas>") == false {
		t.Error("Does not contain '</canvas>', Output:" + html)
	}
}

func TestNewDialog(t *testing.T) {
	dialog := NewDialog()
	html := dialog.ToHTML()

	if strings.Contains(html, "<dialog>") == false {
		t.Error("Does not contain '<dialog>', Output:" + html)
	}

	if strings.Contains(html, "</dialog>") == false {
		t.Error("Does not contain '</dialog>', Output:" + html)
	}
}

func TestNewFieldSet(t *testing.T) {
	fieldSet := NewFieldSet()
	html := fieldSet.ToHTML()

	if strings.Contains(html, "<fieldset>") == false {
		t.Error("Does not contain '<fieldset>', Output:" + html)
	}

	if strings.Contains(html, "</fieldset>") == false {
		t.Error("Does not contain '</fieldset>', Output:" + html)
	}
}

func TestNewIframe(t *testing.T) {
	iframe := NewIframe()
	html := iframe.ToHTML()

	if strings.Contains(html, "<iframe>") == false {
		t.Error("Does not contain '<iframe>', Output:" + html)
	}

	if strings.Contains(html, "</iframe>") == false {
		t.Error("Does not contain '</iframe>', Output:" + html)
	}
}

func TestNewImg(t *testing.T) {
	img := NewImg()
	html := img.ToHTML()

	if strings.Contains(html, "<img />") == false {
		t.Error("Does not contain '<img />', Output:" + html)
	}
}

func TestNewOL(t *testing.T) {
	ol := NewOL()
	html := ol.ToHTML()

	if strings.Contains(html, "<ol>") == false {
		t.Error("Does not contain '<ol>', Output:" + html)
	}

	if strings.Contains(html, "</ol>") == false {
		t.Error("Does not contain '</ol>', Output:" + html)
	}
}

func TestNewP(t *testing.T) {
	p := NewP()
	html := p.ToHTML()

	if strings.Contains(html, "<p>") == false {
		t.Error("Does not contain '<p>', Output:" + html)
	}

	if strings.Contains(html, "</p>") == false {
		t.Error("Does not contain '</p>', Output:" + html)
	}
}

func TestNewParagraph(t *testing.T) {
	paragraph := NewParagraph()
	html := paragraph.ToHTML()

	if strings.Contains(html, "<p>") == false {
		t.Error("Does not contain '<p>', Output:" + html)
	}

	if strings.Contains(html, "</p>") == false {
		t.Error("Does not contain '</p>', Output:" + html)
	}
}

func TestNewPRE(t *testing.T) {
	pre := NewPRE()
	html := pre.ToHTML()

	if strings.Contains(html, "<pre>") == false {
		t.Error("Does not contain '<pre>', Output:" + html)
	}

	if strings.Contains(html, "</pre>") == false {
		t.Error("Does not contain '</pre>', Output:" + html)
	}
}

func TestNewRaw(t *testing.T) {
	raw := NewRaw("<script>alert('test')</script>")
	html := raw.ToHTML()

	if strings.Contains(html, "<script>alert('test')</script>") == false {
		t.Error("Does not contain expected raw content, Output:" + html)
	}
}

func TestNewScript(t *testing.T) {
	script := NewScript("alert('test');")
	html := script.ToHTML()

	if strings.Contains(html, "<script>") == false {
		t.Error("Does not contain '<script>', Output:" + html)
	}

	if strings.Contains(html, "</script>") == false {
		t.Error("Does not contain '</script>', Output:" + html)
	}

	if strings.Contains(html, "alert('test');") == false {
		t.Error("Does not contain expected script content, Output:" + html)
	}
}

func TestNewScriptURL(t *testing.T) {
	script := NewScriptURL("test.js")
	html := script.ToHTML()

	if strings.Contains(html, `<script src="test.js">`) == false {
		t.Error("Does not contain expected script tag with src, Output:" + html)
	}
}

func TestNewSpan(t *testing.T) {
	span := NewSpan()
	html := span.ToHTML()

	if strings.Contains(html, "<span>") == false {
		t.Error("Does not contain '<span>', Output:" + html)
	}

	if strings.Contains(html, "</span>") == false {
		t.Error("Does not contain '</span>', Output:" + html)
	}
}

func TestNewStrong(t *testing.T) {
	strong := NewStrong()
	html := strong.ToHTML()

	if strings.Contains(html, "<strong>") == false {
		t.Error("Does not contain '<strong>', Output:" + html)
	}

	if strings.Contains(html, "</strong>") == false {
		t.Error("Does not contain '</strong>', Output:" + html)
	}
}

func TestNewStyle(t *testing.T) {
	style := NewStyle("body { color: red; }")
	html := style.ToHTML()

	if strings.Contains(html, "<style>") == false {
		t.Error("Does not contain '<style>', Output:" + html)
	}

	if strings.Contains(html, "</style>") == false {
		t.Error("Does not contain '</style>', Output:" + html)
	}

	if strings.Contains(html, "body { color: red; }") == false {
		t.Error("Does not contain expected style content, Output:" + html)
	}
}

func TestNewStyleURL(t *testing.T) {
	style := NewStyleURL("test.css")
	html := style.ToHTML()

	if !strings.Contains(html, `href="test.css"`) || !strings.Contains(html, `rel="stylesheet"`) {
		t.Error("Does not contain required style link attributes, Output:" + html)
	}
}

func TestNewSub(t *testing.T) {
	sub := NewSub()
	html := sub.ToHTML()

	if strings.Contains(html, "<sub>") == false {
		t.Error("Does not contain '<sub>', Output:" + html)
	}

	if strings.Contains(html, "</sub>") == false {
		t.Error("Does not contain '</sub>', Output:" + html)
	}
}

func TestNewSup(t *testing.T) {
	sup := NewSup()
	html := sup.ToHTML()

	if strings.Contains(html, "<sup>") == false {
		t.Error("Does not contain '<sup>', Output:" + html)
	}

	if strings.Contains(html, "</sup>") == false {
		t.Error("Does not contain '</sup>', Output:" + html)
	}
}

func TestNewSwal(t *testing.T) {
	options := SwalOptions{
		Title: "Test Title",
		Text:  "Test Message",
	}
	swal := NewSwal(options)
	html := swal.ToHTML()

	if strings.Contains(html, "<script>") == false {
		t.Error("Does not contain '<script>', Output:" + html)
	}

	if strings.Contains(html, "</script>") == false {
		t.Error("Does not contain '</script>', Output:" + html)
	}

	if strings.Contains(html, "Test Title") == false {
		t.Error("Does not contain expected title, Output:" + html)
	}

	if strings.Contains(html, "Test Message") == false {
		t.Error("Does not contain expected message, Output:" + html)
	}
}

func TestNewTfoot(t *testing.T) {
	tfoot := NewTfoot()
	html := tfoot.ToHTML()

	if strings.Contains(html, "<tfoot>") == false {
		t.Error("Does not contain '<tfoot>', Output:" + html)
	}

	if strings.Contains(html, "</tfoot>") == false {
		t.Error("Does not contain '</tfoot>', Output:" + html)
	}
}

func TestNewTitle(t *testing.T) {
	title := NewTitle()
	html := title.ToHTML()

	if strings.Contains(html, "<title>") == false {
		t.Error("Does not contain '<title>', Output:" + html)
	}

	if strings.Contains(html, "</title>") == false {
		t.Error("Does not contain '</title>', Output:" + html)
	}
}

func TestNewUL(t *testing.T) {
	ul := NewUL()
	html := ul.ToHTML()

	if strings.Contains(html, "<ul>") == false {
		t.Error("Does not contain '<ul>', Output:" + html)
	}

	if strings.Contains(html, "</ul>") == false {
		t.Error("Does not contain '</ul>', Output:" + html)
	}
}

func TestNewWebpage(t *testing.T) {
	webpage := NewWebpage()
	html := webpage.ToHTML()

	if strings.Contains(html, "<html>") == false {
		t.Error("Does not contain '<html>', Output:" + html)
	}

	if strings.Contains(html, "</html>") == false {
		t.Error("Does not contain '</html>', Output:" + html)
	}
}
