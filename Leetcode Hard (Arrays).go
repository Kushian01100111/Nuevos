package kata

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeArray := append(nums1, nums2...)

	sort.Slice(mergeArray, func(i, j int) bool {
		return mergeArray[i] < mergeArray[j]
	})

	median := float64(len(mergeArray)) / 2.0
	expr := math.Mod(median, 1.0)
	if expr == 0.0 {
		temp := (float64(mergeArray[int(median-1)]) + float64(mergeArray[int(median)])) / 2.0
		return temp
	} else {
		temp := float64(mergeArray[int(median)])
		return temp
	}
}
