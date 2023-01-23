package main

import "fmt"

const Unconnection = 114514

type floydGraph struct {
	vertex []string
	matrix [][]int
	pre    [][]int //前驱节点的在vertex对应的下标
}

func createFloydGraph(vertex []string, matrix [][]int) floydGraph {
	temp := floydGraph{
		vertex: vertex,
		matrix: matrix,
	}
	preArr := make([][]int, len(matrix))
	for i := 0; i < len(preArr); i++ {
		preArr[i] = make([]int, len(matrix))
	}
	for i, v := range preArr {
		for j := range v {
			preArr[i][j] = i
		}
	}

	temp.pre = preArr
	return temp
}

func (g *floydGraph) backPath(start, end int) []string {
	pathArr := make([]string, 0)
	pathArr = append(pathArr, g.vertex[end])
	for start != end {
		preNode := g.pre[start][end]
		pathArr = append(pathArr, g.vertex[preNode])
		end = preNode
	}
	for i := 0; i < len(pathArr)/2; i++ {
		pathArr[i], pathArr[len(pathArr)-1-i] = pathArr[len(pathArr)-1-i], pathArr[i]
	}
	return pathArr
} //回溯路径

func (g *floydGraph) showGraph() {
	fmt.Println("距离图:")
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix); j++ {
			fmt.Print(g.matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("前驱节点图:")
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix); j++ {
			fmt.Print(g.vertex[g.pre[i][j]], " ")
		}
		fmt.Println()
	}
}

func (g *floydGraph) floyd() {
	//这三层遍历都是遍历同一个vertex,即遍历三层{"A", "B", "C", "D", "E", "F", "G"}
	for k := 0; k < len(g.matrix); k++ { //遍历中间节点
		tempLen := 0
		for i := 0; i < len(g.matrix); i++ { //遍历出发节点
			for j := 0; j < len(g.matrix); j++ { //遍历终点
				tempLen = g.matrix[i][k] + g.matrix[k][j]
				if g.matrix[i][j] > tempLen {
					g.matrix[i][j] = tempLen
					g.pre[i][j] = g.pre[k][j]
				}
			}
		}
	}

	for k := 0; k < len(g.vertex); k++ {
		fmt.Printf("%v点到各个顶点的最短路径为: \n", g.vertex[k])
		for i := 0; i < len(g.vertex); i++ {
			fmt.Print(g.backPath(k, i), " ")
		}
		fmt.Println()
	}

}

func main() {
	notesData := []string{"A", "B", "C", "D", "E", "F", "G"}
	matrix := [][]int{
		{0, 5, 7, Unconnection, Unconnection, Unconnection, 2},
		{5, 0, Unconnection, 9, Unconnection, Unconnection, 3},
		{7, Unconnection, 0, Unconnection, 8, Unconnection, Unconnection},
		{Unconnection, 9, Unconnection, 0, Unconnection, 4, Unconnection},
		{Unconnection, Unconnection, 8, Unconnection, 0, 5, 4},
		{Unconnection, Unconnection, Unconnection, 4, 5, 0, 6},
		{2, 3, Unconnection, Unconnection, 4, 6, 0},
	}
	g := createFloydGraph(notesData, matrix)
	g.floyd()
	g.showGraph()
}
