package main

import (
	"fmt"
)

func reverseArray(a []int32) []int32 {
	// Write your code here
	j := len(a)

	b := []int32{}

	for i := 0; i < len(a); i++ {
		b = append(b, a[j-1])
		j = j - 1
	}

	return b

}

func main() {

	arr := []int32{1, 5, 8, 6, 2, 7}

	fmt.Println("Initial Array: ", arr)

	res := reverseArray(arr)

	fmt.Println("Reversed Array: ", res)

}
