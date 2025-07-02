package hb

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	a := A()
	html := a.ToHTML()

	if strings.Contains(html, "<a>") == false {
		t.Error("Does not contain '<a>', Output:" + html)
	}
	if strings.Contains(html, "</a>") == false {
		t.Error("Does not contain '</a>', Output:" + html)
	}
}

func TestAbbr(t *testing.T) {
	a := Abbr()
	html := a.ToHTML()

	if strings.Contains(html, "<abbr>") == false {
		t.Error("Does not contain '<abbr>', Output:" + html)
	}
	if strings.Contains(html, "</abbr>") == false {
		t.Error("Does not contain '</abbr>', Output:" + html)
	}
}

func TestAddress(t *testing.T) {
	a := Address()
	html := a.ToHTML()

	if strings.Contains(html, "<address>") == false {
		t.Error("Does not contain '<address>', Output:" + html)
	}
	if strings.Contains(html, "</address>") == false {
		t.Error("Does not contain '</address>', Output:" + html)
	}
}

func TestArticle(t *testing.T) {
	a := Article()
	html := a.ToHTML()

	if strings.Contains(html, "<article>") == false {
		t.Error("Does not contain '<article>', Output:" + html)
	}
	if strings.Contains(html, "</article>") == false {
		t.Error("Does not contain '</article>', Output:" + html)
	}
}

func TestAside(t *testing.T) {
	a := Aside()
	html := a.ToHTML()

	if strings.Contains(html, "<aside>") == false {
		t.Error("Does not contain '<aside>', Output:" + html)
	}
	if strings.Contains(html, "</aside>") == false {
		t.Error("Does not contain '</aside>', Output:" + html)
	}
}

func TestB(t *testing.T) {
	b := B()
	html := b.ToHTML()

	if strings.Contains(html, "<b>") == false {
		t.Error("Does not contain '<b>', Output:" + html)
	}
	if strings.Contains(html, "</b>") == false {
		t.Error("Does not contain '</b>', Output:" + html)
	}
}

func TestBR(t *testing.T) {
	br := BR()
	html := br.ToHTML()

	if strings.Contains(html, "<br />") == false {
		t.Error("Does not contain '<br />', Output:" + html)
	}
}

func TestButton(t *testing.T) {
	button := Button()
	html := button.ToHTML()

	if strings.Contains(html, "<button>") == false {
		t.Error("Does not contain '<button>', Output:" + html)
	}

	if strings.Contains(html, "</button>") == false {
		t.Error("Does not contain '</button>', Output:" + html)
	}
}

func TestCanvas(t *testing.T) {
	canvas := Canvas()
	html := canvas.ToHTML()

	if strings.Contains(html, "<canvas>") == false {
		t.Error("Does not contain '<canvas>', Output:" + html)
	}

	if strings.Contains(html, "</canvas>") == false {
		t.Error("Does not contain '</canvas>', Output:" + html)
	}
}

func TestCaption(t *testing.T) {
	caption := Caption()
	html := caption.ToHTML()

	if strings.Contains(html, "<caption>") == false {
		t.Error("Does not contain '<caption>', Output:" + html)
	}

	if strings.Contains(html, "</caption>") == false {
		t.Error("Does not contain '</caption>', Output:" + html)
	}
}

func TestCode(t *testing.T) {
	code := Code()
	html := code.ToHTML()

	if strings.Contains(html, "<code>") == false {
		t.Error("Does not contain '<code>', Output:" + html)
	}

	if strings.Contains(html, "</code>") == false {
		t.Error("Does not contain '</code>', Output:" + html)
	}
}

func TestDialog(t *testing.T) {
	e := Dialog()
	html := e.ToHTML()

	if strings.Contains(html, "<dialog>") == false {
		t.Error("Does not contain '<dialog>', Output:" + html)
	}

	if strings.Contains(html, "</dialog>") == false {
		t.Error("Does not contain '</dialog>', Output:" + html)
	}
}

func TestDiv(t *testing.T) {
	div := Div()
	html := div.ToHTML()

	if strings.Contains(html, "<div>") == false {
		t.Error("Does not contain '<div>', Output:" + html)
	}

	if strings.Contains(html, "</div>") == false {
		t.Error("Does not contain '</div>', Output:" + html)
	}
}

