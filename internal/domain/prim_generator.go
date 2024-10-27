package domain

import (
	"time"

	"golang.org/x/exp/rand"
)

type PrimGenerator struct{}

func (g *PrimGenerator) Generate(maze *Maze, startCell, endCell *Cell) *Maze {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Set start cell for generate maze
	start := NewCell(0, 0, nil)

	maze.SetGrid(0, 0, Floor)

	frontier := []*Cell{}

	// Add the start cells neighbors to the frontier
	frontier = append(frontier, NewCell(1, 0, start), NewCell(0, 1, start))

	var child, parent *Cell

	for len(frontier) > 0 {
		randIndex := rand.Intn(len(frontier))
		child = frontier[randIndex]
		frontier = append(frontier[:randIndex], frontier[randIndex+1:]...)

		parent = child.GetChild()
		row := parent.GetRow()
		col := parent.GetCol()

		if maze.IsValid(row, col) && maze.GetGrid()[row][col] == Wall {
			maze.SetGrid(child.GetRow(), child.GetCol(), Floor)
			maze.SetGrid(row, col, End)

			neighbors := maze.GetNeighbours(parent, Wall)
			frontier = append(frontier, neighbors...)

			maze.SetGrid(row, col, Floor)
		}

		maze.generateSteps = append(maze.generateSteps, maze.CopyGrid()) // Generate animation
	}

	// Set start and end
	maze.SetStart(startCell)
	maze.SetEnd(endCell)
	maze.generateSteps = append(maze.generateSteps, maze.CopyGrid())

	return maze
}
