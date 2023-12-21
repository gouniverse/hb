package hb

import (
	"strings"
	"testing"
)

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
