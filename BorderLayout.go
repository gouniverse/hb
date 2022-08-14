package hb

import (
	"github.com/gouniverse/utils"
)

const BORDER_LAYOUT_ALIGN_BOTTOM = "bottom"
const BORDER_LAYOUT_ALIGN_CENTER = "center"
const BORDER_LAYOUT_ALIGN_LEFT = "left"
const BORDER_LAYOUT_ALIGN_MIDDLE = "middle"
const BORDER_LAYOUT_ALIGN_RIGHT = "right"
const BORDER_LAYOUT_ALIGN_TOP = "top"

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
	top                   *Tag
	topAlignHorizontal    string
	topAlignVertical      string
	bottom                *Tag
	bottomAlignHorizontal string
	bottomAlignVertical   string
	left                  *Tag
	leftAlignHorizontal   string
	leftAlignVertical     string
	right                 *Tag
	rightAlignHorizontal  string
	rightAlignVertical    string
	center                *Tag
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

func (bl *BorderLayout) AddTop(tag *Tag, alignHorizontal string, alignVertical string) {
	bl.top = tag
	bl.topAlignHorizontal = alignHorizontal
	bl.topAlignVertical = alignVertical
}

func (bl *BorderLayout) AddBottom(tag *Tag, alignHorizontal string, alignVertical string) {
	bl.bottom = tag
	bl.bottomAlignHorizontal = alignHorizontal
	bl.bottomAlignVertical = alignVertical
}

func (bl *BorderLayout) AddLeft(tag *Tag, alignHorizontal string, alignVertical string) {
	bl.left = tag
	bl.leftAlignHorizontal = alignHorizontal
	bl.leftAlignVertical = alignVertical
}

func (bl *BorderLayout) AddRight(tag *Tag, alignHorizontal string, alignVertical string) {
	bl.right = tag
	bl.rightAlignHorizontal = alignHorizontal
	bl.rightAlignVertical = alignVertical
}

func (bl *BorderLayout) AddCenter(tag *Tag, alignHorizontal string, alignVertical string) {
	bl.center = tag
	bl.centerAlignHorizontal = alignHorizontal
	bl.centerAlignVertical = alignVertical
}

func (bl *BorderLayout) SetHeightPercents(height int) {
	bl.height = utils.ToString(height) + "%"
}

func (bl *BorderLayout) SetWidthPercents(width int) {
	bl.width = utils.ToString(width) + "%"
}

func (bl *BorderLayout) SetHeightPixels(height int) {
	bl.height = utils.ToString(height) + "px"
}

func (bl *BorderLayout) SetWidthPixels(width int) {
	bl.width = utils.ToString(width) + "px"
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
	table.
		Attr("class", "BorderLayout").
		Attr("cellspacing", "0").
		Attr("cellpadding", "0").
		Attr("style", "width:"+bl.width+";height:"+bl.height+";")

	if bl.top != nil {
		td := NewTD().
			Attr("class", "Top").
			Attr("style", "height:1px;text-align:"+bl.topAlignHorizontal+";vertical-align:"+bl.topAlignVertical).
			AddChild(bl.top)

		if colspan > 1 {
			td.Attr("colspan", utils.ToString(colspan))
		}

		tr := NewTR().AddChild(td)
		table.AddChild(tr)
	}

	tr := NewTR()

	if bl.left != nil {
		td := NewTD().
			Attr("class", "Left").
			Attr("style", "width:1px;text-align:"+bl.leftAlignHorizontal+";vertical-align:"+bl.leftAlignVertical).
			AddChild(bl.left)

		tr.AddChild(td)
	}

	if bl.center != nil {
		td := NewTD().
			Attr("class", "Center").
			Attr("style", "text-align:"+bl.centerAlignHorizontal+";vertical-align:"+bl.centerAlignVertical).
			AddChild(bl.center)

		tr.AddChild(td)
	}

	if bl.right != nil {
		td := NewTD().
			Attr("class", "Right").
			Attr("style", "width:1px;text-align:"+bl.rightAlignHorizontal+";vertical-align:"+bl.rightAlignVertical).
			AddChild(bl.right)

		tr.AddChild(td)
	}

	table.AddChild(tr)

	if bl.bottom != nil {
		td := NewTD().
			Attr("class", "Bottom").
			Attr("style", "height:1px;text-align:"+bl.bottomAlignHorizontal+";vertical-align:"+bl.bottomAlignVertical).
			AddChild(bl.bottom)

		if colspan > 1 {
			td.Attr("colspan", utils.ToString(colspan))
		}

		tr := NewTR().AddChild(td)
		table.AddChild(tr)
	}

	return table.ToHTML()
}
