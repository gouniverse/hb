package hb

import (
	"strings"
	"testing"
)

func TestNewBR(t *testing.T) {
	br := NewBR()
	html := br.ToHTML()

	if strings.Contains(html, "<br />") == false {
		t.Error("Does not contain '<br />', Output:" + html)
	}
}

func TestNewButton(t *testing.T) {
	header := NewButton()
	html := header.ToHTML()

	if strings.Contains(html, "<button>") == false {
		t.Error("Does not contain '<button>', Output:" + html)
	}

	if strings.Contains(html, "</button>") == false {
		t.Error("Does not contain '</button>', Output:" + html)
	}
}

func TestNewCaption(t *testing.T) {
	caption := NewCaption()
	html := caption.ToHTML()

	if strings.Contains(html, "<caption>") == false {
		t.Error("Does not contain '<caption>', Output:" + html)
	}

	if strings.Contains(html, "</caption>") == false {
		t.Error("Does not contain '</caption>', Output:" + html)
	}
}

func TestNewCode(t *testing.T) {
	code := NewCode()
	html := code.ToHTML()

	if strings.Contains(html, "<code>") == false {
		t.Error("Does not contain '<code>', Output:" + html)
	}

	if strings.Contains(html, "</code>") == false {
		t.Error("Does not contain '</code>', Output:" + html)
	}
}

func TestNewDiv(t *testing.T) {
	header := NewDiv()
	html := header.ToHTML()

	if strings.Contains(html, "<div>") == false {
		t.Error("Does not contain '<div>', Output:" + html)
	}

	if strings.Contains(html, "</div>") == false {
		t.Error("Does not contain '</div>', Output:" + html)
	}
}

func TestNewFooter(t *testing.T) {
	footer := NewFooter()
	html := footer.ToHTML()

	if strings.Contains(html, "<footer>") == false {
		t.Error("Does not contain '<footer>', Output:" + html)
	}

	if strings.Contains(html, "</footer>") == false {
		t.Error("Does not contain '</footer>', Output:" + html)
	}
}

func TestNewForm(t *testing.T) {
	header := NewForm()
	html := header.ToHTML()

	if strings.Contains(html, "<form>") == false {
		t.Error("Does not contain '<form>', Output:" + html)
	}

	if strings.Contains(html, "</form>") == false {
		t.Error("Does not contain '</form>', Output:" + html)
	}
}

func TestNewHeader(t *testing.T) {
	header := NewHeader()
	html := header.ToHTML()

	if strings.Contains(html, "<header>") == false {
		t.Error("Does not contain '<header>', Output:" + html)
	}

	if strings.Contains(html, "</header>") == false {
		t.Error("Does not contain '</header>', Output:" + html)
	}
}

