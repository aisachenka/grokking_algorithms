package main

import "fmt"

func qsort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		left := []int{}
		right := []int{}
		pivot := array[0]
		for i := 1; i < len(array); i++ {
			if array[i] < pivot {
				left = append(left, array[i])
			} else {
				right = append(right, array[i])
			}
		}
		left = append(qsort(left), pivot)
		right = qsort(right)
		return append(left, right...)
	}
}
func main() {
	array := []int{10, 5, 12, 7, 3, 8}
	fmt.Println(qsort(array))
}
