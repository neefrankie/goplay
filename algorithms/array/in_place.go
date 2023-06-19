package array

func replaceElements(arr []int) []int {
	max := -1
	l := len(arr)

	for i := l - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max
		if tmp > max {
			max = tmp
		}
	}

	return arr
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	// Use the two pointer technique to remove the duplicates in-place.
	// The first elemnt shouldn't be touched; it's already in its correct place.
	writerPointer := 1

	for readPointer := 1; readPointer < l; readPointer++ {
		// If the current element we're reading is different from the
		// previous element...
		if nums[readPointer] != nums[readPointer-1] {
			// Copy it to the next position at the front, tracked by writePointer.
			nums[writerPointer] = nums[readPointer]
			// Increment writePointer, because the next element should be written one space over.
			writerPointer++
		}
	}

	// The corrent length value.
	return writerPointer
}

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
// Input: nums = [0,1,0,3,12]
// Output:       [1,3,12,0,0]
func moveZeroes(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}

	lastNoneZeroIdx := 0
	for cur := 0; cur < l; cur++ {
		if nums[cur] != 0 {
			nums[lastNoneZeroIdx] = nums[cur]
			lastNoneZeroIdx++
		}
	}

	for i := lastNoneZeroIdx; i < l; i++ {
		nums[i] = 0
	}
}

// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.
// Input: nums = [3,1,2,4]
// Output:       [2,4,3,1]
func sortArrayByParity(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}

	lastOddIdx := 0
	for cur := 0; cur < l; cur++ {
		if nums[cur]%2 == 0 {
			nums[lastOddIdx], nums[cur] = nums[cur], nums[lastOddIdx]
			lastOddIdx++
		}
	}

	return nums
}

// Given an integer array nums and an integer val,
// remove all occurrences of val in nums in-place.
// The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
// Input:     nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// See https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1157/
func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	k := 0
	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
// Input: nums = [-4,-1,0,3,10]
// Output:       [0,1,9,16,100]
func sortedSquares(nums []int) []int {
	left := 0
	l := len(nums)
	right := l - 1
	out := make([]int, l)

	for i := l - 1; i >= 0; i-- {
		lv := nums[left] * nums[left]
		rv := nums[right] * nums[right]
		if lv > rv {
			out[i] = lv
			left++
		} else {
			out[i] = rv
			right--
		}
	}

	return out
}
