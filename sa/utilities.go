package sa

import (
	"fmt"
)

func printArray(a []int) {
	out := ""
	for _, v := range a {
		out += fmt.Sprintf("%v ", v)
	}
	fmt.Printf("%v\n", out)
}
