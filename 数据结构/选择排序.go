package main

import "fmt"

func SelectSort(temp []int) {
	for i := 0; i < len(temp)-1; i++ {
		//先找最小值
		min := temp[i]
		minindex := i
		for j := i; j < len(temp); j++ {
			if temp[j] < min {
				minindex = j
				min = temp[j]
			} //更新最小值
		}
		if minindex != i {
			temp[i], temp[minindex] = temp[minindex], temp[i]
		}

	}
}

func main() {

	temp := []int{400, 103, 200, 10, 100}
	SelectSort(temp)
	fmt.Println(temp)
}
