package main

import "fmt"

func ShellSort(arr []int) {

	//交换法(慢
	//for gap := len(arr); gap > 0; gap /= 2 {
	//	for i := gap; i < len(arr); i++ {
	//		for j := i - gap; j >= 0; j -= gap {
	//			if arr[j] > arr[j+gap] {
	//				arr[j], arr[j+gap] = arr[j+gap], arr[j]
	//			}
	//		}
	//	}
	//}
	//移位法(快
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {

			insertIndex := i
			insertVal := arr[insertIndex]
			if arr[insertIndex] < arr[insertIndex-gap] {
				for insertIndex-gap >= 0 && arr[insertIndex-gap] > insertVal {
					arr[insertIndex] = arr[insertIndex-gap]
					insertIndex -= gap
				}

				arr[insertIndex] = insertVal
			}

		}

	}
}

func main() {
	temp := []int{400, 103, 200, 10, 100}
	ShellSort(temp)
	fmt.Println(temp)
}
