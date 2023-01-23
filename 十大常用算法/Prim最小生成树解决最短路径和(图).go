package main

import "fmt"

type primGraph struct {
	weight [][]int  //邻接矩阵表示权重
	data   []string //记录节点数据
	verxs  int      //节点数
}

func graphCreate(verxs int, weight [][]int, data []string) primGraph {
	Mgraph := primGraph{verxs: verxs, weight: weight, data: data}
	return Mgraph
} //根据传入的参数生成图

func (g *primGraph) showGraph() {
	for _, v1 := range g.weight {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}

func (g *primGraph) primGetEdges(first int) {
	visited := make([]int, g.verxs)
	visited[first] = 1 //1表示访问过,0表示未访问过
	h1, h2 := -1, -1
	minWeight := 100
	for k := 0; k < g.verxs-1; k++ { //控制循环轮数,verxs-1次
		for i := 0; i < g.verxs; i++ { //寻找两个节点,它们边的权重最短
			for j := 0; j < g.verxs; j++ { //已访问节点与未访问节点
				if visited[i] == 1 && visited[j] == 0 && g.weight[i][j] < minWeight { //更新最小权重及两个节点的索引
					minWeight = g.weight[i][j]
					h1 = i
					h2 = j
				}
			}
		}
		fmt.Println(g.data[h1], g.data[h2], minWeight)
		//更新visited
		visited[h2] = 1
		minWeight = 100
	}

} //Prim算法获取极小路径子图

func main() {
	notesData := []string{"A", "B", "C", "D", "E", "F", "G"}
	weights := [][]int{
		{100, 5, 7, 100, 100, 100, 2},
		{5, 100, 100, 9, 100, 100, 3},
		{7, 100, 100, 100, 8, 100, 100},
		{100, 9, 100, 100, 100, 4, 100},
		{100, 100, 8, 100, 100, 5, 4},
		{100, 100, 100, 4, 5, 100, 6},
		{2, 3, 100, 100, 4, 6, 100},
	}
	graph := graphCreate(len(notesData), weights, notesData)
	graph.showGraph()
	graph.primGetEdges(0)
}
