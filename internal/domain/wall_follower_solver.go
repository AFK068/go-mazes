package domain

import "fmt"

type WallFollowerSolver struct{}

const MaxSteps = 100000

func (solver *WallFollowerSolver) Solve(maze *Maze) (bool, []Grid) {
	var path []Grid
	path = append(path, maze.CopyGrid())

	current := maze.GetStart()
	dirIndex := 0
	directions := []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	visited := make(map[int]bool)
	visited[maze.GetIndex(current)] = true
	steps := 0

	for current.GetRow() != maze.GetEnd().GetRow() || current.GetCol() != maze.GetEnd().GetCol() {
		if steps > MaxSteps {
			fmt.Println("Exceeded maximum steps, possible infinite loop.")
			return false, path
		}

		rightDirIndex := (dirIndex + 1) % 4
		rightDir := directions[rightDirIndex]
		rightRow := current.GetRow() + rightDir.dx
		rightCol := current.GetCol() + rightDir.dy

		frontDir := directions[dirIndex]
		frontRow := current.GetRow() + frontDir.dx
		frontCol := current.GetCol() + frontDir.dy

		leftDirIndex := (dirIndex + 3) % 4
		leftDir := directions[leftDirIndex]
		leftRow := current.GetRow() + leftDir.dx
		leftCol := current.GetCol() + leftDir.dy

		switch {
		case maze.IsValid(rightRow, rightCol) && maze.GetGrid()[rightRow][rightCol] != Wall:
			dirIndex = rightDirIndex
			current = NewCell(rightRow, rightCol, current)

			if visited[maze.GetIndex(current)] {
				maze.SetGrid(current.GetRow(), current.GetCol(), Visited)
			} else {
				maze.SetGrid(current.GetRow(), current.GetCol(), Path)
			}
		case maze.IsValid(frontRow, frontCol) && maze.GetGrid()[frontRow][frontCol] != Wall:
			current = NewCell(frontRow, frontCol, current)
			if visited[maze.GetIndex(current)] {
				maze.SetGrid(current.GetRow(), current.GetCol(), Visited)
			} else {
				maze.SetGrid(current.GetRow(), current.GetCol(), Path)
			}
		case maze.IsValid(leftRow, leftCol) && maze.GetGrid()[leftRow][leftCol] != Wall:
			dirIndex = leftDirIndex
			current = NewCell(leftRow, leftCol, current)

			if visited[maze.GetIndex(current)] {
				maze.SetGrid(current.GetRow(), current.GetCol(), Visited)
			} else {
				maze.SetGrid(current.GetRow(), current.GetCol(), Path)
			}
		default:
			dirIndex = (dirIndex + 2) % 4 // Turn around
			backDir := directions[dirIndex]
			backRow := current.GetRow() + backDir.dx
			backCol := current.GetCol() + backDir.dy

			if maze.IsValid(backRow, backCol) && maze.GetGrid()[backRow][backCol] == Floor {
				current = NewCell(backRow, backCol, current)
				if visited[maze.GetIndex(current)] {
					maze.SetGrid(current.GetRow(), current.GetCol(), Visited)
				} else {
					maze.SetGrid(current.GetRow(), current.GetCol(), Path)
				}
			}
		}

		visited[maze.GetIndex(current)] = true
		steps++

		path = append(path, maze.CopyGrid())
	}

	maze.SetGrid(maze.GetEnd().GetRow(), maze.GetEnd().GetCol(), End)
	path = append(path, maze.CopyGrid())

	return true, path
}
