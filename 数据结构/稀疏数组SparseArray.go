package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//先创建棋盘
	chess := [11][11]int{}
	//为棋盘赋上有效值
	chess[1][2] = 1
	chess[2][3] = 2
	//遍历棋盘找出有效值的位数
	sum := 0
	for _, v := range chess {
		for _, v1 := range v {
			if v1 != 0 {
				sum++
			}
		}
	}
	//创建稀疏数组
	//sparseArr := [sum][3]int{}
	sparseArr := make([][3]int, sum+1)
	sparseArr[0][0] = 11
	sparseArr[0][1] = 11
	sparseArr[0][2] = sum

	//将有效数据录入稀疏数组
	count := 0
	for i, v := range chess {
		for j, v1 := range v {
			if v1 != 0 {
				count++
				sparseArr[count][0] = i
				sparseArr[count][1] = j
				sparseArr[count][2] = v1
			}
		}
	}
	//将稀疏数组存入文件
	src, err1 := os.Create("./chess.txt")
	if err1 != nil {
		fmt.Printf("%v", err1)
		return
	}
	defer src.Close()

	for i := 0; i < len(sparseArr); i++ {
		for j := 0; j < len(sparseArr[i]); j++ {
			src.WriteString(strconv.Itoa(sparseArr[i][j]) + "\t")
		}
		src.WriteString("\n")
	}

	f, err2 := os.Open("./chess.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer f.Close()

	//buf := make([]byte, 1024) //4k大小的临时缓冲区
	//
	//for {
	//	n,err2:=src.Read(buf)
	//	if err2 == io.EOF {
	//		break
	//	}
	//}

	//至此稀疏数组的构建完成,现在执行复原操作
	//fmt.Println(sparseArr)
	//如下是用make内置函数以非字面量创建二维数组
	newChess := make([][]int, sparseArr[0][0])
	for i := range newChess {
		newChess[i] = make([]int, sparseArr[0][1])
	}
	//fmt.Println(newChess)
	//开始根据稀疏数组的数据复原
	for i := 1; i < len(sparseArr); i++ {
		newChess[sparseArr[i][0]][sparseArr[i][1]] = sparseArr[i][2]
	}
	for _, v := range newChess {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}
}
