package main

import "fmt"

func findSmallest(array []int) int {
	smallest := array[0]
	smallest_index := 0
	for i := 1; i < len(array); i++ {
		if array[i] < smallest {
			smallest = array[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selection_sort(array []int) []int {
	sorted_array := make([]int, len(array))
	for i := 0; i < len(sorted_array); i++ {
		smallest := findSmallest(array)
		sorted_array[i] = array[smallest]
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return sorted_array
}

func main() {
	list := []int{12, 562, 23445, 432, 234, 4, 2434, 23, 7, 8, 24, 6, 2, 1}
	fmt.Println(selection_sort(list))
}
