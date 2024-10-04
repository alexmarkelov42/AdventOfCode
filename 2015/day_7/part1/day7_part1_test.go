package day7_part1

import (
	"testing"
)

func TestParseNum(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		number uint16
	}{
		{"test1", "123 -> x", 123},
	}
	for _, val := range cases {
		result := parseNum(val.input)
		if result != val.number {
			t.Errorf("%s: got %d while goal is %d\n", val.name, result, val.number)
		}
	}
}

func TestEvalExpr(t *testing.T) {
	wires["a"] = &wire{"123", 0, false}
	wires["b"] = &wire{"NOT a", 0, false}
	wires["c"] = &wire{"456", 0, false}
	wires["d"] = &wire{"a AND c", 0, false}
	wires["e"] = &wire{"a LSHIFT 3", 0, false}
	wires["f"] = &wire{"c RSHIFT 3", 0, false}
	wires["g"] = &wire{"c OR a", 0, false}
	cases := []struct {
		name     string
		wireName string
		result   uint16
	}{
		{"test 1", "a", 123},
		{"test 2", "b", ^uint16(123)},
		{"test 3", "d", uint16(123) & uint16(456)},
		{"test 4", "e", uint16(123) << 3},
		{"test 5", "f", uint16(456) >> 3},
		{"test 6", "g", 123 | 456},
	}
	for _, val := range cases {
		result := evalExpr(val.wireName)
		if result != val.result {
			t.Errorf("%s: got %d while goal is %d\n", val.name, result, val.result)
		}
	}
}
