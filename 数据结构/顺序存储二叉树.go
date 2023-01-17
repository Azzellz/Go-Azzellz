package main

import "fmt"

type arrBinaryTree struct {
	arr []int
}

func (arr *arrBinaryTree) arrPreOrder(index int) {
	if len(arr.arr) == 0 || arr.arr == nil {
		fmt.Println("无效数组")
		return
	}

	fmt.Println(arr.arr[index])
	if index*2+1 <= len(arr.arr)-1 {
		arr.arrPreOrder(index*2 + 1)
	} //向左递归
	if index*2+2 <= len(arr.arr)-1 {
		arr.arrPreOrder(index*2 + 2)
	} //右边递归

}

func (arr *arrBinaryTree) arrInfixOrder(index int) {
	if len(arr.arr) == 0 || arr.arr == nil {
		fmt.Println("无效数组")
		return
	}

	if index*2+1 <= len(arr.arr)-1 {
		arr.arrInfixOrder(index*2 + 1)
	} //向左递归
	fmt.Println(arr.arr[index])
	if index*2+2 <= len(arr.arr)-1 {
		arr.arrInfixOrder(index*2 + 2)
	} //右边递归
}

func (arr *arrBinaryTree) arrPostOrder(index int) {
	if len(arr.arr) == 0 || arr.arr == nil {
		fmt.Println("无效数组")
		return
	}

	if index*2+1 <= len(arr.arr)-1 {
		arr.arrPostOrder(index*2 + 1)
	} //向左递归

	if index*2+2 <= len(arr.arr)-1 {
		arr.arrPostOrder(index*2 + 2)
	} //右边递归
	fmt.Println(arr.arr[index])
}

func main() {
	arr := arrBinaryTree{arr: []int{1, 2, 3, 4, 5, 6, 7}}
	arr.arrPostOrder(0)
}
