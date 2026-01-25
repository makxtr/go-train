package main

/**
Have the function SimpleMode(arr) take the array of numbers stored in arr and return the number that appears most frequently (the mode).
For example: if arr contains [10, 4, 5, 2, 4] the output should be 4.
If there is more than one mode return the one that appeared in the array first (ie. [5, 10, 10, 6, 5] should return 5 because it appeared first).
If there is no mode return -1. The array will not be empty.
Examples
Input: []int {5,5,2,2,1}
Output: 5
Input: []int {3,4,1,6,10}
Output: -1
*/

func SimpleMode(arr []int) int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	if maxCount == 1 {
		return -1
	}

	for _, num := range arr {
		if counts[num] == maxCount {
			return num
		}
	}

	return -1
}
