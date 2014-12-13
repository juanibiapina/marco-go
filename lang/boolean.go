package lang

type boolean bool

func MakeBoolean(b bool) boolean {
	if b {
		return true
	} else {
		return false
	}
}

func (b boolean) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (b boolean) IsTrue() bool {
	return bool(b)
}

func (b boolean) Equal(o Expr) bool {
	switch other := o.(type) {
	case boolean:
		return b == other
	default:
		return false
	}
}

func (b boolean) Not() boolean {
	return !b
}
