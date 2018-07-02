package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return float64((nums1[0] + nums2[0]) / 2)
}

func findKSortedArrays(nums1 []int, nums2 []int, k int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > k/2 && len2 > k/2 {
		findKSortedArraysDouble(nums1, nums2, k)
	} else if len1 > k {
		findKSortedArraysSingle(nums1, nums2, k)
	} else {
		findKSortedArraysSingle(nums2, nums1, k)
	}
}

func findKSortedArraysDouble(nums1 []int, nums2 []int, k int) float64 {
	k1 := nums1[k/2-1]
	k2 := nums2[k/2-1]

	if k1 == k2 {
		return float64(k1)
	} else if k1 > k2 {
		findKSortedArrays(nums1[k/2:], nums2, k/2+1)
	} else {
		findKSortedArrays(nums1, nums2[k/2:], k/2+1)
	}
}

func findKSortedArraysSingle(nums1 []int, nums2 []int, k int) float64 {
	l2 := len(nums2)
	k1 = nums1[k-l2]
	k2 = nums2[l2]
	if k1 == k2 {
		return k1
	} else if k1 > k2 {

	} else {
	}
	return float64(0)
}
