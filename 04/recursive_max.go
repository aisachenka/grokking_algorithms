package main

import "fmt"

func recursive_max(array []int) int {
	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		} else {
			return array[1]
		}
	}
	max := recursive_max(array[1:])
	if array[0] > max {
		return array[0]
	}
	return max

}

func main() {
	array := []int{11, 12, 3, 4, 15, 6, 7, 8, 9, 10, 7, 8}
	fmt.Println(recursive_max(array))
}
