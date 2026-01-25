package main

/*

Have the function DistinctList(arr) take the array of numbers stored in arr
and determine the total number of duplicate entries.
For example if the input is [1, 2, 2, 2, 3] then your program should output 2
because there are two duplicates of one of the elements.

	Examples
	Input: []int {0,-2,-2,5,5,5}
	Output: 3
	Input: []int {100,2,101,4}
	Output: 0
*/

func DistinctList(arr []int) int {
	seen := make(map[int]struct{})
	duplicates := 0
	for _, num := range arr {
		if _, ok := seen[num]; ok {
			duplicates++
		} else {
			seen[num] = struct{}{}
		}
	}
	return duplicates
}
