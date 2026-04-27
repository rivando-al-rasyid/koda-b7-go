package main

import "fmt"

func slices() {
	numbers := []int{50, 75, 66, 20, 32, 90}
	end := make([]int, len(numbers)-3)
	copy(end, numbers[len(numbers)-3:])
	front := numbers[:len(numbers)-3]
	result := append(front, 88)
	result = append(result, end...)
	fmt.Println(result)

}