func TestFieldset(t *testing.T) {
	e := FieldSet()
	html := e.ToHTML()

	if strings.Contains(html, "<fieldset>") == false {
		t.Error("Does not contain '<fieldset>', Output:" + html)
	}

	if strings.Contains(html, "</fieldset>") == false {
		t.Error("Does not contain '</fieldset>', Output:" + html)
	}
}

func TestFooter(t *testing.T) {
	footer := Footer()
	html := footer.ToHTML()

	if strings.Contains(html, "<footer>") == false {
		t.Error("Does not contain '<footer>', Output:" + html)
	}

	if strings.Contains(html, "</footer>") == false {
		t.Error("Does not contain '</footer>', Output:" + html)
	}
}

func TestForm(t *testing.T) {
	form := Form()
	html := form.ToHTML()

	if strings.Contains(html, "<form>") == false {
		t.Error("Does not contain '<form>', Output:" + html)
	}

	if strings.Contains(html, "</form>") == false {
		t.Error("Does not contain '</form>', Output:" + html)
	}
}

func TestHeader(t *testing.T) {
	header := Header()
	html := header.ToHTML()

	if strings.Contains(html, "<header>") == false {
		t.Error("Does not contain '<header>', Output:" + html)
	}

	if strings.Contains(html, "</header>") == false {
		t.Error("Does not contain '</header>', Output:" + html)
	}
}

func TestH1(t *testing.T) {
	h1 := H1()
	html := h1.ToHTML()

	if strings.Contains(html, "<h1>") == false {
		t.Error("Does not contain '<h1>', Output:" + html)
	}
	if strings.Contains(html, "</h1>") == false {
		t.Error("Does not contain '</h1>', Output:" + html)
	}
}

func TestH2(t *testing.T) {
	h2 := H2()
	html := h2.ToHTML()

	if strings.Contains(html, "<h2>") == false {
		t.Error("Does not contain '<h2>', Output:" + html)
	}
	if strings.Contains(html, "</h2>") == false {
		t.Error("Does not contain '</h2>', Output:" + html)
	}
}

func TestH3(t *testing.T) {
	h3 := H3()
	html := h3.ToHTML()

	if strings.Contains(html, "<h3>") == false {
		t.Error("Does not contain '<h3>', Output:" + html)
	}
	if strings.Contains(html, "</h3>") == false {
		t.Error("Does not contain '</h3>', Output:" + html)
	}
}

func TestH4(t *testing.T) {
	h4 := H4()
	html := h4.ToHTML()

	if strings.Contains(html, "<h4>") == false {
		t.Error("Does not contain '<h4>', Output:" + html)
	}

	if strings.Contains(html, "</h4>") == false {
		t.Error("Does not contain '</h4>', Output:" + html)
	}
}

func TestH5(t *testing.T) {
	h5 := H5()
	html := h5.ToHTML()

	if strings.Contains(html, "<h5>") == false {
		t.Error("Does not contain '<h5>', Output:" + html)
	}
	if strings.Contains(html, "</h5>") == false {
		t.Error("Does not contain '</h5>', Output:" + html)
	}
}

func TestH6(t *testing.T) {
	h6 := H6()
	html := h6.ToHTML()

	if strings.Contains(html, "<h6>") == false {
		t.Error("Does not contain '<h6>', Output:" + html)
	}
	if strings.Contains(html, "</h6>") == false {
		t.Error("Does not contain '</h6>', Output:" + html)
	}
}

func TestHeading1(t *testing.T) {
	h := Heading1()
	html := h.ToHTML()

	if strings.Contains(html, "<h1>") == false {
		t.Error("Does not contain '<h1>', Output:" + html)
	}

	if strings.Contains(html, "</h1>") == false {
		t.Error("Does not contain '</h1>', Output:" + html)
	}
}

func TestHeading2(t *testing.T) {
	h := Heading2()
	html := h.ToHTML()

	if strings.Contains(html, "<h2>") == false {
		t.Error("Does not contain '<h2>', Output:" + html)
	}

	if strings.Contains(html, "</h2>") == false {
		t.Error("Does not contain '</h2>', Output:" + html)
	}
}

