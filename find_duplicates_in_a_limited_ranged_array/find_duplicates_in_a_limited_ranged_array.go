// FIND A DUPLICATE ELEMENT IN A LIMITED RANGE ARRAY
// Go program to find duplicate elements
//
// Example: [1,2,3,4,4]
// Duplicates: 4

package main

import (
	"fmt"
)

// Func to find duplicates
func findDuplicate(arr []int) {
	// Vars
	arrayCount := len(arr) // input array counts
	visited := []int{}     // visited array

	// First loop to iterate the array
	for i := 0; i < arrayCount; i++ {

		// Check visited var with current element
		for _, v := range visited {
			if v == arr[i] {
				fmt.Printf("Found duplicates: %v\n", v)
			}
		}

		// Flag into visited var for each element
		visited = append(visited, arr[i])
	}
}

// Entry point
func main() {
	//Declare limited range
	array := []int{1, 2, 3, 4, 4}
	findDuplicate(array)
}
