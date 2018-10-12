package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int  {

	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row,col int
	fmt.Fscanf(file,"%d %d",&row,&col)
	maze := make([][]int,row)
	for i := range maze{
		maze[i] = make([]int,col)
		for j:= range maze[i]{
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}

type point struct {
	i,j int
}

func (p point)add (r point) point  {
	return point{p.i+r.i,p.j+r.j}
}
func (p point) at(grid [][]int) (int,bool) {
	//越界判断
	if p.i <0 || p.i >= len(grid) {
		return 0,false
	}
	if p.j <0 || p.j >= len(grid[p.i]) {
		return 0,false
	}
	return grid[p.i][p.j],true
}
var dirs = [4]point{
	//上左下右
	{-1,0},{0,-1},{1,0},{0,1},
}
//广度优先搜索走迷宫
func walk(maze [][]int,start,end point) [][]int {

	steps := make([][]int,len(maze))
	for i:= range steps {
		steps[i] = make([]int,len(maze[i]))
	}
	Q := []point{start}
	for len(Q)>0 {
		cur := Q[0]
		Q = Q[1:]
		//走到头了
		if cur == end {
			break
		}

		for _,dir := range dirs {
			next := cur.add(dir)
			/*
			走动条件
			1 maze at next is 0
			2 steps at next is 0
			3 next != start
			*/
			val,ok := next.at(maze)
			if !ok || val ==1 {
				continue
			}
			val,ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps,_ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q,next)
		}
	}
	return steps
}
func main() {
	maze:=readMaze("migong")


	for _,row := range maze{
		for _,col := range row{
			fmt.Printf("%d ",col)
		}
		fmt.Println()
	}
	fmt.Println("--------")
	steps := walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})

	for _,row := range steps{
		for _,col := range row{
			fmt.Printf("%3d",col)
		}
		fmt.Println()
	}
}
