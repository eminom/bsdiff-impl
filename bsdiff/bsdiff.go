package bsdiff

func SortOutString(a string) []int {
	return SortOut([]byte(a))
}

func SortOut(raw []byte) []int {
	length := len(raw)
	var bucket [256]int
	for _, v := range raw {
		bucket[v]++
	}
	k := 1
	for i := 0; i < 256; i++ {
		if bucket[i] != 0 {
			bucket[i] = k
			k++
		}
	}

	q := make([]int, 1+length)
	for i, v := range raw {
		q[i] = bucket[v]
	}
	// printArray(q)
	nq := make([][2]int, 1+length)
	for h := 1; h < length; h *= 2 {
		for i := 0; i < 1+length; i++ {
			nq[i][0] = q[i]
			if i+h >= length {
				nq[i][1] = 0
			} else {
				nq[i][1] = q[i+h]
			}
		}
		radix := make([][]int, 1+length)
		for i := 0; i < 1+length; i++ {
			radix[nq[i][1]] = append(radix[nq[i][1]], i)
		}
		radix2 := make([][]int, 1+length)
		for _, vr := range radix {
			for _, v := range vr {
				radix2[nq[v][0]] = append(radix2[nq[v][0]], v)
			}
		}
		var prev int
		nk := -1
		for _, vr := range radix2 {
			prev = -1
			for _, v := range vr {
				vnow := nq[v][0]*(1+length) + nq[v][1]
				if prev != vnow {
					nk++
				}
				q[v] = nk
				prev = vnow
			}
		}
		// printArray(q)
	}

	p := make([]int, 1+length)
	for i := 0; i < 1+length; i++ {
		p[q[i]] = i
	}
	return p[1:]
}