func TestHeading3(t *testing.T) {
	h := Heading3()
	html := h.ToHTML()

	if strings.Contains(html, "<h3>") == false {
		t.Error("Does not contain '<h3>', Output:" + html)
	}

	if strings.Contains(html, "</h3>") == false {
		t.Error("Does not contain '</h3>', Output:" + html)
	}
}

func TestHeading4(t *testing.T) {
	h := Heading4()
	html := h.ToHTML()

	if strings.Contains(html, "<h4>") == false {
		t.Error("Does not contain '<h4>', Output:" + html)
	}

	if strings.Contains(html, "</h4>") == false {
		t.Error("Does not contain '</h4>', Output:" + html)
	}
}

func TestHeading5(t *testing.T) {
	h := Heading5()
	html := h.ToHTML()

	if strings.Contains(html, "<h5>") == false {
		t.Error("Does not contain '<h5>', Output:" + html)
	}

	if strings.Contains(html, "</h5>") == false {
		t.Error("Does not contain '</h5>', Output:" + html)
	}
}

func TestHeading6(t *testing.T) {
	h := Heading6()
	html := h.ToHTML()

	if strings.Contains(html, "<h6>") == false {
		t.Error("Does not contain '<h6>', Output:" + html)
	}

	if strings.Contains(html, "</h6>") == false {
		t.Error("Does not contain '</h6>', Output:" + html)
	}
}

func TestHR(t *testing.T) {
	hr := HR()
	html := hr.ToHTML()

	if strings.Contains(html, "<hr />") == false {
		t.Error("Does not contain '<hr />', Output:" + html)
	}
}

func TestHyperlink(t *testing.T) {
	h := Hyperlink()
	html := h.ToHTML()

	if strings.Contains(html, "<a>") == false {
		t.Error("Does not contain '<a>', Output:" + html)
	}

	if strings.Contains(html, "</a>") == false {
		t.Error("Does not contain '</a>', Output:" + html)
	}
}

func TestI(t *testing.T) {
	i := I()
	html := i.ToHTML()

	if strings.Contains(html, "<i>") == false {
		t.Error("Does not contain '<i>', Output:" + html)
	}
	if strings.Contains(html, "</i>") == false {
		t.Error("Does not contain '</i>', Output:" + html)
	}
}

func TestIframe(t *testing.T) {
	i := Iframe(`url`)
	html := i.ToHTML()

	if strings.Contains(html, "<iframe src=\"url\">") == false {
		t.Error("Does not contain '<iframe src=\"url\">', Output:" + html)
	}
	if strings.Contains(html, "</iframe>") == false {
		t.Error("Does not contain '</iframe>', Output:" + html)
	}
}

func TestImg(t *testing.T) {
	img := Img("url")
	html := img.ToHTML()

	if strings.Contains(html, "<img src=\"url\" />") == false {
		t.Error("Does not contain '<img src=\"url\" />', Output:" + html)
	}
}

func TestImage(t *testing.T) {
	img := Image("url")
	html := img.ToHTML()

	if strings.Contains(html, "<img src=\"url\" />") == false {
		t.Error("Does not contain '<img src=\"url\" />', Output:" + html)
	}
}

func TestInput(t *testing.T) {
	input := Input()
	html := input.ToHTML()

	if strings.Contains(html, "<input />") == false {
		t.Error("Does not contain '<input />', Output:" + html)
	}
}

func TestLabel(t *testing.T) {
	label := Label()
	html := label.ToHTML()

	if strings.Contains(html, "<label>") == false {
		t.Error("Does not contain '<label>', Output:" + html)
	}
	if strings.Contains(html, "</label>") == false {
		t.Error("Does not contain '</label>', Output:" + html)
	}
}

func TestLI(t *testing.T) {
	li := LI()
	html := li.ToHTML()

	if strings.Contains(html, "<li>") == false {
		t.Error("Does not contain '<li>', Output:" + html)
	}
	if strings.Contains(html, "</li>") == false {
		t.Error("Does not contain '</li>', Output:" + html)
	}
}

func TestLink(t *testing.T) {
	link := Link()
	html := link.ToHTML()

	if strings.Contains(html, "<link />") == false {
		t.Error("Does not contain '<link />', Output:" + html)
	}
}

func TestMain(t *testing.T) {
	main := Main()
	html := main.ToHTML()

	if strings.Contains(html, "<main>") == false {
		t.Error("Does not contain '<main>', Output:" + html)
	}
	if strings.Contains(html, "</main>") == false {
		t.Error("Does not contain '</main>', Output:" + html)
	}
}

