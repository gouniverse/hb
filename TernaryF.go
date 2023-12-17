package hb

// TernaryF is a 1 line if/else statement whose options are functions
func TernaryF(condition bool, trueFunc func() *Tag, falseFunc func() *Tag) *Tag {
	if condition {
		return trueFunc()
	}

	return falseFunc()
}
