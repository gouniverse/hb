package hb

func If(condition bool, trueTag *Tag) *Tag {
	if condition {
		return trueTag
	}
	return nil
}

func IfElse(condition bool, trueTag *Tag, falseTag *Tag) *Tag {
	if condition {
		return trueTag
	}
	return falseTag
}
