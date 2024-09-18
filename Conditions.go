package hb

func If(condition bool, trueTag TagInterface) TagInterface {
	if condition {
		return trueTag
	}
	return nil
}

func IfF(condition bool, trueFunc func() TagInterface) TagInterface {
	if condition {
		return trueFunc()
	}

	return nil
}

func IfElse(condition bool, trueTag TagInterface, falseTag TagInterface) TagInterface {
	if condition {
		return trueTag
	}
	return falseTag
}

func IfFElseF(condition bool, trueFunc func() TagInterface, falseFunc func() TagInterface) TagInterface {
	if condition {
		return trueFunc()
	}

	return falseFunc()
}
