package hb

// This file contains idiomatic functions for creating HTML tags
// These functions all start with the New prefix
// Shortcuts are defined in shortcuts.go

// NewBR represents a BR tag
// Shortcut method exists: BR()
func NewBR() *Tag {
	return &Tag{TagName: "br"}
}

// NewButton represents a BUTTON tag
// Shortcut method exists: Button()
func NewButton() *Tag {
	return &Tag{TagName: "button"}
}

// NewCaption is a shortcut to create a new CAPTION tag
// Shortcut method exists: Caption()
func NewCaption() *Tag {
	return &Tag{TagName: "caption"}
}

// NewHeading1 represents a H1 tag
func NewHeading1() *Tag {
	return &Tag{
		TagName: "h1",
	}
}

// NewHeading2 represents a H1 tag
func NewHeading2() *Tag {
	return &Tag{
		TagName: "h2",
	}
}

// NewHeading3 represents a H1 tag
func NewHeading3() *Tag {
	return &Tag{
		TagName: "h3",
	}
}

// NewHeading4 represents a H1 tag
func NewHeading4() *Tag {
	return &Tag{
		TagName: "h4",
	}
}

// NewHeading5 represents a H1 tag
func NewHeading5() *Tag {
	return &Tag{
		TagName: "h5",
	}
}

// NewHeading6 represents a H1 tag
func NewHeading6() *Tag {
	return &Tag{
		TagName: "h6",
	}
}

// NewH1 represents a H1 tag
func NewH1() *Tag {
	return NewHeading1()
}

// NewH2 represents a H2 tag
func NewH2() *Tag {
	return NewHeading2()
}

// NewH3 represents a H3 tag
func NewH3() *Tag {
	return NewHeading3()
}

// NewH4 represents a H4 tag
func NewH4() *Tag {
	return NewHeading4()
}

// NewH5 represents a H5 tag
func NewH5() *Tag {
	return NewHeading5()
}

// NewH6 represents a H6 tag
func NewH6() *Tag {
	return NewHeading6()
}
