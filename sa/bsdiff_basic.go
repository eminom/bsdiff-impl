/*
RawSort:  The formatting of double algorithm
*/
package sa

type RawSort struct{}

func (r *RawSort) SortString(a string) []int {
	return r.Sort([]byte(a))
}

func (*RawSort) Sort(raw []byte) []int {
	n := len(raw)
	var bucket [256]int
	for _, v := range raw {
		if 0 == bucket[v] {
			bucket[v] = 1
		}
	}

	// reference value ranged: 1 ~ number of items.
	// reference value always retain their relative order.
	// so the specific ord order is not critical
	// referenve value 0 is given to the zero-width string(which is unique and there is only one)

	// reference value will range [0,n] (n+1 different values when SA is fully formatted)
	var k int = 1
	for i := 0; i < 256; i++ {
		if 0 != bucket[i] {
			bucket[i] = k
			k++
		}
	}

	// array of reference value: q
	q := make([]int, 1+n)
	for i, v := range raw {
		q[i] = bucket[v]
	}
	// printArray(q)

	nq := make([][2]int, 1+n)
	for h := 1; h < n; h *= 2 {
		// generate the 'next' q(reference value array)
		// with two parts:
		// the higher h and the lower h:
		for i := 0; i < 1+n; i++ {
			nq[i][0] = q[i]
			if i+h > n {
				nq[i][1] = 0
			} else {
				nq[i][1] = q[i+h]
			}
		}
		// Sort the lower part.
		// As a result, the lower part is always ready
		// This part can be omitted gracefully.
		radix := make([][]int, 1+n)
		for i := 0; i < 1+n; i++ {
			radix[nq[i][1]] = append(radix[nq[i][1]], i)
		}

		// Sort the higher part.
		radix2 := make([][]int, 1+n)
		for _, vr := range radix {
			for _, v := range vr {
				radix2[nq[v][0]] = append(radix2[nq[v][0]], v)
			}
		}

		var prev int
		// the previous index (suffix-id)
		// update the refrence value for array q
		var nk int = 0
		for _, vr := range radix2 {
			prev = -1
			for _, v := range vr {
				if prev < 0 || nq[prev][1] != nq[v][1] {
					nk++
				}
				q[v] = nk - 1
				prev = v
			}
		}
		// printArray(q)
	}

	p := make([]int, 1+n)
	for i := 0; i < 1+n; i++ {
		p[q[i]] = i
	}
	return p[1:]
}
