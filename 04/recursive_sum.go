package main

import "fmt"

func recursive_sum(array []int) int {
	if len(array) == 0 {
		return 0
	} else {
		return array[0] + recursive_sum(array[1:])
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(recursive_sum(array))
}
