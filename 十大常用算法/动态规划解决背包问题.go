package main

import "fmt"

func dynamic(weight int, arrPrice []int, arrWeight []int) [][]int {
	//1.先创建一个表记录最优子结构,并且先初始化表的第一行第一列为0
	list := make([][]int, len(arrPrice)+1)
	for i := range list {
		for j := 0; j < weight+1; j++ {
			list[i] = append(list[i], 0)
		}
	}
	notebook := make([][]int, len(arrPrice)+1)
	for i := range notebook {
		for j := 0; j < weight+1; j++ {
			notebook[i] = append(notebook[i], 0)
		}
	}
	//for i := 0; i < len(list)+1; i++ {
	//	list[i][0] = 0
	//}
	//for i := 0; i < weight+1; i++ {
	//	list[0][i] = 0
	//}
	//规定三种物品各自都只有一件,1:吉他(1KG): 1500元 ; 2:电脑(3KG) 2000元 ; 3:音响(4KG) 3000元
	//2.开始遍历填表
	//max := func(a, b int) int {
	//	if a > b {
	//		return a
	//	} else {
	//		return b
	//	}
	//}

	for i := 1; i < len(list); i++ {
		//i表示当前物品;j表示背包的容重情况;i-1表示上一个商品
		for j := 1; j < len(list[0]); j++ { //要装入的商品重量大于当前背包容重,那么参照上一个商品的容量
			if arrWeight[i-1] > j {
				list[i][j] = list[i-1][j]

			} else { //否则,比较最大值
				if list[i-1][j] < arrPrice[i-1]+list[i-1][j-arrWeight[i-1]] {
					notebook[i][j] = 1
					list[i][j] = arrPrice[i-1] + list[i-1][j-arrWeight[i-1]]
				} else {
					list[i][j] = list[i-1][j]
				}
				//list[i][j] = max(list[i-1][j], arrPrice[i-1]+list[i-1][j-arrWeight[i-1]])

			}
		}
	}

	//3.根据notebook和list打印出需要放入背包的物品
	//直接从最后一行考虑,最后的一行的数据是最大的
	i := len(list) - 1
	j := len(list[0]) - 1
	for i > 0 && j > 0 {
		if notebook[i][j] == 1 {
			fmt.Println(i, "号物品")
			j -= arrWeight[i-1]
		}
		i--
	}

	return list
}

func main() {
	list := dynamic(4, []int{2000, 2000, 2000}, []int{2, 2, 2})
	for i, v1 := range list {
		if i == 0 {
			continue
		}
		fmt.Println()
		for j, v2 := range v1 {
			if j == 0 {
				continue
			}
			fmt.Print(" ", v2, " ")
		}
	}
}
