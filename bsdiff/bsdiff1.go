package bsdiff

import (
	"fmt"
)

var (
	_ = fmt.Sprintf
)

// These two interfaces are not compatible at all.
func SortOutString1(a string) []int {
	return SortOut1([]byte(a))
}

func fetchBig(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func SortOut1(ib []byte) []int {
	// fmt.Printf("test<%v>\n", string(ib))
	n := len(ib)

	// reference value array.
	x := make([]int, n+1)
	for i, v := range ib {
		x[i] = int(v)
	}
	x[n] = -1

	// suffix-id(partial) to be sorted.
	y := make([]int, n+1)

	// buckets
	bucket := make([]int, fetchBig(n, 256))

	// Fall for the first round
	for i := 0; i < n; i++ {
		bucket[x[i]]++
	}
	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	// suffix array: [0..n-1] to be sorted.
	sa := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := x[i]
		bucket[v]--
		sa[bucket[v]] = i
	}

	for h := 1; h*2 < n; h *= 2 {
		p := 0
		for i := n - h; i < n; i++ {
			y[p] = i
			p++
		}
		for _, v := range sa {
			if v >= h {
				y[p] = v - h
				p++
			}
		}

		for i := range bucket {
			bucket[i] = 0
		}
		for i := 0; i < n; i++ {
			bucket[x[y[i]]]++
		}
		for i := 1; i < len(bucket); i++ {
			bucket[i] += bucket[i-1]
		}
		// get the next SA
		for i := n - 1; i >= 0; i-- {
			bucket[x[y[i]]]--
			sa[bucket[x[y[i]]]] = y[i]
		}
		// use y to evaluate the reference-value array(x)
		p, y[sa[0]] = 1, 0
		for i := 1; i < n; i++ {
			lhs := sa[i]
			rhs := sa[i-1]
			// try this: abcbcb111(will crash this line)
			// if x[lhs] != x[rhs] || x[lhs+h] != x[rhs+h]

			if x[lhs] != x[rhs] || !((lhs+h < n && rhs+h < n && x[lhs+h] == x[rhs+h]) || (lhs+h >= n && rhs+h >= n)) {
				y[sa[i]] = p
				p++
			} else {
				y[sa[i]] = p - 1
			}
		}
		y, x = x, y
	}
	return sa
}
