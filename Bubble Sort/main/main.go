package main

import "fmt"

func main() {

	s := []int{1, 4, 3, 5, 8, 2, 7, 9, 6}

	swapped := false

	for i := 0; i < len(s); i++ {
		swapped = false

		for j := 0; j < len(s)-1-i; j++ {
			tmp := 0

			if s[j] > s[j+1] {
				tmp = s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
				swapped = true
				fmt.Println(" => swapped ", s[j], " ", s[j+1])
			}

		}

		if !swapped {
			break
		}

		fmt.Println("Iteration:, ", i)
		fmt.Println(s)

	}
}
