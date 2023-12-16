package hb

import (
	"strings"
	"testing"
)

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
