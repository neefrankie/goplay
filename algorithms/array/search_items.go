package array

import (
	"fmt"
	"sort"
)

func checkIfExist(arr []int) bool {
	l := len(arr)

	for i := 0; i < l-1; i++ {
		fmt.Printf("Outer loop %d: %d\n", i, arr[i])
		for j := i + 1; j < l; j++ {
			fmt.Printf("Inner loop %d: %d\n", j, arr[j])
			if (arr[i] == 2*arr[j]) || (2*arr[i] == arr[j]) {
				return true
			}
		}
	}

	return false
}

type slope int

const (
	slopeUp slope = 1 << iota
	slopeDown
)

func validMountainArray1(arr []int) bool {
	l := len(arr)

	if l < 3 {
		return false
	}

	var dir slope

	for i := 1; i < l; i++ {
		if arr[i-1] == arr[i] {
			return false
		}

		if arr[i-1] < arr[i] {
			if dir&slopeDown > 0 {
				return false
			}

			dir = dir | slopeUp
		}

		if arr[i-1] > arr[i] {
			if dir&slopeUp <= 0 {
				return false
			}

			dir = dir | slopeDown
		}
	}

	if (dir&slopeUp) > 0 && (dir&slopeDown) > 0 {
		return true
	}

	return false
}

// validMountainArray2 by lee215
func validMountainArray2(arr []int) bool {
	l := len(arr)
	if l < 3 {
		return false
	}

	left := 0
	right := l - 1

	for i := 0; i < right; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}

	for left+1 < l && arr[left] < arr[left+1] {
		left++
	}

	for right > 0 && arr[right-1] > arr[right] {
		right--
	}

	if left == right && left < l-1 && right > 0 {
		return true
	}

	return false
}

// Given an integer array nums, return the third distinct maximum number in this array.
// If the third maximum does not exist, return the maximum number.
// Input: nums = [3,2,1]
// Output: 1
// Input: nums = [1,2]
// Output: 2
// Input: nums = [2,2,3,1]
// Output: 1
func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	l := len(nums)
	if l < 3 {
		return nums[0]
	}

	// 3, 2, 2, 1
	count := 1
	for i := 1; i < l; i++ {
		if nums[i-1] != nums[i] {
			count++
		}

		if count == 3 {
			return nums[i]
		}
	}

	return nums[0]
}

// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
// Reference: Section 8.2 Counting Sort, Chapter 8, Introduction to Algorithms
func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	counter := make([]int, l+1)

	for _, v := range nums {
		counter[v] = counter[v] + 1
	}

	fmt.Printf("%v", counter)

	out := make([]int, 0)
	for i := 1; i < len(counter); i++ {
		if counter[i] == 0 {
			out = append(out, i)
		}
	}

	return out
}
