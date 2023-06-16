package array

import (
	"fmt"
	"strconv"
	"strings"
)

func pivotIndex(nums []int) int {
	l := len(nums)
	sumLeft := make([]int, l)
	sumRight := make([]int, l)

	sum := 0
	for i := 0; i < l; i++ {
		if i != 0 {
			sum += nums[i-1]
		}

		sumLeft[i] = sum
	}

	sum = 0
	for i := l - 1; i >= 0; i-- {
		if i != l-1 {
			sum += nums[i+1]
		}

		sumRight[i] = sum
	}

	fmt.Printf("Left %v, right %v", sumLeft, sumRight)

	for i := 0; i < l; i++ {
		if sumLeft[i] == sumRight[i] {
			return i
		}
	}

	return -1
}

// You are given an integer array nums where the largest integer is unique.
//
// Determine whether the largest element in the array is at least twice as much as every other number in the array.
// If it is, return the index of the largest element, or return -1 otherwise.
func dominantIndex(nums []int) int {
	maxIdx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	m := nums[maxIdx]
	for i := 0; i < len(nums); i++ {
		if i != maxIdx && m < nums[i]*2 {
			return -1
		}
	}

	return maxIdx
}

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry == 1 {
		d := []int{1}
		d = append(d, digits...)
		return d
	}

	return digits
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	var output string
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry = carry + int(a[i]-'0')
			i--
		}

		if j >= 0 {
			carry = carry + int(b[j]-'0')
			j--
		}

		output = strconv.Itoa(carry%2) + output

		carry = carry / 2
	}

	return output
}

// Given two strings needle and haystack,
// return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
func strStr(haystack string, needle string) int {
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		h := haystack[i]
		n := needle[j]

		if h == n {
			j++
		} else {
			// Back to the starting point where
			// the first char in needle is found in
			// haystack
			i -= j
			j = 0
		}

		i++
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}

	}

	return prefix
}