func TestNewHeading1(t *testing.T) {
	h1 := NewHeading1()
	html := h1.ToHTML()

	expecteds := []string{"<h1>", "</h1>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHeading2(t *testing.T) {
	h2 := NewHeading2()
	html := h2.ToHTML()

	expecteds := []string{"<h2>", "</h2>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHeading3(t *testing.T) {
	h3 := NewHeading3()
	html := h3.ToHTML()

	expecteds := []string{"<h3>", "</h3>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHeading4(t *testing.T) {
	h4 := NewHeading4()
	html := h4.ToHTML()

	expecteds := []string{"<h4>", "</h4>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHeading5(t *testing.T) {
	h5 := NewHeading5()
	html := h5.ToHTML()

	expecteds := []string{"<h5>", "</h5>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHeading6(t *testing.T) {
	h6 := NewHeading6()
	html := h6.ToHTML()

	expecteds := []string{"<h6>", "</h6>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH1(t *testing.T) {
	h1 := NewH1()
	html := h1.ToHTML()

	expecteds := []string{"<h1>", "</h1>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH2(t *testing.T) {
	h2 := NewH2()
	html := h2.ToHTML()

	expecteds := []string{"<h2>", "</h2>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH3(t *testing.T) {
	h3 := NewH3()
	html := h3.ToHTML()

	expecteds := []string{"<h3>", "</h3>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH4(t *testing.T) {
	h4 := NewH4()
	html := h4.ToHTML()

	expecteds := []string{"<h4>", "</h4>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH5(t *testing.T) {
	h5 := NewH5()
	html := h5.ToHTML()

	expecteds := []string{"<h5>", "</h5>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewH6(t *testing.T) {
	h6 := NewH6()
	html := h6.ToHTML()

	expecteds := []string{"<h6>", "</h6>"}
	for _, expected := range expecteds {
		if strings.Contains(html, expected) == false {
			t.Error("Does not contain '" + expected + "', Output:" + html)
		}
	}
}

func TestNewHR(t *testing.T) {
	header := NewHR()
	html := header.ToHTML()

	if strings.Contains(html, "<hr />") == false {
		t.Error("Does not contain '<hr />', Output:" + html)
	}
}

func TestNewHTML(t *testing.T) {
	s := NewHTML("<script>alert('Hello')</script>")
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

func TestNewHyperlink(t *testing.T) {
	link := NewHyperlink().
		Href("https://example.com").
		Text("Example")
	html := link.ToHTML()

	if strings.Contains(html, "<a href=\"https://example.com\">Example</a>") == false {
		t.Error("Does not contain '<a href=\"https://example.com\">Example</a>', Output:" + html)
	}
}

func TestNewI(t *testing.T) {
	i := NewI()
	html := i.ToHTML()

	if strings.Contains(html, "<i>") == false {
		t.Error("Does not contain '<i>', Output:" + html)
	}
	if strings.Contains(html, "</i>") == false {
		t.Error("Does not contain '</i>', Output:" + html)
	}
}

func TestNewIWithContent(t *testing.T) {
	i := NewI().HTML("italic")
	html := i.ToHTML()

	if strings.Contains(html, "<i>") == false {
		t.Error("Does not contain '<i>', Output:" + html)
	}
	if strings.Contains(html, "</i>") == false {
		t.Error("Does not contain '</i>', Output:" + html)
	}
	if strings.Contains(html, "italic") == false {
		t.Error("Does not contain 'italic', Output:" + html)
	}
}

func TestNewImage(t *testing.T) {
	s := NewImage()
	html := s.ToHTML()

	if strings.Contains(html, "<img />") == false {
		t.Error("Does not contain '<img />', Output:" + html)
	}
}

func TestNewInput(t *testing.T) {
	s := NewInput()
	html := s.ToHTML()

	if strings.Contains(html, "<input />") == false {
		t.Error("Does not contain '<input />', Output:" + html)
	}

	if strings.Contains(html, "</input>") {
		t.Error("Must not contain '</input>' as it is a self closing tag, Output:" + html)
	}
}

func TestNewLabel(t *testing.T) {
	s := NewLabel()
	html := s.ToHTML()

	if strings.Contains(html, "<label>") == false {
		t.Error("Does not contain '<label>', Output:" + html)
	}

	if strings.Contains(html, "</label>") == false {
		t.Error("Does not contain '</label>', Output:" + html)
	}
}

func TestNewLI(t *testing.T) {
	li := NewLI()
	html := li.ToHTML()

	if strings.Contains(html, "<li>") == false {
		t.Error("Does not contain '<li>', Output:" + html)
	}

	if strings.Contains(html, "</li>") == false {
		t.Error("Does not contain '</li>', Output:" + html)
	}
}

func TestNewLink(t *testing.T) {
	link := NewLink()
	html := link.ToHTML()

	if strings.Contains(html, "<link />") == false {
		t.Error("Does not contain '<link />', Output:" + html)
	}
}

func TestNewMain(t *testing.T) {
	main := NewMain()
	html := main.ToHTML()

	if strings.Contains(html, "<main>") == false {
		t.Error("Does not contain '<main>', Output:" + html)
	}

	if strings.Contains(html, "</main>") == false {
		t.Error("Does not contain '</main>', Output:" + html)
	}
}

func TestNewMeta(t *testing.T) {
	meta := NewMeta()
	html := meta.ToHTML()

	if strings.Contains(html, "<meta />") == false {
		t.Error("Does not contain '<meta />', Output:" + html)
	}
}

func TestNewNav(t *testing.T) {
	nav := NewNav()
	html := nav.ToHTML()

	if strings.Contains(html, "<nav>") == false {
		t.Error("Does not contain '<nav>', Output:" + html)
	}

	if strings.Contains(html, "</nav>") == false {
		t.Error("Does not contain '</nav>', Output:" + html)
	}
}

func TestNewNavbar(t *testing.T) {
	navbar := NewNavbar()
	html := navbar.ToHTML()

	if strings.Contains(html, "<navbar>") == false {
		t.Error("Does not contain '<navbar>', Output:" + html)
	}

	if strings.Contains(html, "</navbar>") == false {
		t.Error("Does not contain '</navbar>', Output:" + html)
	}
}

func TestNewOption(t *testing.T) {
	e := NewOption()
	html := e.ToHTML()

	if strings.Contains(html, "<option>") == false {
		t.Error("Does not contain '<option>', Output:" + html)
	}

	if strings.Contains(html, "</option>") == false {
		t.Error("Does not contain '</option>', Output:" + html)
	}
}

func TestNewOptionEmptyValue(t *testing.T) {
	e := NewOption().Value("").HTML("- select -")
	html := e.ToHTML()

	if strings.Contains(html, `<option value="">- select -</option>`) == false {
		t.Error(`Does not contain '<option value="">- select -</option>', Output:` + html)
	}
}

// TestPod is a unit test for the NewPod and ToHTML methods of the Pod struct.
//
// Args:
//
//	t (testing.T): A testing object provided by the testing package.
//
// Returns:
//
//	None.
//
// Raises:
//
//	AssertionError: If the output of ToHTML method is not an empty string.
//
// Example:
//
//	>>> pod := NewPod()
//	>>> html := pod.ToHTML()
//	>>> assert(html == "")
func TestPod(t *testing.T) {
	pod := NewPod()
	html := pod.ToHTML()

	if html != "" {
		t.Error("Does not equal '' empty string, Output:" + html)
	}

	pod2 := NewPod().Children([]TagInterface{
		NewBR(),
		NewI(),
	})
	html2 := pod2.ToHTML()

	if html2 != "<br /><i></i>" {
		t.Error("Does not equal '<br /><i></i>', Output:" + html2)
	}
}

func TestNewSection(t *testing.T) {
	s := NewSection()
	html := s.ToHTML()

	if strings.Contains(html, "<section>") == false {
		t.Error("Does not contain '<section>', Output:" + html)
	}

	if strings.Contains(html, "</section>") == false {
		t.Error("Does not contain '</section>', Output:" + html)
	}
}

func TestNewSelect(t *testing.T) {
	s := NewSelect()
	html := s.ToHTML()

	if strings.Contains(html, "<select>") == false {
		t.Error("Does not contain '<select>', Output:" + html)
	}

	if strings.Contains(html, "</select>") == false {
		t.Error("Does not contain '</select>', Output:" + html)
	}
}

func TestTagTable(t *testing.T) {
	tag := NewTable()
	h := tag.ToHTML()
	if strings.Contains(h, "<table></table>") == false {
		t.Error("Does not contain '<table></table>'", "Output:"+h)
	}
}

func TestTagTBody(t *testing.T) {
	tag := NewTbody()
	h := tag.ToHTML()
	if strings.Contains(h, "<tbody></tbody>") == false {
		t.Error("Does not contain '<tbody></tbody>'", "Output:"+h)
	}
}

func TestTagThead(t *testing.T) {
	tag := NewThead()
	h := tag.ToHTML()
	if strings.Contains(h, "<thead></thead>") == false {
		t.Error("Does not contain '<thead></thead>'", "Output:"+h)
	}
}

func TestTagTR(t *testing.T) {
	tag := NewTR()
	h := tag.ToHTML()
	if strings.Contains(h, "<tr></tr>") == false {
		t.Error("Does not contain '<tr></tr>'", "Output:"+h)
	}
}

func TestTagTH(t *testing.T) {
	tag := NewTH()
	h := tag.ToHTML()
	if strings.Contains(h, "<th></th>") == false {
		t.Error("Does not contain '<th></th>'", "Output:"+h)
	}
}

func TestTagTD(t *testing.T) {
	tag := NewTD()
	h := tag.ToHTML()
	if strings.Contains(h, "<td></td>") == false {
		t.Error("Does not contain '<td></td>'", "Output:"+h)
	}
}

func TestNewTag(t *testing.T) {
	s := NewTag(`x-login-form`)
	html := s.ToHTML()

	if strings.Contains(html, "<x-login-form>") == false {
		t.Error("Does not contain '<x-login-form>', Output:" + html)
	}

	if strings.Contains(html, "</x-login-form>") == false {
		t.Error("Does not contain '</x-login-form>', Output:" + html)
	}
}

func TestNewTemplate(t *testing.T) {
	s := NewTemplate()
	html := s.ToHTML()

	if strings.Contains(html, "<template>") == false {
		t.Error("Does not contain '<template>', Output:" + html)
	}

	if strings.Contains(html, "</template>") == false {
		t.Error("Does not contain '</template>', Output:" + html)
	}
}

func TestNewText(t *testing.T) {
	s := NewText("<script>alert('Hello')</script>")
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

func TestNewTextArea(t *testing.T) {
	e := NewTextArea()
	html := e.ToHTML()

	if strings.Contains(html, "<textarea>") == false {
		t.Error("Does not contain '<textarea>', Output:" + html)
	}

	if strings.Contains(html, "</textarea>") == false {
		t.Error("Does not contain '</textarea>', Output:" + html)
	}
}

func TestNewWrap(t *testing.T) {
	wrap := NewWrap().Children([]TagInterface{
		NewImage().Src("foo.jpg"),
		NewBR(),
		NewDiv().HTML("Image description"),
	})
	html := wrap.ToHTML()

	expected := `<img src="foo.jpg" /><br /><div>Image description</div>`
	if html != expected {
		t.Error(`Must be exactly '`+expected+`', Output: `, html)
	}
}
