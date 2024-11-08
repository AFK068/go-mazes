package domain

type WallFollowerSolver struct{}

const MaxSteps = 100000

func (solver *WallFollowerSolver) Solve(maze *Maze) (found bool, path []Grid, coinsCollected int) {
	path = append(path, maze.CopyGrid())

	current := maze.Start
	dirIndex := 0
	directions := []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	visited := make(map[int]bool)
	visited[maze.GetIndex(current)] = true
	steps := 0

	for current.Row != maze.End.Row || current.Col != maze.End.Col {
		if steps > MaxSteps {
			return false, path, coinsCollected
		}

		rightDirIndex := (dirIndex + 1) % 4
		rightDir := directions[rightDirIndex]
		rightRow := current.Row + rightDir.dx
		rightCol := current.Col + rightDir.dy

		frontDir := directions[dirIndex]
		frontRow := current.Row + frontDir.dx
		frontCol := current.Col + frontDir.dy

		leftDirIndex := (dirIndex + 3) % 4
		leftDir := directions[leftDirIndex]
		leftRow := current.Row + leftDir.dx
		leftCol := current.Col + leftDir.dy

		switch {
		case maze.IsValid(rightRow, rightCol) && maze.Grid[rightRow][rightCol] != Wall:
			dirIndex = rightDirIndex
			current = NewCell(rightRow, rightCol, current)
		case maze.IsValid(frontRow, frontCol) && maze.Grid[frontRow][frontCol] != Wall:
			current = NewCell(frontRow, frontCol, current)
		case maze.IsValid(leftRow, leftCol) && maze.Grid[leftRow][leftCol] != Wall:
			dirIndex = leftDirIndex
			current = NewCell(leftRow, leftCol, current)
		default:
			dirIndex = (dirIndex + 2) % 4 // Turn around
			backDir := directions[dirIndex]
			backRow := current.Row + backDir.dx
			backCol := current.Col + backDir.dy

			if maze.IsValid(backRow, backCol) && maze.Grid[backRow][backCol] == Floor {
				current = NewCell(backRow, backCol, current)
			}
		}

		if maze.Grid[current.Row][current.Col] == Money {
			coinsCollected++
		}

		if visited[maze.GetIndex(current)] {
			maze.SetGrid(current.Row, current.Col, Visited)
		} else {
			maze.SetGrid(current.Row, current.Col, Path)
		}

		visited[maze.GetIndex(current)] = true
		steps++

		path = append(path, maze.CopyGrid())
	}

	maze.SetGrid(maze.End.Row, maze.End.Col, End)
	path = append(path, maze.CopyGrid())

	return true, path, coinsCollected
}
