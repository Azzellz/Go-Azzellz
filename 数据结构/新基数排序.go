package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newBucketsSort(arr []int) {
	buckets := make([][]int, 0)
	for i := 0; i < 10; i++ {
		temp := make([]int, len(arr))
		buckets = append(buckets, temp)
	}
	//找到arr中的最高位,用来确定次数
	max := arr[0]
	//fmt.Println(time.Now())
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	//获取最大值的位数
	count := 0
	temp := max
	for temp != 0 {
		temp /= 10
		count++
	}
	//再创建一个大小为10的数组来记录每个桶中有效数据的个数
	bucketsCount := make([]int, 10)
	for i := 0; i < count; i++ { //总

		for j := 0; j < len(arr); j++ {
			if i == 0 { //个位数的特殊情况
				//获取个位数
				t := arr[j] % 10
				buckets[t][bucketsCount[t]] = arr[j]
				bucketsCount[t]++
			} else {
				t := arr[j]
				for s := 0; s < i; s++ {
					t /= 10
				}
				c := t % 10
				buckets[c][bucketsCount[c]] = arr[j]
				bucketsCount[c]++
			}
		} //把arr中的每个元素放入对应的桶中

		//从桶中按顺序复原

		tempIndex := 0
		for k := 0; k < 10; k++ {
			tempI := 0
			if bucketsCount[k] != 0 {
				for bucketsCount[k] != 0 {
					arr[tempIndex] = buckets[k][tempI]
					bucketsCount[k]--
					tempI++
					tempIndex++
				}

			}
		}

	}
}
func makeArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = rand.Intn(800000000)
	}
	return temp
}
func main() {
	//temp := []int{400, 103, 200, 10, 100}
	fmt.Println(time.Now())
	temp := makeArray(8000000)
	newBucketsSort(temp)
	fmt.Println(time.Now())
	//fmt.Println(temp)
}
