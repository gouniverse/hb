package hb

import (
	"strconv"
)

const BORDER_LAYOUT_ALIGN_BOTTOM = "bottom"
const BORDER_LAYOUT_ALIGN_CENTER = "center"
const BORDER_LAYOUT_ALIGN_LEFT = "left"
const BORDER_LAYOUT_ALIGN_MIDDLE = "middle"
const BORDER_LAYOUT_ALIGN_RIGHT = "right"
const BORDER_LAYOUT_ALIGN_TOP = "top"

var _ TagInterface = (*BorderLayout)(nil)

func NewBorderLayout() *BorderLayout {
	return &BorderLayout{}
}

/**
* The BorderLayout type is an advanced layout.
*
* It arranges its children in five regions: top, bottom, left, right, and center.
* Each region may contain no more than one widget.
* <code>
* // Creating a new instance of BorderLayout
* borderlayout = NewBorderLayout()
*     AddChild("HEADER", BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE);
*     AddChild("CONTENT", BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE);
*     AddChild("FOOTER", BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE);
*     AddChild("MENU", BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE);
*     AddChild("ADDS", BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE);
* </code>
 */
type BorderLayout struct {
	Tag
	width                 string
	height                string
	top                   TagInterface
	topAlignHorizontal    string
	topAlignVertical      string
	bottom                TagInterface
	bottomAlignHorizontal string
	bottomAlignVertical   string
	left                  TagInterface
	leftAlignHorizontal   string
	leftAlignVertical     string
	right                 TagInterface
	rightAlignHorizontal  string
	rightAlignVertical    string
	center                TagInterface
	centerAlignHorizontal string
	centerAlignVertical   string
}

func (bl *BorderLayout) applyDefaultAlignments() {
	if bl.height == "" {
		bl.height = "100%"
	}

	if bl.width == "" {
		bl.width = "100%"
	}

	if bl.centerAlignHorizontal == "" {
		bl.centerAlignHorizontal = BORDER_LAYOUT_ALIGN_CENTER
	}

	if bl.centerAlignVertical == "" {
		bl.centerAlignVertical = BORDER_LAYOUT_ALIGN_MIDDLE
	}

	if bl.topAlignHorizontal == "" {
		bl.topAlignHorizontal = BORDER_LAYOUT_ALIGN_CENTER
	}

	if bl.topAlignVertical == "" {
		bl.topAlignVertical = BORDER_LAYOUT_ALIGN_MIDDLE
	}

	if bl.bottomAlignHorizontal == "" {
		bl.bottomAlignHorizontal = BORDER_LAYOUT_ALIGN_CENTER
	}

	if bl.bottomAlignVertical == "" {
		bl.bottomAlignVertical = BORDER_LAYOUT_ALIGN_MIDDLE
	}

	if bl.leftAlignHorizontal == "" {
		bl.leftAlignHorizontal = BORDER_LAYOUT_ALIGN_CENTER
	}

	if bl.leftAlignVertical == "" {
		bl.leftAlignVertical = BORDER_LAYOUT_ALIGN_MIDDLE
	}

	if bl.rightAlignHorizontal == "" {
		bl.rightAlignHorizontal = BORDER_LAYOUT_ALIGN_CENTER
	}

	if bl.rightAlignVertical == "" {
		bl.rightAlignVertical = BORDER_LAYOUT_ALIGN_MIDDLE
	}
}

func (bl *BorderLayout) AddTop(tag TagInterface, alignHorizontal, alignVertical string) *BorderLayout {
	bl.top = tag
	bl.topAlignHorizontal = alignHorizontal
	bl.topAlignVertical = alignVertical

	return bl
}

func (bl *BorderLayout) AddBottom(tag TagInterface, alignHorizontal, alignVertical string) *BorderLayout {
	bl.bottom = tag
	bl.bottomAlignHorizontal = alignHorizontal
	bl.bottomAlignVertical = alignVertical

	return bl
}

func (bl *BorderLayout) AddLeft(tag TagInterface, alignHorizontal, alignVertical string) *BorderLayout {
	bl.left = tag
	bl.leftAlignHorizontal = alignHorizontal
	bl.leftAlignVertical = alignVertical

	return bl
}

func (bl *BorderLayout) AddRight(tag TagInterface, alignHorizontal, alignVertical string) *BorderLayout {
	bl.right = tag
	bl.rightAlignHorizontal = alignHorizontal
	bl.rightAlignVertical = alignVertical

	return bl
}

func (bl *BorderLayout) AddCenter(tag TagInterface, alignHorizontal, alignVertical string) *BorderLayout {
	bl.center = tag
	bl.centerAlignHorizontal = alignHorizontal
	bl.centerAlignVertical = alignVertical

	return bl
}

func (bl *BorderLayout) SetHeightPercents(height int) *BorderLayout {
	bl.height = strconv.Itoa(height) + "%"

	return bl
}

func (bl *BorderLayout) SetWidthPercents(width int) *BorderLayout {
	bl.width = strconv.Itoa(width) + "%"

	return bl
}

func (bl *BorderLayout) SetHeightPixels(height int) *BorderLayout {
	bl.height = strconv.Itoa(height) + "px"

	return bl
}

func (bl *BorderLayout) SetWidthPixels(width int) *BorderLayout {
	bl.width = strconv.Itoa(width) + "px"

	return bl
}

// BorderLayout returns HTML representation of the layout
func (bl *BorderLayout) ToHTML() string {
	bl.applyDefaultAlignments()

	colspan := 0
	if bl.left != nil {
		colspan++
	}
	if bl.center != nil {
		colspan++
	}
	if bl.right != nil {
		colspan++
	}

	table := NewTable()
	table.TagAttributes = bl.TagAttributes
	table.Class("BorderLayout").
		Attr("cellspacing", "0").
		Attr("cellpadding", "0").
		Style("width:" + bl.width + ";height:" + bl.height + ";")

	if bl.top != nil {
		td := NewTD().Class("Top").
			Style("height:1px;text-align:" + bl.topAlignHorizontal + ";vertical-align:" + bl.topAlignVertical).
			AddChild(bl.top)

		if colspan > 1 {
			td.Attr("colspan", strconv.Itoa(colspan))
		}

		tr := NewTR().AddChild(td)
		table.AddChild(tr)
	}

	tr := NewTR()

	if bl.left != nil {
		td := NewTD().Class("Left").
			Style("width:1px;height:100%;text-align:" + bl.leftAlignHorizontal + ";vertical-align:" + bl.leftAlignVertical).
			AddChild(bl.left)

		tr.AddChild(td)
	}

	if bl.center != nil {
		td := NewTD().Class("Center").
			Style("text-align:" + bl.centerAlignHorizontal + ";vertical-align:" + bl.centerAlignVertical).
			AddChild(bl.center)

		tr.AddChild(td)
	}

	if bl.right != nil {
		td := NewTD().Class("Right").
			Style("width:1px;height:100%;text-align:" + bl.rightAlignHorizontal + ";vertical-align:" + bl.rightAlignVertical).
			AddChild(bl.right)

		tr.AddChild(td)
	}

	table.AddChild(tr)

	if bl.bottom != nil {
		td := NewTD().Class("Bottom").
			Style("height:1px;text-align:" + bl.bottomAlignHorizontal + ";vertical-align:" + bl.bottomAlignVertical).
			AddChild(bl.bottom)

		if colspan > 1 {
			td.Attr("colspan", strconv.Itoa(colspan))
		}

		tr := NewTR().AddChild(td)
		table.AddChild(tr)
	}

	return table.ToHTML()
}
