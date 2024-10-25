package domain

type DFSSolver struct{}

func (dfs *DFSSolver) Solve(maze *Maze) bool {
	stack := []*Cell{}
	start := maze.GetStart()
	current := start
	stack = append(stack, start)
	visited := make(map[int]bool)
	visited[maze.GetIndex(start)] = true

	for maze.GetGrid()[start.GetRow()][start.GetCol()] != End {
		if nextMovePossible(maze, stack[len(stack)-1], visited) {
			current = nextFesableMove(maze, stack[len(stack)-1], visited)
			visited[maze.GetIndex(current)] = true
			stack = append(stack, current)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return false
			}
			current = stack[len(stack)-1]
		}
	}

	current.GetChild()
	return false
}

func nextMovePossible(maze *Maze, cell *Cell, visited map[int]bool) bool {
	neighbors := maze.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[maze.GetIndex(neighbor)] {
			return true
		}
	}

	return false
}

func nextFesableMove(maze *Maze, cell *Cell, visited map[int]bool) *Cell {
	neighbors := maze.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[maze.GetIndex(neighbor)] {
			return neighbor
		}
	}
	return nil
}