func TestMeta(t *testing.T) {
	meta := Meta()
	html := meta.ToHTML()

	if strings.Contains(html, "<meta />") == false {
		t.Error("Does not contain '<meta />', Output:" + html)
	}
}

func TestNav(t *testing.T) {
	nav := Nav()
	html := nav.ToHTML()

	if strings.Contains(html, "<nav>") == false {
		t.Error("Does not contain '<nav>', Output:" + html)
	}
	if strings.Contains(html, "</nav>") == false {
		t.Error("Does not contain '</nav>', Output:" + html)
	}
}

func TestNavbar(t *testing.T) {
	navbar := Navbar()
	html := navbar.ToHTML()

	if strings.Contains(html, "<navbar>") == false {
		t.Error("Does not contain '<navbar>', Output:" + html)
	}
	if strings.Contains(html, "</navbar>") == false {
		t.Error("Does not contain '</navbar>', Output:" + html)
	}
}

func TestOL(t *testing.T) {
	ol := OL()
	html := ol.ToHTML()

	if strings.Contains(html, "<ol>") == false {
		t.Error("Does not contain '<ol>', Output:" + html)
	}
	if strings.Contains(html, "</ol>") == false {
		t.Error("Does not contain '</ol>', Output:" + html)
	}
}

func TestOption(t *testing.T) {
	option := Option()
	html := option.ToHTML()

	if strings.Contains(html, "<option>") == false {
		t.Error("Does not contain '<option>', Output:" + html)
	}
	if strings.Contains(html, "</option>") == false {
		t.Error("Does not contain '</option>', Output:" + html)
	}
}

func TestP(t *testing.T) {
	p := P()
	html := p.ToHTML()

	if strings.Contains(html, "<p>") == false {
		t.Error("Does not contain '<p>', Output:" + html)
	}
	if strings.Contains(html, "</p>") == false {
		t.Error("Does not contain '</p>', Output:" + html)
	}
}

func TestParagraph(t *testing.T) {
	p := Paragraph()
	html := p.ToHTML()

	if strings.Contains(html, "<p>") == false {
		t.Error("Does not contain '<p>', Output:" + html)
	}
	if strings.Contains(html, "</p>") == false {
		t.Error("Does not contain '</p>', Output:" + html)
	}
}

func TestPRE(t *testing.T) {
	pre := PRE()
	html := pre.ToHTML()

	if strings.Contains(html, "<pre>") == false {
		t.Error("Does not contain '<pre>', Output:" + html)
	}
	if strings.Contains(html, "</pre>") == false {
		t.Error("Does not contain '</pre>', Output:" + html)
	}
}

func TestRaw(t *testing.T) {
	s := Raw("<script>alert('Hello')</script>")
	html := s.ToHTML()

	if strings.Contains(html, "<script>") == false {
		t.Error("Does not contain '<script>', Output:" + html)
	}

	if strings.Contains(html, "alert('Hello')") == false {
		t.Error("Does not contain 'alert('Hello')', Output:" + html)
	}

	if strings.Contains(html, "</script>") == false {
		t.Error("Does not contain '</script>', Output:" + html)
	}
}

func TestScript(t *testing.T) {
	s := Script(`window.alert('Hello')`)
	html := s.ToHTML()

	if strings.Contains(html, "<script>") == false {
		t.Error("Does not contain '<script>', Output:" + html)
	}
	if strings.Contains(html, "</script>") == false {
		t.Error("Does not contain '</script>', Output:" + html)
	}

	if strings.Contains(html, "window.alert('Hello')") == false {
		t.Error("Does not contain 'window.alert('Hello')', Output:" + html)
	}
}

func TestScriptURL(t *testing.T) {
	s := ScriptURL("https://example.com/script.js")
	html := s.ToHTML()

	if strings.Contains(html, "<script src=\"https://example.com/script.js\"></script>") == false {
		t.Error("Does not contain '<script src=\"https://example.com/script.js\"></script>', Output:" + html)
	}
}

func TestSection(t *testing.T) {
	s := Section()
	html := s.ToHTML()

	if strings.Contains(html, "<section>") == false {
		t.Error("Does not contain '<section>', Output:" + html)
	}
	if strings.Contains(html, "</section>") == false {
		t.Error("Does not contain '</section>', Output:" + html)
	}
}

