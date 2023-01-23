package main

import "fmt"

func binarySearch(arr []int, value int) int {
	right := len(arr) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > value {
			right = mid - 1
		} else if arr[mid] < value {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 6, 9}
	fmt.Println(binarySearch(arr, 6))
}
