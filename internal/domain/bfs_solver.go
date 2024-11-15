package domain

type BFSSolver struct{}

// Breadth-first search algorithm.
func (bfs *BFSSolver) Solve(maze *Maze) (found bool, path []Grid, coinsCollected int) {
	queue := []*Cell{}

	start := maze.Start
	queue = append(queue, start)

	visited := make(map[int]bool)
	visited[maze.GetIndex(start)] = true

	path = append(path, maze.CopyGrid())

	var current *Cell

	for len(queue) != 0 {
		current = queue[0]
		queue = queue[1:]

		if maze.End.Col == current.Col && maze.End.Row == current.Row {
			for current != nil {
				maze.SetGrid(current.Row, current.Col, MainPath)
				current = current.Parent

				path = append(path, maze.CopyGrid())
			}

			maze.SetGrid(maze.Start.Row, maze.Start.Col, Start)
			maze.SetGrid(maze.End.Row, maze.End.Col, End)
			path = append(path, maze.CopyGrid())

			return true, path, coinsCollected
		}

		path = append(path, maze.CopyGrid())

		for _, neighbor := range maze.GetNeighbours(current, Floor) {
			if !visited[maze.GetIndex(neighbor)] {
				queue = append(queue, neighbor)
				visited[maze.GetIndex(neighbor)] = true

				if maze.Grid[neighbor.Row][neighbor.Col] == Money {
					coinsCollected++
				}

				maze.SetGrid(neighbor.Row, neighbor.Col, Path)
			}
		}
	}

	return false, path, coinsCollected
}
