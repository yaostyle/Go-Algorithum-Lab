// REARRANGE ODD AND EVEN VALUES IN ALTERNATE FASHION IN
// ASCENDING ORDER
// Go program to rearrange odd and even values
// Example given array: [9,8,13,2,19,14]
// Output: 2, 9, 8, 13, 14, 19

package main

import (
	"fmt"
	"sort"
)

func sortArrayInAscOrder(arr []int) []int {
	// Vars
	arrayCount := len(arr) // Length of array (6)
	evenList := []int{}    // Even list
	oddList := []int{}     // Odd list

	//Rearrange new array in ASC order
	sort.Sort(sort.IntSlice(arr))

	// Iterate the given array and sort into even/odd list
	for i := 0; i < arrayCount; i++ {
		// If it can be divided by 2, it's an even number
		// Set to evenList
		if arr[i]%2 == 0 {
			evenList = append(evenList, arr[i])
		} else {
			oddList = append(oddList, arr[i])
		}
	}

	// If the given array starts with odd, flag it
	startWithEven := false
	if arr[0]%2 == 0 {
		startWithEven = true
	}
	// startWithEven: false

	// Rearrange the array
	a := 0 // Even list index
	b := 0 // Odd list index
	for index, _ := range arr {
		// If first element is even
		if startWithEven {
			arr[index] = evenList[a]
			a++
			startWithEven = false
		} else {
			//First element is odd
			arr[index] = oddList[b]
			b++
			startWithEven = true
		}
	}

	return arr
}

func main() {
	// Given array
	array := []int{9, 8, 13, 2, 19, 14}
	// Run the logic
	newArray := sortArrayInAscOrder(array)
	// Print out result
	fmt.Printf("New ASC array: %v", newArray)
}
