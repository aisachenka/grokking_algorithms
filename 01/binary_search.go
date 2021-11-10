package main

import "fmt"

func binary_search(array []int, element int) {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == element {
			fmt.Println(mid)
		}
		if array[mid] > element {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20}
	binary_search(list, 5)
}
