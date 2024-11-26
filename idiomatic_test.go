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
