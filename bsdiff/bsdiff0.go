package bsdiff

func SortOutString0(a string) []int {
	return SortOut0([]byte(a))
}

func SortOut0(raw []byte) []int {
	n := len(raw)
	x := make([]int, n)
	var bucket [256]int
	for _, v := range raw {
		bucket[v]++
	}
	k := 0
	for i := 0; i < 256; i++ {
		if bucket[i] != 0 {
			bucket[i] = k
			k++
		}
	}
	for i, v := range raw {
		x[i] = bucket[v]
	}
	y := make([]int, n)
	//h*2 < n &&
	for h := 1; k < n; h *= 2 {
		radix := make([][]int, n+1)
		// Pay attention: the shortest gets the frontest:
		// not (not i<= 0~n-1)
		for i := 0; i < n; i++ {
			if i+h >= n {
				radix[0] = append(radix[0], i)
			} else {
				radix[x[i+h]+1] = append(radix[x[i+h]+1], i)
			}
		}
		rad2 := make([][]int, n)
		for _, vr := range radix {
			for _, s := range vr {
				rad2[x[s]] = append(rad2[x[s]], s)
			}
		}
		prev := -1
		k = 0
		for _, vr := range rad2 {
			for _, s := range vr {
				if prev < 0 || x[prev] != x[s] || (prev+h < n) != (s+h < n) || x[prev+h] != x[s+h] {
					k++
				}
				y[s] = k - 1
				prev = s
			}
		}
		x, y = y, x
	}
	rv := make([]int, n)
	for i := 0; i < n; i++ {
		rv[x[i]] = i
	}
	return rv
}
