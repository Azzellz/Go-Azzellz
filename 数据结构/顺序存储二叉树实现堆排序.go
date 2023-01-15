package main

import "fmt"

func heapify(arr []int, i int, len int) {

	left := 2*i + 1
	right := 2*i + 2
	//假设最大值就是子节点
	max := i

	if left < len && arr[left] > arr[max] {
		max = left
	} //与左子树比较,更新最大值

	if right < len && arr[right] > arr[max] {
		max = right
	} //与右子树比较,更新最大值

	//判断最大值的索引是否发生改变,若索引发生改变则交换先前的值
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, max, len)
	}
} //i表示非叶子树的节点

func HeapSort(arr []int) {
	//先调整初态堆

	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	//交换并且再次调整
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)
	}
}

func main() {
	arr := []int{10, 9, 8, 5, 100}
	fmt.Println("排序前：", arr)
	HeapSort(arr)
	fmt.Println("排序后：", arr)
}
