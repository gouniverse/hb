package hb

import (
	"strings"
	"testing"
)

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
