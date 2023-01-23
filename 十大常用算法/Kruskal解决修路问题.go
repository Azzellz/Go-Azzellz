package main

import "fmt"

const (
	INF = 100000 //表示未连通的占位数
)

type edge struct {
	start  string
	end    string
	weight int
} //边包含两端节点和边的权值

type kruskalGraph struct {
	vertexs  []string //节点组
	matrix   [][]int  //邻接矩阵
	edgesNum int      //图的边数
	edges    []edge   //边的切片
	ends     []int    //记录各顶点终点的数组,随算法动态添加元素
}

func (g *kruskalGraph) getEnd(i int) int {
	for g.ends[i] != 0 {
		i = g.ends[i]
	}
	return i
} //返回i这个下标对于节点的终点的下标,用来判断是否会产生回路

func createKruskalGraph(verts []string, matrix [][]int) *kruskalGraph {
	temp := &kruskalGraph{vertexs: verts, matrix: matrix}
	//temp.edgesNum = temp.countEdges()
	temp.putEdgesInArray()
	temp.edgesNum = len(temp.edges)
	temp.ends = make([]int, temp.edgesNum)
	return temp

}

func (g *kruskalGraph) kruskal() []edge {
	//创建一个结果edge切片用来储存最小生成树
	resultEdges := make([]edge, 0)
	//遍历图的中各个边(排序过后的
	for i := 0; i < g.edgesNum; i++ {
		//分别取出当前遍历到的边的start和end对应的索引
		p1 := g.backIndex(g.edges[i].start)
		p2 := g.backIndex(g.edges[i].end)
		//分别得到p1和p2的终点
		m := g.getEnd(p1)
		n := g.getEnd(p2)
		//判断是否构成回路
		if m != n { //如果两个点终点不同就加入到结果切片,并且更新Ends
			g.ends[m] = n
			resultEdges = append(resultEdges, g.edges[i])
		}
	}
	return resultEdges
}

func (g *kruskalGraph) countEdges() (count int) {

	for _, v1 := range g.matrix {
		for _, v2 := range v1 {
			if v2 != INF && v2 != 0 {
				count++
			}
		}
	}
	return
}

func (g *kruskalGraph) showGraph() {
	for _, v1 := range g.matrix {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

}

func (g *kruskalGraph) backIndex(dest string) int {
	for i, v := range g.vertexs {
		if v == dest {
			return i
		}
	}
	return -1
} //返回目标节点对应的下标

func (g *kruskalGraph) putEdgesInArray() {
	for i := 0; i < len(g.matrix); i++ {
		for j := i + 1; j < len(g.matrix[i]); j++ {
			if g.matrix[i][j] != INF {
				g.edges = append(g.edges, edge{weight: g.matrix[i][j], start: g.vertexs[i], end: g.vertexs[j]})
			}
		}
	}
	g.sortEdges()
} //遍历邻接矩阵将边放入graph结构体中的edges切片

func (g *kruskalGraph) sortEdges() {
	for i := 0; i < len(g.edges)-1; i++ {
		for j := 0; j < len(g.edges)-1-i; j++ {
			if g.edges[j].weight > g.edges[j+1].weight {
				g.edges[j], g.edges[j+1] = g.edges[j+1], g.edges[j]
			}
		}
	}
}

func main() {
	notesData := []string{"A", "B", "C", "D", "E", "F", "G"}
	matrix := [][]int{
		{0, 12, INF, INF, INF, 16, 14},
		{12, 0, 10, INF, INF, 7, INF},
		{INF, 10, 0, 3, 5, 6, INF},
		{INF, INF, 3, 0, 4, INF, INF},
		{INF, INF, 5, 4, 0, 2, 8},
		{16, 7, 6, INF, 2, 0, 9},
		{14, INF, INF, INF, 8, 9, 0},
	}
	g := createKruskalGraph(notesData, matrix)
	g.showGraph()
	fmt.Println(g.kruskal())
}
