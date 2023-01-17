package main

import "fmt"

func InsertSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		insertVal := arr[i+1]
		insertIndex := i

		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		arr[insertIndex+1] = insertVal

	}
}

func main() {
	temp := []int{400, 103, 200, 10, 100}
	InsertSort(temp)
	fmt.Println(temp)
}
