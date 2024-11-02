package main

/*
FindMedianSortedArrays
  - Runtime: 0 ms | Beats 100.00%
  - Memory: 4.24 MB | Beats 40.40%
*/
func LongestPalindrome(s string) string {
	var (
		rv          = GetLongestPalindromeFromMiddle(s)
		lenS, lenRV = len(s), len(rv)
	)

	for i := 1; ; i++ {
		if lenRV >= lenS-i {
			return rv
		}
		v := GetLongestPalindromeFromMiddle(s[i:])
		if len(v) > lenRV {
			rv = v
			lenRV = len(rv)
			if lenRV == lenS-i {
				return rv
			}
		}
		v = GetLongestPalindromeFromMiddle(s[:lenS-i])
		if len(v) > lenRV {
			rv = v
			lenRV = len(rv)
		}
	}
}

func GetLongestPalindromeFromMiddle(s string) string {
	var (
		rv          = s[0:1]
		lenS        = len(s)
		left, right int
	)

	if lenS%2 == 0 {
		right = lenS / 2
		left = right - 1
	} else {
		left = lenS/2 - 1
		right = left + 2
	}

	for left >= 0 {
		if s[left] != s[right] {
			break
		}
		rv = s[left : right+1]
		left--
		right++
	}

	return rv
}