func TestSelect(t *testing.T) {
	s := Select()
	html := s.ToHTML()

	if strings.Contains(html, "<select>") == false {
		t.Error("Does not contain '<select>', Output:" + html)
	}
	if strings.Contains(html, "</select>") == false {
		t.Error("Does not contain '</select>', Output:" + html)
	}
}

func TestSmall(t *testing.T) {
	s := Small()
	html := s.ToHTML()

	if strings.Contains(html, "<small>") == false {
		t.Error("Does not contain '<small>', Output:" + html)
	}
	if strings.Contains(html, "</small>") == false {
		t.Error("Does not contain '</small>', Output:" + html)
	}
}

func TestSpan(t *testing.T) {
	s := Span()
	html := s.ToHTML()

	if strings.Contains(html, "<span>") == false {
		t.Error("Does not contain '<span>', Output:" + html)
	}
	if strings.Contains(html, "</span>") == false {
		t.Error("Does not contain '</span>', Output:" + html)
	}
}

func TestStrong(t *testing.T) {
	s := Strong()
	html := s.ToHTML()

	if strings.Contains(html, "<strong>") == false {
		t.Error("Does not contain '<strong>', Output:" + html)
	}

	if strings.Contains(html, "</strong>") == false {
		t.Error("Does not contain '</strong>', Output:" + html)
	}
}

func TestStyle(t *testing.T) {
	s := Style(`color: red;`)
	html := s.ToHTML()

	if strings.Contains(html, "<style>") == false {
		t.Error("Does not contain '<style>', Output:" + html)
	}
	if strings.Contains(html, "</style>") == false {
		t.Error("Does not contain '</style>', Output:" + html)
	}

	if strings.Contains(html, "color: red;") == false {
		t.Error("Does not contain 'color: red;', Output:" + html)
	}
}

func TestStyleURL(t *testing.T) {
	s := StyleURL("https://example.com/style.css")
	html := s.ToHTML()

	if strings.Contains(html, "<link href=\"https://example.com/style.css\" rel=\"stylesheet\" />") == false {
		t.Error("Does not contain '<link href=\"https://example.com/style.css\" rel=\"stylesheet\" />', Output:" + html)
	}
}

func TestSub(t *testing.T) {
	s := Sub()
	html := s.ToHTML()

	if strings.Contains(html, "<sub>") == false {
		t.Error("Does not contain '<sub>', Output:" + html)
	}

	if strings.Contains(html, "</sub>") == false {
		t.Error("Does not contain '</sub>', Output:" + html)
	}
}

func TestSup(t *testing.T) {
	s := Sup()
	html := s.ToHTML()

	if strings.Contains(html, "<sup>") == false {
		t.Error("Does not contain '<sup>', Output:" + html)
	}

	if strings.Contains(html, "</sup>") == false {
		t.Error("Does not contain '</sup>', Output:" + html)
	}
}

func TestSwal(t *testing.T) {
	s := Swal(SwalOptions{Title: "Hello", Text: "World"})
	html := s.ToHTML()

	if strings.Contains(html, `Swal.fire({"showCancelButton":false,"showConfirmButton":false,"showDenyButton":false,"text":"World","title":"Hello"});`) == false {
		t.Error(`Does not contain 'Swal.fire({"showCancelButton":false,"showConfirmButton":false,"showDenyButton":false,"text":"World","title":"Hello"});', Output:` + html)
	}
}

func TestTemplate(t *testing.T) {
	tmpl := Template()
	html := tmpl.ToHTML()

	if strings.Contains(html, "<template>") == false {
		t.Error("Does not contain '<template>', Output:" + html)
	}
	if strings.Contains(html, "</template>") == false {
		t.Error("Does not contain '</template>', Output:" + html)
	}
}

func TestText(t *testing.T) {
	s := Text("<script>alert('Hello')</script>")
	html := s.ToHTML()

	if strings.Contains(html, "&lt;script&gt;") == false {
		t.Error("Does not contain '&lt;script&gt;', Output:" + html)
	}

	if strings.Contains(html, "alert(&#39;Hello&#39;)") == false {
		t.Error("Does not contain 'alert(&#39;Hello&#39;)', Output:" + html)
	}

	if strings.Contains(html, "&lt;/script&gt") == false {
		t.Error("Does not contain '&lt;/script&gt', Output:" + html)
	}
}

