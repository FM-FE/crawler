package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d\n", &row, &col)


	// make([][]int, row, col) is false,
	// the first parameter(row) means allocation space,
	// the second parameter(col) means Reserved allocation space

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			if j == 4 {
				fmt.Fscanf(file, "%d\n", &maze[i][j])
			}else {
				fmt.Fscanf(file, "%d", &maze[i][j])
			}
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point {
	{-1,0}, {0,-1}, {1,0}, {0,1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int,len(maze))
	for i:= range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if  cur == end {
			break
		}

		for _,dir := range dirs {
			next := cur.add(dir)

			// next at maze is 0
			// and next steps is 0
			// and next != start

			//if next == end { // another way to judge end
			//	return
			//}

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q,next)
		}

	}
	return steps
}

func main() {

	maze := readMaze("src/learngo/maze/maze.in")
	for _, i := range maze {
		for _, j := range i {
			fmt.Printf("%d  ",j)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, i := range steps {
		for _, j:= range i{
			fmt.Printf("%d\t",j)
		}
		fmt.Println()
	}

}
