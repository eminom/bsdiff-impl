package sa

type StdDoubleAlgo struct{}

// These two interfaces are not compatible at all.
func (d *StdDoubleAlgo) SortString(a string) []int {
	return d.Sort([]byte(a))
}

func (*StdDoubleAlgo) Sort(ib []byte) []int {
	// fmt.Printf("test<%v>\n", string(ib))
	n := len(ib)

	// buckets
	var bucket [256]int // 256 different values

	// Fall for the first round
	for i := 0; i < n; i++ {
		bucket[ib[i]]++
	}

	// reference value, starting from 0
	p := 0
	for i, v := range bucket {
		if v != 0 {
			bucket[i] = p
			p++
		}
	}

	// reference value array.(range 0~n-1(maximum))
	x := make([]int, n)
	for i, v := range ib {
		x[i] = bucket[v]
	}

	// Must do it again
	for i := 0; i < len(bucket); i++ {
		bucket[i] = 0
	}
	for i := 0; i < n; i++ {
		bucket[ib[i]]++
	}
	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	// suffix array: [0..n-1] to be sorted.
	sa := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := ib[i]
		bucket[v]--
		sa[bucket[v]] = i
	}

	// suffix-id(partial) to be sorted.
	y := make([]int, n)

	// up to n different values, starting with the minimum of 0
	wv := make([]int, n)

	for h := 1; p < n; h *= 2 {
		p = 0
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

		for i := range wv {
			wv[i] = 0
		}
		for i := 0; i < n; i++ {
			wv[x[y[i]]]++
		}
		for i := 1; i < len(wv); i++ {
			wv[i] += wv[i-1]
		}
		// get the next SA
		for i := n - 1; i >= 0; i-- {
			wv[x[y[i]]]--
			sa[wv[x[y[i]]]] = y[i]
		}
		// use y to evaluate the reference-value array(x)
		p, y[sa[0]] = 1, 0
		for i := 1; i < n; i++ {
			lhs := sa[i]
			rhs := sa[i-1]
			// try this: abcbcb111(will crash this line)
			// if x[lhs] != x[rhs] || x[lhs+h] != x[rhs+h]

			// lhs and rhs can not be >= n for the same time
			// For one: lhs+h >= n && rhs+h >= n will never happen
			if x[lhs] != x[rhs] || (lhs+h < n) != (rhs+h < n) || x[lhs+h] != x[rhs+h] {
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
