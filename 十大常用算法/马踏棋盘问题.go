package main

import (
	"container/list"
	"fmt"
)

type point struct {
	x int
	y int
}

type horseChess struct {
	chessBoard  [][]int
	step        int  //看看走的步数有没有到棋盘数减1
	flag        bool //判断是否完成整个棋盘的walk
	boardLength int  //格子数
	visited     [][]bool
}

func createHorseChess(row int, col int) *horseChess {
	temp := make([][]bool, row)
	temp2 := make([][]int, row)
	for i := 0; i < row; i++ {
		temp[i] = make([]bool, col)
		temp2[i] = make([]int, col)
	}
	return &horseChess{chessBoard: temp2, boardLength: row * col, step: 1, visited: temp}
}

func (chess *horseChess) getNextPosition(now point) *list.List {
	res := list.New()
	if now.x-2 >= 0 && now.y-1 >= 0 {
		res.PushBack(point{x: now.x - 2, y: now.y - 1})
	}

	if now.x-1 >= 0 && now.y-2 >= 0 {
		res.PushBack(point{x: now.x - 1, y: now.y - 2})
	}

	if now.x-2 >= 0 && now.y+1 < len(chess.chessBoard) {
		res.PushBack(point{x: now.x - 2, y: now.y + 1})
	}

	if now.x-1 >= 0 && now.y+2 < len(chess.chessBoard) {
		res.PushBack(point{x: now.x - 1, y: now.y + 2})
	}

	if now.x+2 < len(chess.chessBoard[0]) && now.y-1 >= 0 {
		res.PushBack(point{x: now.x + 2, y: now.y - 1})
	}

	if now.x+1 < len(chess.chessBoard[0]) && now.y-2 >= 0 {
		res.PushBack(point{x: now.x + 1, y: now.y - 2})
	}

	if now.x+2 < len(chess.chessBoard[0]) && now.y+1 < len(chess.chessBoard) {
		res.PushBack(point{x: now.x + 2, y: now.y + 1})
	}

	if now.x+1 < len(chess.chessBoard[0]) && now.y+2 < len(chess.chessBoard) {
		res.PushBack(point{x: now.x + 1, y: now.y + 2})
	}
	return res
} //返回一个保存着下一步可以走的point切片

func (chess *horseChess) tryWalk(begin point, step int) {
	chess.chessBoard[begin.x][begin.y] = step //计步
	chess.visited[begin.x][begin.y] = true    //标记为已经访问过

	next := chess.getNextPosition(begin) //获取接下来可以走的可能性

	for next.Len() != 0 {
		tmp := next.Remove(next.Front())
		p, ok := tmp.(point)
		if ok {
			if !chess.visited[p.x][p.y] { //递归
				chess.tryWalk(p, step+1)
			}
		}
	}

	if step < chess.boardLength && !chess.flag {
		chess.chessBoard[begin.x][begin.y] = 0
		chess.visited[begin.x][begin.y] = false
	} else {
		chess.flag = true //完成
	}
}

func (chess *horseChess) show() {
	for _, v1 := range chess.chessBoard {
		fmt.Println()
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
	}
}

func main() {
	chess := createHorseChess(9, 9)
	chess.tryWalk(point{x: 0, y: 0}, 1)
	//fmt.Println(chess.chessBoard)
	chess.show()
}
