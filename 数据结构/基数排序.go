package main

import (
	"fmt"
	"math/rand"
	"time"
)

type newArr struct {
	arr []int
	len int
}

func (n *newArr) pop() int {
	if n.len == 0 {
		fmt.Printf("None")
		return 0
	}
	temp := n.arr[0]
	n.len--
	if n.len == 0 {
		n.arr = make([]int, 0)
	} else {
		n.arr = append(n.arr[:0], n.arr[1:]...)
	}
	return temp
} //弹出一个数据的方法

func (n *newArr) push(i int) {
	n.arr = append(n.arr, i)
	n.len++
}

func BucketsSort(arr []int) {
	//创建十个桶
	buckets := make([]newArr, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = newArr{make([]int, 0), 0}
	}
	//找到arr中的最高位,用来确定次数
	max := arr[0]
	//fmt.Println(time.Now())
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	//fmt.Println(time.Now())
	//fmt.Println(max)
	//获取最大值的位数
	count := 0
	temp := max
	for temp != 0 {
		temp /= 10
		count++
	}
	//fmt.Println(count)
	for i := 0; i < count; i++ { //总
		fmt.Println("阶段一:")
		fmt.Println(time.Now())
		for j := 0; j < len(arr); j++ {
			if i == 0 { //个位数的特殊情况
				//获取个位数
				buckets[arr[j]%10].push(arr[j])
			} else {
				t := arr[j]
				for s := 0; s < i; s++ {
					t /= 10
				}
				buckets[t%10].push(arr[j])
			}
		} //把arr中的每个元素放入对应的桶中
		fmt.Println(time.Now())

		//从桶中按顺序复原
		fmt.Println("阶段二:")
		//fmt.Println(time.Now())
		tempIndex := 0
		for k := 0; k < 10; k++ {
			if buckets[k].len != 0 {
				fmt.Println(time.Now())
				for buckets[k].len != 0 {
					arr[tempIndex] = buckets[k].pop()
					tempIndex++
				}
				fmt.Println(time.Now())
			}
		}
		//fmt.Println(time.Now())

	}
}

func MakeArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = rand.Intn(80000000)
	}
	return temp
}

func main() {
	fmt.Println(time.Now())

	//temp := []int{400, 103, 200, 10, 100, 250}
	temp := MakeArray(800000)
	BucketsSort(temp)
	fmt.Println(time.Now())
	//fmt.Println(temp)

}
