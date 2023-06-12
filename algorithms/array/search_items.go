package array

import "fmt"

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
	first := -1
	second := -1
	third := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] >= first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] >= second {
			third = second
			second = nums[i]
		} else if nums[i] >= third {
			third = nums[i]
		}
	}

	fmt.Printf("%d, %d, %d", first, second, third)

	if third > -1 {
		return third
	} else {
		return first
	}
}
