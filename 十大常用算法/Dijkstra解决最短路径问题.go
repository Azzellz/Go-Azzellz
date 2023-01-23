package main

import "fmt"

const (
	Unconnection1 = 65535
)

type dijkstraGraph struct {
	vertexs []string //节点组
	matrix  [][]int  //邻接矩阵
}

type visitVertex struct {
	alreadyArr  []int //记录已经访问过的节点的下标
	preVisitArr []int //记录各个节点的前驱节点
	dis         []int //记录各个顶点的最小距离,动态更新
}

func createVisitVertex(vertexNum int, currentVertex int) *visitVertex {
	vv := &visitVertex{alreadyArr: make([]int, vertexNum), preVisitArr: make([]int, vertexNum)}
	tmp := make([]int, vertexNum)
	for i := 0; i < vertexNum; i++ {
		if i != currentVertex {
			tmp[i] = Unconnection1
		}
	}
	tmp[currentVertex] = 0
	vv.alreadyArr[currentVertex] = 1
	vv.dis = tmp
	return vv
} //根据传进来的初始节点下标和节点数创建并且返回一个visitVertex对象

func (v *visitVertex) isVisited(i int) bool {
	return v.alreadyArr[i] == 1
} //判断节点是否被访问过

func (v *visitVertex) updateDis(index int, len int) {
	v.dis[index] = len
} //更新最短值

func (v *visitVertex) updatePre(pre, index int) {
	v.preVisitArr[pre] = index
} //更新前驱节点

func (v *visitVertex) getDis(index int) int {
	return v.dis[index]
} //返回距离出发顶点的最小值

func (v *visitVertex) updateArr() int {
	min := Unconnection1
	index := 0
	for i := 0; i < len(v.alreadyArr); i++ {
		if v.alreadyArr[i] == 0 && v.dis[i] < min {
			min = v.dis[i]
			index = i
		}
	}
	v.alreadyArr[index] = 1
	return index
}

func (g *dijkstraGraph) showGraph() {
	for _, v1 := range g.matrix {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}

func createDijkstraGraph(vertexs []string, matrix [][]int) *dijkstraGraph {
	return &dijkstraGraph{vertexs: vertexs, matrix: matrix}
}

func (g *dijkstraGraph) dijkstra(first int) {
	vV := createVisitVertex(len(g.vertexs), first)
	g.update(first, vV)
	for i := 1; i < len(g.vertexs); i++ {
		g.update(vV.updateArr(), vV)
	}
	fmt.Println("good")
} //从first开始执行dijkstra

func (g *dijkstraGraph) update(index int, vv *visitVertex) {
	length := 0
	for i := 0; i < len(g.matrix[index]); i++ {
		length = vv.getDis(index) + g.matrix[index][i]
		if !vv.isVisited(i) && length < vv.getDis(i) {
			vv.updateDis(i, length)
			vv.updatePre(i, index)
		}
	}
}

func main() {
	notesData := []string{"A", "B", "C", "D", "E", "F", "G"}
	matrix := [][]int{
		{Unconnection1, 5, 7, Unconnection1, Unconnection1, Unconnection1, 2},
		{5, Unconnection1, Unconnection1, 9, Unconnection1, Unconnection1, 3},
		{7, Unconnection1, Unconnection1, Unconnection1, 8, Unconnection1, Unconnection1},
		{Unconnection1, 9, Unconnection1, Unconnection1, Unconnection1, 4, Unconnection1},
		{Unconnection1, Unconnection1, 8, Unconnection1, Unconnection1, 5, 4},
		{Unconnection1, Unconnection1, Unconnection1, 4, 5, Unconnection1, 6},
		{2, 3, Unconnection1, Unconnection1, 4, 6, Unconnection1},
	}
	g := createDijkstraGraph(notesData, matrix)
	g.showGraph()
	g.dijkstra(6)
}
