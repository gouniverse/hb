package hb

// Ternary is a 1 line if/else statement.
func Ternary(condition bool, trueTag *Tag, falseTag *Tag) *Tag {
	if condition {
		return trueTag
	}

	return falseTag
}
