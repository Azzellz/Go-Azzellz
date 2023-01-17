package main

import "fmt"

func BinarySearch(arr []int, left, right int, dst int) {
	if left > right || dst < arr[0] || dst > arr[len(arr)-1] {
		fmt.Println("No found")
	}
	//mid := (left + right) / 2
	fmt.Println("Finding")
	//标准二分的mid求法
	mid := left + (right-left)*(dst-arr[left])/(arr[right]-arr[left]) //这是插值查找的mid求法,插值查找公式
	//插值查找适用于数据连续均匀的数组
	if dst > arr[mid] { //右递归
		BinarySearch(arr, mid+1, right, dst)
	} else if dst < arr[mid] {
		BinarySearch(arr, left, mid-1, dst)
	} else {
		arrT := make([]int, 0)
		arrT = append(arrT, mid)
		//向左扫描
		temp := mid - 1
		for {
			if temp < 0 || arr[temp] != dst {
				break
			}
			arrT = append(arrT, temp)
			temp--
		}
		//向右扫描
		temp = mid + 1
		for {
			if temp > right || arr[temp] != dst {
				break
			}
			arrT = append(arrT, temp)
			temp++
		}
		fmt.Println(arrT)

	}

}
func main() {
	temp := []int{-4000, -500, -100, 20, 50, 80, 100}
	BinarySearch(temp, 0, len(temp)-1, -4000)

}
