package main

import "fmt"

func BubbleSort(temp []int) {
	count := 0
	for i := 0; i < len(temp)-1; i++ {
		flag := true
		for j := 0; j < len(temp)-1-i; j++ {
			if temp[j] > temp[j+1] {
				temp[j+1], temp[j] = temp[j], temp[j+1]
				flag = false
			}
			count++
		}
		if flag {
			break
		}
	}
	fmt.Println(count)
}

func main() {

	a := []int{3, 1, 2, 100, 4, 5}
	BubbleSort(a)
	fmt.Println(a)
}
