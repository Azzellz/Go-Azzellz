package main

import "fmt"

type miGong struct {
	chessMap [][]int
}

func (m *miGong) makeMap(length int, width int) {
	temp := make([][]int, length)
	for i := 0; i < length; i++ {
		temp[i] = make([]int, width)
	}
	//让第一行和最后一行变成1
	for i := 0; i < width; i++ {
		temp[0][i] = 1
		temp[width][i] = 1
	}
	//竖墙
	for i := 0; i < length; i++ {
		temp[i][0] = 1
		temp[i][width-1] = 1
	}
	temp[3][1] = 1
	temp[3][2] = 1
	temp[5][3] = 1
	temp[5][4] = 1
	temp[5][5] = 1
	m.chessMap = temp
}

func (m *miGong) showMap() {
	for _, v1 := range m.chessMap {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}

func (m *miGong) setWay(i, j int) bool {
	if m.chessMap[6][5] == 2 { //递归终止条件
		return true
	} else {
		if m.chessMap[i][j] == 0 {
			m.chessMap[i][j] = 2
			if m.setWay(i+1, j) {
				return true
			} else if m.setWay(i, j+1) {
				return true
			} else if m.setWay(i-1, j) {
				return true
			} else if m.setWay(i, j-1) {
				return true
			} else { //回溯
				m.chessMap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {

	m := miGong{}
	m.makeMap(8, 7)
	m.setWay(1, 1)
	m.showMap()

}
