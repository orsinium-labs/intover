package intover

import (
	"fmt"
	"math"
	"testing"
)

type args[I integer] struct {
	a  I
	op rune
	b  I
}

func check[I integer](t *testing.T, args args[I], wantVal I, wantOk bool) {
	name := fmt.Sprintf("%T:%d%s%d", args.a, args.a, string(args.op), args.b)
	t.Run(name, func(t *testing.T) {
		gotVal, gotOk := Do(args.a, args.op, args.b)
		if gotOk != wantOk {
			t.Errorf("ok: got = %v, want %v", gotOk, wantOk)
		}
		if gotVal != wantVal {
			t.Errorf("value: got = %v, want %v", gotVal, wantVal)
		}
	})
}

func TestDo(t *testing.T) {
	ok := true
	over := false

	// int8 + int8
	check(t, args[int8]{4, '+', 5}, 9, ok)
	check(t, args[int8]{126, '+', 5}, -125, over)
	check(t, args[int8]{127, '+', 127}, -2, over)
	check(t, args[int8]{127, '+', -126}, 1, ok)
	check(t, args[int8]{0, '+', 0}, 0, ok)
	check(t, args[int8]{0, '+', 0}, 0, ok)

	// int16 + int16
	check(t, args[int16]{126, '+', 5}, 131, ok)
	check(t, args[int16]{math.MaxInt16, '+', 1}, math.MinInt16, over)
	check(t, args[int16]{math.MaxInt16, '+', math.MaxInt16}, -2, over)
}
