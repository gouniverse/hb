package hb

func If(condition bool, trueTag *Tag) *Tag {
	if condition {
		return trueTag
	}
	return nil
}

func IfF(condition bool, trueFunc func() *Tag) *Tag {
	if condition {
		return trueFunc()
	}

	return nil
}

func IfElse(condition bool, trueTag *Tag, falseTag *Tag) *Tag {
	if condition {
		return trueTag
	}
	return falseTag
}

func IfFElseF(condition bool, trueFunc func() *Tag, falseFunc func() *Tag) *Tag {
	if condition {
		return trueFunc()
	}

	return falseFunc()
}
