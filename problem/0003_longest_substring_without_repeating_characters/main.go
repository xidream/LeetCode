package main

/*
LengthOfLongestSubstring
  - Runtime: 0 ms | Beats 100.00%
  - Memory: 5.01 MB | Beats 24.54%
*/
func LengthOfLongestSubstring(s string) int {
	var (
		start, size, rv, lenS = 0, 0, 0, len(s)
		prevIdxes             = make(map[rune]int)
	)
	for idx, char := range s {
		if prevIdx, found := prevIdxes[char]; found && prevIdx >= start {
			size = idx - start
			if size > rv {
				rv = size
			}
			start = prevIdx + 1
		}
		prevIdxes[char] = idx
	}
	size = lenS - start
	if size > rv {
		return size
	}
	return rv
}
