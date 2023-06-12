package dma

import (
	"fmt"
	"strings"
)

// This file answers questions of Exercise 3.1 of
// Discrete Mathematics and Its Applications.

// SumList finds the sum of all the integers in a list.
// Question 3.
func SumList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var sum int

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}

// LargestDifference produces as output the largest difference obtained
// by subtracting an integer in the list from the one following it.
// Question 4.
func LargestDifference(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	var max = arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff > max {
			max = diff
		}
	}

	return max
}

// FindDuplicates takes as input a list of n integers
// in non-decreasing order and produces the list of all
// values that occur more than one.
// A list of integers is non-decreasing if each integer
// in the list is at least as large as the previous integer
// in the list.
// Question 5.
func FindDuplicates(arr []int) []int {
	var items = make([]int, 0)

	var i = 0
	var size = len(arr)
	for i < size {
		// Starting from the next one, stops at the first one
		// that is not equal to arr[i]:
		// [2, 2, 3, 4, 6, 6, 6]
		var j = i + 1
		// Count duplicates.
		// This programmed could be easily modified to tell how many
		// duplicates there are for each item.
		var count = 0
		for ; j < size && arr[i] == arr[j]; j++ {
			count++
		}
		// At least one duplicate is found.
		if count > 0 {
			items = append(items, arr[i])
		}
		// Next round loop starts from j.
		i = j
	}

	return items
}

// CountNegativeInts implements question 6:
// Takes as input a list of n integers and finds the
// number of negative integers in the list.
func CountNegativeInts(ints []int) int {
	var count int

	for _, n := range ints {
		if n < 0 {
			count++
		}
	}

	return count
}

// LastEvenInt implements question 7:
// Takes as input a list of n integers
// and find the location of the last even integer in the list
// or returns -1 if there are no event integers in the list.
func LastEvenInt(arr []int) int {
	var pos int = -1
	for i, n := range arr {
		if n%2 == 0 {
			pos = i
		}
	}

	return pos
}

// IsPalindrome implements question 9 to test if a string is a palindrome.
// A palindrome is a string that reads the same forward and backward. For example:
// A man, a plan, a canal, Panama!
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	var left = 0
	var right = len(s) - 1

	for left < right {
		leftCh := s[left]
		rightCh := s[right]
		if !isLetter(leftCh) {
			left++
			continue
		}
		if !isLetter(rightCh) {
			right--
			continue
		}
		if leftCh != rightCh {
			return false
		}
		left++
		right--
	}

	return true
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

// XPower implements question 10 to compute x^n,
// where x is a real number and n is an integer.
func XPower(x float64, n int) float64 {
	if n >= 0 {
		return xPower(x, n)
	}

	return 1 / xPower(x, -n)
}

func xPower(x float64, n int) float64 {
	if n < 0 {
		panic("n must be non-negative")
	}
	var result float64 = 1

	for i := 1; i <= n; i++ {
		result = result * x
	}

	return result
}

// LongestWord implements question 22:
// find the longest word in an English sentence.
func LongestWord(sentence string) string {
	// Ensure we know the end of sentence is reached.
	if !strings.HasSuffix(sentence, "\n") {
		sentence += "\n"
	}

	var inWorld = false

	var word string
	var start = 0
	for i, ch := range sentence {
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			// The first time read a space after a char.
			if inWorld && i-start > len(word) {
				word = sentence[start:i]
			}
			inWorld = false
		} else if !inWorld {
			// The first time read a char after discarding space
			inWorld = true
			start = i
		}
	}

	return word
}

func TernarySearch(x int, arr []int) int {
	var low = 0
	var high = len(arr) - 1

	for low < high {
		diff := high - low + 1
		if diff < 3 {
			if x == arr[low] {
				return low
			} else if x == arr[high] {
				return high
			} else {
				return -1
			}
		}

		m := (diff) / 3
		fmt.Printf("Divider %d\n", m)

		if x == arr[m] {
			return m
		} else if x < arr[m] {
			high = m
		} else if x == arr[2*m] {
			return 2 * m
		} else if x < arr[2*m] {
			low = m + 1
			high = 2 * m
		} else {
			low = 2*m + 1
		}
	}

	return -1
}

func ClosestPair(arr []float64) (float64, float64) {
	var a, b float64

	var distance = arr[1] - arr[0]
	if distance < 0 {
		distance = -distance
	}

	for outer := 1; outer < len(arr); outer++ {
		for inner := 0; inner < outer; inner++ {
			d := arr[outer] - arr[inner]
			if d < 0 {
				d = -d
			}

			if d < distance {
				distance = d
				a = arr[inner]
				b = arr[outer]
			}
		}
	}

	return a, b
}

func GreaterThanPreviousTermsSum(arr []int) []int {
	var result = make([]int, 0)
	var sum = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > sum {
			result = append(result, arr[i])
		}
		sum += arr[i]
	}

	return result
}
