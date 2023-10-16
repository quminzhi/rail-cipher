package rail

// Encode -
//
//	  s - income string
//	  n - how many rails we will use
//	OUT: ciphered string
func Encode(s string, n int) string {
	ss := []rune(s)

	l := len(ss)
	r := make([]rune, l)
	indexes := getIndexes(n-1, l-1)

	for i, idx := range indexes {
		r[i] = ss[idx]
	}

	return string(r)
}

// Decode -
//
//	  s - ciphered string
//	  n - how many rails we will use
//	OUT: decoded string
func Decode(s string, n int) string {
	ss := []rune(s)

	l := len(ss)
	r := make([]rune, l)
	indexes := getIndexes(n-1, l-1)

	for i, idx := range indexes {
		r[idx] = ss[i]
	}

	return string(r)
}

func getIndexes(n, m int) (r []int) {
	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			for j := 0; j <= m; j = j + 2*n {
				r = append(r, j)
			}
		case n:
			for j := n; j <= m; j = j + 2*n {
				r = append(r, j)
			}
		default:
			for j := i; j <= m; {
				r = append(r, j)
				next := j + 2*n - 2*i
				if next <= m {
					r = append(r, next)
				}
				j += 2 * n
			}
		}
	}
	return r
}
