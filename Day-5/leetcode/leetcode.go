package leetcode

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2[:]...)
	var m float64
	for i := 0; i < len(s)-1; i++ {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	fmt.Println(s)
	if a := len(s) - 1; a%2 == 0 {
		m = float64(s[a/2])
	} else {
		m = float64(s[a/2] + s[a/2+1])
		m = m / 2
	}
	fmt.Println(m)
	return m
}

func LongestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	fmt.Println(ans)
	return ans
}