func TestTextArea(t *testing.T) {
	s := TextArea()
	html := s.ToHTML()

	if strings.Contains(html, "<textarea>") == false {
		t.Error("Does not contain '<textarea>', Output:" + html)
	}
	if strings.Contains(html, "</textarea>") == false {
		t.Error("Does not contain '</textarea>', Output:" + html)
	}
}

func TestTable(t *testing.T) {
	s := Table()
	html := s.ToHTML()

	if strings.Contains(html, "<table>") == false {
		t.Error("Does not contain '<table>', Output:" + html)
	}
	if strings.Contains(html, "</table>") == false {
		t.Error("Does not contain '</table>', Output:" + html)
	}
}

func TestTbody(t *testing.T) {
	s := Tbody()
	html := s.ToHTML()

	if strings.Contains(html, "<tbody>") == false {
		t.Error("Does not contain '<tbody>', Output:" + html)
	}
	if strings.Contains(html, "</tbody>") == false {
		t.Error("Does not contain '</tbody>', Output:" + html)
	}
}

func TestTD(t *testing.T) {
	s := TD()
	html := s.ToHTML()

	if strings.Contains(html, "<td>") == false {
		t.Error("Does not contain '<td>', Output:" + html)
	}
	if strings.Contains(html, "</td>") == false {
		t.Error("Does not contain '</td>', Output:" + html)
	}
}

func TestTfoot(t *testing.T) {
	s := Tfoot()
	html := s.ToHTML()

	if strings.Contains(html, "<tfoot>") == false {
		t.Error("Does not contain '<tfoot>', Output:" + html)
	}
	if strings.Contains(html, "</tfoot>") == false {
		t.Error("Does not contain '</tfoot>', Output:" + html)
	}
}

func TestTH(t *testing.T) {
	s := TH()
	html := s.ToHTML()

	if strings.Contains(html, "<th>") == false {
		t.Error("Does not contain '<th>', Output:" + html)
	}
	if strings.Contains(html, "</th>") == false {
		t.Error("Does not contain '</th>', Output:" + html)
	}
}

func TestThead(t *testing.T) {
	s := Thead()
	html := s.ToHTML()

	if strings.Contains(html, "<thead>") == false {
		t.Error("Does not contain '<thead>', Output:" + html)
	}
	if strings.Contains(html, "</thead>") == false {
		t.Error("Does not contain '</thead>', Output:" + html)
	}
}

func TestTR(t *testing.T) {
	s := TR()
	html := s.ToHTML()

	if strings.Contains(html, "<tr>") == false {
		t.Error("Does not contain '<tr>', Output:" + html)
	}
	if strings.Contains(html, "</tr>") == false {
		t.Error("Does not contain '</tr>', Output:" + html)
	}
}

func TestTitle(t *testing.T) {
	e := Title()
	html := e.ToHTML()

	if strings.Contains(html, "<title>") == false {
		t.Error("Does not contain '<title>', Output:" + html)
	}

	if strings.Contains(html, "</title>") == false {
		t.Error("Does not contain '</title>', Output:" + html)
	}
}

func TestUL(t *testing.T) {
	s := UL()
	html := s.ToHTML()

	if strings.Contains(html, "<ul>") == false {
		t.Error("Does not contain '<ul>', Output:" + html)
	}
	if strings.Contains(html, "</ul>") == false {
		t.Error("Does not contain '</ul>', Output:" + html)
	}
}

func TestWebpage(t *testing.T) {
	s := Webpage()
	html := s.ToHTML()

	prefix := `<!DOCTYPE html><html><head><meta charset="utf-8" />`

	if !strings.HasPrefix(html, prefix) {
		t.Error("Does not start with: ", prefix, ", Output: ", html)
	}

	suffix := `</head><body></body></html>`

	if !strings.HasSuffix(html, suffix) {
		t.Error("Does not end with: ", suffix, ", Output: ", html)
	}
}

func TestWrap(t *testing.T) {
	w := Wrap()
	html := w.ToHTML()

	if html != "" {
		t.Error("Expected empty string for Wrap(), Output:", html)
	}
}
