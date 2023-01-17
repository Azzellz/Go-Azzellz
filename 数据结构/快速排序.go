package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low < high {
		pos := GetPosition(arr, low, high)
		QuickSort(arr, low, pos-1)
		QuickSort(arr, pos+1, high)
	}
}

func GetPosition(arr []int, low, high int) int {
	pivot := arr[low] //每次都以左边界为中轴
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		//结束完这上面的循环就找到位置了
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	//这时候low和high相遇
	arr[low] = pivot
	return low
} //用来返回分界
func main() {
	temp := []int{400, 103, 200, 10, 100}
	QuickSort(temp, 0, len(temp)-1)
	fmt.Println(temp)
}
