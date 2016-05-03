// Package kmp provides ...
package kmp

func preKMP(s string) []int {
	preArr := make([]int, len(s))
	j, i := 0, 1
	for i < len(s) {
		if s[i] == s[j] {
			preArr[i] = j + 1
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = preArr[j]
		}
	}
	return preArr
}

// KMP algorithm
func KMP(pat, content string) []int {
	pats := preKMP(pat)
	var res []int
	i, j, pl := 0, 0, len(pat)

	for i < len(content) {
		if content[i] == pat[j] {
			j++
			i++
		} else if j == 0 {
			i++
		} else {
			j = pats[j-1]
		}

		if pl == j {
			res = append(res, i-j)
			j = 0
		}
	}
	return res
}

// Brute forece matching
func BruteF(pat, content string) []int {
	var res []int
	i, pl := 0, len(pat)

	for i < len(content)-pl {
		if content[i:i+pl] == pat {
			res = append(res, i)
		}
		i++
	}
	return res
}
