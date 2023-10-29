package intover

type integer interface {
	uintptr |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Must[I integer](val I, ok bool) I {
	if !ok {
		panic("integer overflow")
	}
	return val
}

func Do[I integer](a I, op rune, b I) (I, bool) {
	switch op {
	case '+':
		return Add(a, b)
	case '-':
		return Sub(a, b)
	case '*':
		return Mul(a, b)
	case '/':
		return Div(a, b)
	default:
		panic("invalid operation")
	}
}

func Add[I integer](a, b I) (I, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Sub[I integer](a, b I) (I, bool) {
	c := a - b
	if (c < a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Mul[I integer](a, b I) (I, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}

func Div[I integer](a, b I) (I, bool) {
	if b == 0 {
		return 0, false
	}
	c := a / b
	status := (c < 0) == ((a < 0) != (b < 0))
	return c, status
}
