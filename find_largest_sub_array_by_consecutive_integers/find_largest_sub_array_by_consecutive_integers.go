// !!! (NOT COMPLETE) FIND LARGEST SUB-ARRAY FORMED BY CONSECUTIVE INTEGERS
// Go program to find largest sub-arrays
//
// Example: [2,0,2,1,4,3,1,0]
// Largest sub-array: [1,5]

package main

//"fmt"

// Function to find subarray is a consecutive integers or not
func isConsecutive(arr []int, i int, j int, min int, max int) bool {
	if (max - min) != (j - i) {
		return false
	}

	visited := []bool{}

	for k := i; k <= j; k++ {
		if visited[arr[k]-min] {
			return false
		}

		visited[arr[k]-min] = true
	}

	return true
}

// Function to find a smallest number in an array
// and return its index
func compareNum(arr []int, method string) int {
	// assuming the first element is the smallest number
	smallestNum := arr[0]
	var index int

	// Loop through the array

	// Base on the method, do the comparison
	switch method {
	case "min":
		for i, v := range arr {
			if v < smallestNum {
				index = i
			}
		}
	case "max":
		for i, v := range arr {
			if v > smallestNum {
				index = i
			}
		}
	default:
		index = -1
	}

	return index
}

// Function to find the largest sub-array
func findLargestSubArray(arr []int) {
	// Vars
	length := 1
	//start, end := 0, 0
	arrayCount := len(arr)
	min_val, max_val := 0, 0

	// Iterating the array
	for i := 0; i < arrayCount; i++ {
		min_val = arr[i]
		max_val = arr[i]

		for j := i + 1; j < arrayCount; j++ {
			min_val = compareNum(arr, "min")
			max_val = compareNum(arr, "max")

			// Check if subarray is formed by consecutive integers
			if isConsecutive(arr, i, j, min_val, max_val) {
				if length < max_val-min_val+1 {
					length = max_val - min_val + 1
				}
			}
		}

	}

}

// Main entry point
func main() {
	// Declare given array
	array := []int{2, 0, 2, 1, 4, 3, 1, 0}

	// Find the largest sub-array
	findLargestSubArray(array)
}
