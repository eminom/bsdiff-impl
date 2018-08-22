package bsdiff

import (
	"fmt"
	"testing"
)

func printArray(a []int, t *testing.T) {
	out := ""
	for _, v := range a {
		out += fmt.Sprintf("%v ", v)
	}
	t.Logf("%v", out)
}

func testBsdiff(str string, t *testing.T) {
	o := SortOutString(str)
	printArray(o, t)
	for _, v := range o {
		t.Logf("%v", str[v:])
	}
}

func TestBsdiff(t *testing.T) {
	testBsdiff("banana", t)
	testBsdiff("b\x00anana", t)
	testBsdiff("b\x00an\x00ana", t)
	testBsdiff("b\x00an\x00ana\x00", t)
}
