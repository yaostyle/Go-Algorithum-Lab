// ADD ONE INTEGER TO A GIVEN ARRAY
// A golang way to add one integer in a list of array
// Also, in order to solve issue w/ additional leading digit
// such as [9 9 9] to [1 0 0 0], program has to recreate
// a new list.
//
// Example: [1 2 3] -> [1 2 4]
// Example: [9 9 9] -> [1 0 0 0]

package main

import (
	"fmt"
)

func addOne(arr []int) []int {
	lastPos := len(arr) - 1
	carry := 1
	sum := 0

	for i := lastPos; i >= 0; i-- {
		sum = arr[i] + carry

		if sum >= 10 {
			carry = 1
			arr[i] = sum - 10
		} else {
			arr[i] = arr[i] + carry
			carry = 0
		}

		if i == 0 && carry == 1 {
			arr = append([]int{1}, arr...)
		}
	}
	return arr
}

func main() {
	var a = []int{1, 2, 8}
	fmt.Println(a)

	na := addOne(a)
	fmt.Println(na)

	a = []int{9, 9, 9}
	fmt.Println(a)

	na = addOne(a)
	fmt.Println(na)
}
