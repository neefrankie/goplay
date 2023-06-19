package array

import (
	"fmt"
	"math"
	"sort"
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

func reverseString(s []byte) []byte {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums)-1; i += 2 {
		sum += nums[i]
	}

	return sum
}

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number.
// Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	for i := 0; i < l; i++ {
		b := target - numbers[i]

		low := i + 1
		high := l - 1
		for low <= high {

			m := (low + high) / 2

			if b == numbers[m] {
				fmt.Printf("Compare %d == %d\n", b, numbers[m])
				return []int{i + 1, m + 1}
			} else if b < numbers[m] {
				high = m - 1
			} else {
				low = m + 1
			}
		}
	}

	return []int{-1, -1}
}

// See https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1153/discuss/2128501/Two-Pointers-or-Visual-Explanation-JAVA
func twoSum2(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for numbers[l]+numbers[r] != target {
		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{l + 1, r + 1}
}

// Given an array of positive integers nums and a positive integer target,
// return the minimal length of a subarray whose sum is greater than or equal to target.
// If there is no such subarray, return 0 instead.
func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sumOfCurrentWindow := 0
	res := math.MaxInt

	for ; right < len(nums); right++ {
		sumOfCurrentWindow += nums[right]

		for sumOfCurrentWindow >= target {
			if res > (right - left + 1) {
				res = right - left + 1
				sumOfCurrentWindow -= nums[left]
				left++
			}
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverseArray(nums[:l-k])
	reverseArray(nums[l-k:])
	reverseArray(nums)
}

func reverseArray(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
