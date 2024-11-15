package domain

type DFSSolver struct{}

// Dipth First Search algorithm.
func (dfs *DFSSolver) Solve(maze *Maze) (found bool, path []Grid, coinsCollected int) {
	stack := []*Cell{}
	start := maze.Start
	current := start
	stack = append(stack, start)
	visited := make(map[int]bool)
	visited[maze.GetIndex(start)] = true

	path = append(path, maze.CopyGrid())

	for maze.End.Col != current.Col || maze.End.Row != current.Row {
		if maze.NextMovePossible(stack[len(stack)-1], visited) {
			current = maze.NextFesableMove(stack[len(stack)-1], visited)
			visited[maze.GetIndex(current)] = true

			if maze.Grid[current.Row][current.Col] == Money {
				coinsCollected++
			}

			stack = append(stack, current)
			maze.SetGrid(current.Row, current.Col, Path)
			path = append(path, maze.CopyGrid())
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return false, path, coinsCollected
			}

			maze.SetGrid(current.Row, current.Col, Floor)
			current = stack[len(stack)-1]

			path = append(path, maze.CopyGrid())
		}
	}

	// Set the end cell
	maze.SetGrid(maze.End.Row, maze.End.Col, End)
	path = append(path, maze.CopyGrid())

	return true, path, coinsCollected
}
