package domain

type BFSSolver struct{}

func (dfs *BFSSolver) Solve(maze *Maze) (bool, []Grid) {
	queue := []*Cell{}

	start := maze.GetStart()
	queue = append(queue, start)

	visited := make(map[int]bool)
	visited[maze.GetIndex(start)] = true

	var path []Grid
	path = append(path, maze.CopyGrid())

	var current *Cell
	for len(queue) != 0 {
		current = queue[0]
		queue = queue[1:]

		if maze.GetEnd().GetCol() == current.GetCol() && maze.GetEnd().GetRow() == current.GetRow() {
			for current != nil {
				maze.SetGrid(current.GetRow(), current.GetCol(), MainPath)
				current = current.GetParent()

				path = append(path, maze.CopyGrid())
			}

			maze.SetGrid(maze.GetStart().GetRow(), maze.GetStart().GetCol(), Start)
			maze.SetGrid(maze.GetEnd().GetRow(), maze.GetEnd().GetCol(), End)
			path = append(path, maze.CopyGrid())

			return true, path
		}

		path = append(path, maze.CopyGrid())
		for _, neighbor := range maze.GetNeighbours(current, Floor) {

			if !visited[maze.GetIndex(neighbor)] {
				queue = append(queue, neighbor)
				visited[maze.GetIndex(neighbor)] = true

				maze.SetGrid(neighbor.GetRow(), neighbor.GetCol(), Path)
			}
		}
	}

	return false, path
}
