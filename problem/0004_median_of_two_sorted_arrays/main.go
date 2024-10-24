package main

/*
FindMedianSortedArrays
  - Runtime: 0 ms | Beats 100.00%
  - Memory: 6.36 MB | Beats 17.18%
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		size1            = len(nums1)
		size2            = len(nums2)
		size             = size1 + size2
		stop             = size / 2
		idx1, idx2, left = 0, 0, 0
	)

	for n := 0; n < stop; n++ {
		if idx1 == size1 {
			left = nums2[idx2]
			idx2++
			continue
		}
		if idx2 == size2 {
			left = nums1[idx1]
			idx1++
			continue
		}
		if nums1[idx1] < nums2[idx2] {
			left = nums1[idx1]
			idx1++
			continue
		}
		left = nums2[idx2]
		idx2++
	}

	if size%2 == 0 {
		if idx1 == size1 {
			return float64(left+nums2[idx2]) / 2
		}
		if idx2 == size2 {
			return float64(left+nums1[idx1]) / 2
		}
		if nums1[idx1] < nums2[idx2] {
			return float64(left+nums1[idx1]) / 2
		}
		return float64(left+nums2[idx2]) / 2
	}

	if idx1 == size1 {
		return float64(nums2[idx2])
	}
	if idx2 == size2 {
		return float64(nums1[idx1])
	}
	if nums1[idx1] < nums2[idx2] {
		return float64(nums1[idx1])
	}
	return float64(nums2[idx2])
}
