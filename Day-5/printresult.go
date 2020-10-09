package main

import (
	"learngo/GolangThirty/Day-5/leetcode"
)

func main() {
	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 4, 6, 8, 10}
	leetcode.FindMedianSortedArrays(nums1, nums2)

	s := "abcbcda"
	leetcode.LongestPalindrome(s)
}
