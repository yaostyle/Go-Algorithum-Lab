// FIND PAIR WITH GIVEN SUM IN THE ARRAY
// Go program to find the pair(s) with giving sum .
//
// Example: [8,7,2,5,3,1]
// Given Sum: 10
// Pairs: [8,2], [7,3]

package main

import (
	"fmt"
)

func findPair(arr []int, sum int) {
	arrayCount := len(arr)
	for i := 0; i < arrayCount; i++ {
		for j := i + 1; j < arrayCount; j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("Found pair at [%v %v]\n", arr[i], arr[j])
			}
		}
	}
}

func main() {
	givenSum := 10
	array := []int{8, 7, 2, 5, 3, 1}

	fmt.Printf("Given sum: %v\n", givenSum)
	findPair(array, givenSum)
}
