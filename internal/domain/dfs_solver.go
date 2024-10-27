package domain

type DFSSolver struct{}

// Dipth First Search algorithm.
func (dfs *DFSSolver) Solve(maze *Maze) (bool, []Grid, int) {
	stack := []*Cell{}
	start := maze.GetStart()
	current := start
	stack = append(stack, start)
	visited := make(map[int]bool)
	visited[maze.GetIndex(start)] = true

	var path []Grid
	coinsCollected := 0
	path = append(path, maze.CopyGrid())

	for maze.GetEnd().GetCol() != current.GetCol() || maze.GetEnd().GetRow() != current.GetRow() {
		if nextMovePossible(maze, stack[len(stack)-1], visited) {
			current = nextFesableMove(maze, stack[len(stack)-1], visited)
			visited[maze.GetIndex(current)] = true

			if maze.GetGrid()[current.GetRow()][current.GetCol()] == Money {
				coinsCollected++
			}

			stack = append(stack, current)
			maze.SetGrid(current.GetRow(), current.GetCol(), Path)
			path = append(path, maze.CopyGrid())
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return false, path, coinsCollected
			}

			maze.SetGrid(current.GetRow(), current.GetCol(), Floor)
			current = stack[len(stack)-1]

			path = append(path, maze.CopyGrid())
		}
	}

	// Set the end cell
	maze.SetGrid(maze.GetEnd().GetRow(), maze.GetEnd().GetCol(), End)
	path = append(path, maze.CopyGrid())

	return true, path, coinsCollected
}

// Сheck if there are possible moves from the current position.
func nextMovePossible(maze *Maze, cell *Cell, visited map[int]bool) bool {
	neighbors := maze.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[maze.GetIndex(neighbor)] {
			return true
		}
	}

	return false
}

// Take an unvisited neighbor.
func nextFesableMove(maze *Maze, cell *Cell, visited map[int]bool) *Cell {
	neighbors := maze.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[maze.GetIndex(neighbor)] {
			return neighbor
		}
	}

	return nil
}
