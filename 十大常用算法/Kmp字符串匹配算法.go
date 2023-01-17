package main

import "fmt"

func baoli(a, b string) int {
	//判断b是否为a的字串
	i, j := 0, 0
	//i指向a,j指向b;分别作为字符指针
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else { //回溯
			i = i - (j - 1)
			j = 0
		}
	}

	if j == len(b) {
		return i - j
	} //说明匹配成功

	return -1

} //暴力匹配字符串

func kmpNext(str string) []int {
	arr := make([]int, len(str))
	arr[0] = 0 //表的第一个元素一定是0
	for i, j := 1, 0; i < len(str); i++ {

		for j > 0 && str[i] != str[j] {
			j = arr[j-1]
		} //kmp的核心

		if str[i] == str[j] {
			j++
		}
		arr[i] = j
	}
	return arr
} //先得到str的部分匹配表

func kmpSearch(str1, str2 string, next []int) int {

	for i, j := 0, 0; i < len(str1); i++ {

		for j > 0 && str1[i] != str2[j] {
			j = next[j-1]
		}

		if str1[i] == str2[j] {
			j++
		}
		if j == len(str2) {
			return i - j + 1
		}
	}
	return -1
}

func main() {
	a := "WDNMDDdd  Cao"
	b := "Cao"
	//fmt.Println(baoli(a, b))
	fmt.Println(kmpSearch(a, b, kmpNext(b)))
}
