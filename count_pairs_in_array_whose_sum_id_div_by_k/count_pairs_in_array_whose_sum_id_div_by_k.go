// COUNT PAIRS IN ARRAY WHOSE SUM IS DIVISIBLE BY K
// Use Golang method to find the sum of pairs that can be divided by K
// Example Array: [2,2,1,7,5,3]
// Given K=4
// Pair count(s): 5

package main

import (
	"fmt"
)

func findParisDivideBy(arr []int, k int) int {
	// Vars
	paircount := 0         // Total pair counts
	arrayCount := len(arr) // Length of given array

	// Iterate the array (2 loops: x, y)
	for x := 0; x < arrayCount; x++ {
		for y := x + 1; y < arrayCount; y++ {
			pairSum := 0 // Sum of the pair
			// Add up the pair
			pairSum = arr[x] + arr[y]

			// Divide it by K
			if pairSum%k == 0 {
				// If dividisible, adds 1 to count
				paircount++
			}
		}
	}

	return paircount
}

func main() {
	// Given array
	array := []int{2, 2, 1, 7, 5, 3}
	k := 4

	// findParis
	paris := findParisDivideBy(array, k)

	// Print out result
	fmt.Printf("Pair counts: %v", paris)
}
