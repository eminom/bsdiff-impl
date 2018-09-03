// +build ignore

package main

import (
	"fmt"

	"./sa"
)

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}
	fmt.Println()
}

func main() {
	instr := "aac" //
	// instr := "banana" //
	var a = new(sa.RawSort)
	fmt.Printf("try <%v>\n", instr)
	o := a.SortString(instr)
	printArray(o)
	for _, v := range o {
		fmt.Printf("%v\n", instr[v:])
	}
}
