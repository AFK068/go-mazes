package domain

import (
	"golang.org/x/exp/rand"
)

type PrimGenerator struct{}

func (g *PrimGenerator) Generate(maze *Maze, startCell, endCell *Cell) *Maze {
	// Set start cell for generate maze.
	start := NewCell(0, 0, nil)

	maze.SetGrid(0, 0, Floor)

	frontier := []*Cell{}

	// Add the start cells neighbors to the frontier.
	frontier = append(frontier, NewCell(1, 0, start), NewCell(0, 1, start))

	var child, parent *Cell

	for len(frontier) > 0 {
		randIndex := rand.Intn(len(frontier))
		child = frontier[randIndex]
		frontier = append(frontier[:randIndex], frontier[randIndex+1:]...)

		parent = child.GetChild()
		row := parent.Row
		col := parent.Col

		if maze.IsValid(row, col) && maze.Grid[row][col] == Wall {
			maze.SetGrid(child.Row, child.Col, Floor)
			maze.SetGrid(row, col, End)

			neighbors := maze.GetNeighbours(parent, Wall)
			frontier = append(frontier, neighbors...)

			maze.SetGrid(row, col, Floor)
		}

		maze.GenerateSteps = append(maze.GenerateSteps, maze.CopyGrid()) // Generate animation
	}

	// Set start and end
	maze.SetStart(startCell)
	maze.SetEnd(endCell)
	maze.GenerateSteps = append(maze.GenerateSteps, maze.CopyGrid())

	return maze
}
