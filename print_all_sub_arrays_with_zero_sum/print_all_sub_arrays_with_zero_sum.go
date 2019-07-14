// PRINT ALL SUB-ARRAYS WITH 0 SUM
// Go program to find zero sum in a given array
//
// Example: [4,2,-3,-1,0,4]
// 0 Sum Pairs:
// [3,4,-7]
// [4,-7,3]
// [-7,3,1,3]
// [3,1,-4]
// [3,1,3,1,-4,-2,-2]
// [3,4,-7,3,1,3,1,-4,-2,-2]

package main

import (
	"fmt"
)

func findZeroSum(arr []int) {
	arrayCount := len(arr)
	sum := 0

	for i := 0; i < arrayCount; i++ {

		sum = 0
		for j := 0; j < arrayCount; j++ {
			sum += arr[j]
			if sum == 0 {
				fmt.Printf("Found pair at: [%v %v]\n", arr[i], arr[j])
			}

		}
	}
}

func main() {
	array := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	fmt.Printf("Given arry: %v\n", array)

	findZeroSum(array)
}
