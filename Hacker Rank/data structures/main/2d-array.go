package main

import (
	"fmt"
)

func hourglassSum(arr [6][6]int32) int32 {
	// Write your code here
	var tmpSum int32 = 0
	var maxSum int32 = -99999999

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {

			if j+2 < 6 && i+2 < 6 {

				tmpSum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

				if tmpSum > maxSum {
					maxSum = tmpSum
				}
			}

		}
	}

	return maxSum

}

func main() {

	arr := [6][6]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	fmt.Println("Initial Array: ", arr)

	res := hourglassSum(arr)

	fmt.Println("Sum: ", res)

}
