package rail

// Encode encrypts plaintext with given plaintext, k which is a tuple (d, r)
// where d is the depth of the cipher and r is the number of times the algorithm should repeat itself
func Encode(s string, k []int) string {
	if len(k) != 2 {
		return ""
	}
	d := k[0]
	r := k[1]

	var ct string
	for i := 0; i < r; i++ {
		ct = EncodeOnce(s, d)
		s = ct
	}

	return ct
}

// EncodeOnce encrypts plaintext with given s and k and returns ciphertext
// @s: income string
// @n: how many rails we will use
func EncodeOnce(s string, n int) string {
	ss := []rune(s)

	l := len(ss)
	r := make([]rune, l)
	indexes := getIndexes(n-1, l-1)

	for i, idx := range indexes {
		r[i] = ss[idx]
	}

	return string(r)
}

// Decode decrypts ciphertext with given k which is a tuple (d, r)
// where d is the depth of the cipher and r is the number of times the algorithm should repeat itself
func Decode(s string, k []int) string {
	if len(k) != 2 {
		return ""
	}
	d := k[0]
	r := k[1]

	var pt string
	for i := 0; i < r; i++ {
		pt = DecodeOnce(s, d)
		s = pt
	}

	return pt
}

// DecodeOnce decrypts ciphertext with given k and returns plaintext
// @s: ciphered string
// @n: how many rails we will use
func DecodeOnce(s string, n int) string {
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
