package hb

import (
	"testing"
)

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
