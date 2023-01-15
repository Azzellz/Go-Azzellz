package main

import "fmt"

func merge(arr []int, left int, right int, middle int, temp []int) {
	i := left       //初始化左域的索引
	j := middle + 1 //初始化右域的索引
	t := 0          //初始化temp中转数组的索引

	//先将temp填充
	for i <= middle && j <= right {
		if arr[i] >= arr[j] {
			temp[t] = arr[j]
			t++
			j++
		} else {
			temp[t] = arr[i]
			t++
			i++
		}
	}
	//处理有剩余的数据
	for i <= middle {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = arr[j]
		j++
		t++
	}
	//拷贝temp给arr
	t = 0
	templeft := left
	for templeft <= right {
		arr[templeft] = temp[t]
		t++
		templeft++
	}
} //并-

func divAndMerge(arr []int, left int, right int, temp []int) {
	if left < right {
		middle := (left + right) / 2
		divAndMerge(arr, left, middle, temp)    //左
		divAndMerge(arr, middle+1, right, temp) //右
		merge(arr, left, right, middle, temp)

	}
}

func main() {
	temp := []int{400, 103, 200, 10, 100}
	new := make([]int, len(temp))
	divAndMerge(temp, 0, len(temp)-1, new)
	fmt.Println(temp)
}
