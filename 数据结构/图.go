package main

import (
	"fmt"
)

type graph struct {
	arr         []string //记录图的成员
	mappingList [][]int  //映射关系表,这个表的值对应着关系权
	memberNum   int      //图的成员数量
	isVisited   []bool   //遍历的时候用来判断是否已经被访问
	isBonded    []bool   //是否有建立连接
	bfsQueue    bfsQueue
}

type bfsQueue struct {
	arr []int
	len int
}

func (q *bfsQueue) addQueue(n ...int) {
	for i := 0; i < len(n); i++ {
		q.arr = append(q.arr, n[i])
		q.len++
	}
}

func (q *bfsQueue) removeQueue() int {
	temp := q.arr[0]
	q.arr = q.arr[1:]
	q.len--
	return temp
}

func (g *graph) addMember(member string) {
	g.memberNum++
	g.arr = append(g.arr, member)
	g.mappingList = append(g.mappingList, make([]int, g.memberNum))
	g.isVisited = append(g.isVisited, false)
	g.isBonded = append(g.isBonded, false)
	for i := range g.mappingList {
		for len(g.mappingList[i]) < g.memberNum {
			g.mappingList[i] = append(g.mappingList[i], 0)
		}
	}
} //添加图的成员

func (g *graph) makeBond(i int, j int, weight int) {
	if i > len(g.arr)-1 || j > len(g.arr)-1 {
		fmt.Println("越界访问")
		return
	}
	g.mappingList[i][j] = weight
	g.mappingList[j][i] = weight
	g.isBonded[i] = true
	g.isBonded[j] = true
} //将图的成员间建立起连接

func (g *graph) showGraph() {
	for _, v1 := range g.mappingList {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}

func (g *graph) getFirstNeighbor(index int) int {
	for i := 0; i < g.memberNum; i++ {
		if g.mappingList[index][i] == 1 {
			return i
		}
	}
	return -1 //表示没有邻接节点了
} //得到第一个邻接节点

func (g *graph) getNextNeighbor(v1, v2 int) int {
	for i := v2 + 1; i < g.memberNum; i++ {
		if g.mappingList[v1][i] == 1 {
			return i
		}
	}
	return -1
} //得到最初传入dfs节点的下一个邻接节点

func (g *graph) dfsCore(i int) {
	fmt.Println(g.arr[i])
	g.isVisited[i] = true //确定为已经访问
	//查找节点i的第一个邻接节点w
	w := g.getFirstNeighbor(i)
	//如果i存在邻接节点
	for w != -1 {
		if !g.isVisited[w] {
			g.dfsCore(w)
		} //如果w节点未被访问过则递归
		//若w已经被访问过,那么则更新w
		w = g.getNextNeighbor(i, w)
	}
}
func (g *graph) dfs() {
	for i := range g.arr {
		if !g.isVisited[i] && g.isBonded[i] {
			g.dfsCore(i)
		}
	}
}

func (g *graph) bfsCore(i int) {

	//生成一个队列,记录节点顺序
	g.bfsQueue.addQueue(i)
	g.isVisited[i] = true
	fmt.Println(g.arr[i])
	for g.bfsQueue.len != 0 {
		u := g.bfsQueue.removeQueue() //u表示队列头节点代表的下标
		w := g.getFirstNeighbor(u)    //邻接节点的下标
		for w != -1 {
			if !g.isVisited[w] {
				fmt.Println(g.arr[w])
				g.isVisited[w] = true
				g.bfsQueue.addQueue(w)
			}
			w = g.getNextNeighbor(u, w)
		}

	}

}
func (g *graph) bfs() {
	for i := 0; i < g.memberNum; i++ {
		if g.isBonded[i] && !g.isVisited[i] {
			g.bfsCore(i)
		}
	}
}

func createGraph() graph {
	return graph{arr: make([]string, 0), mappingList: make([][]int, 0), isBonded: make([]bool, 0), bfsQueue: bfsQueue{arr: make([]int, 0)}}
}

func main() {
	graph := createGraph()
	graph.addMember("A")
	graph.addMember("B")
	graph.addMember("C")
	graph.makeBond(1, 0, 1)
	graph.makeBond(2, 1, 1)
	graph.showGraph()
	graph.dfs()
}
