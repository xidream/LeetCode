package main

/*
TwoSum
  - Runtime: 0 ms | Beats 100.00%
  - Memory: 5.68 MB | Beats 8.82%
*/
func TwoSum(nums []int, target int) []int {
	want := make(map[int]int)
	for idx, num := range nums {
		if wantIdx, found := want[num]; found {
			return []int{wantIdx, idx}
		}
		want[target-num] = idx
	}
	return nil
}
