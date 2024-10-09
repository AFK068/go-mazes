package domain

import (
	"time"

	"golang.org/x/exp/rand"
)

type PrimGenerator struct{}

func (g *PrimGenerator) Generate(maze *Maze) *Maze {
	rand.Seed(uint64(time.Now().UnixNano()))

	start := NewCell(0, 0, nil)
	maze.SetGrid(start.GetRow(), start.GetCol(), Start)

	frontier := []*Cell{
		NewCell(1, 0, start),
		NewCell(0, 1, start),
	}

	var child *Cell
	var parent *Cell

	for len(frontier) > 0 {
		randIndex := rand.Intn(len(frontier))
		child = frontier[randIndex]
		frontier = append(frontier[:randIndex], frontier[randIndex+1:]...)

		parent = child.GetChild()
		row := parent.GetRow()
		col := parent.GetCol()

		if (maze.IsValid(row, col)) && (maze.GetGrid()[row][col] == Wall) {
			maze.SetGrid(child.GetRow(), child.GetCol(), Floor)
			maze.SetGrid(row, col, End)

			neighbors := maze.GetNeighbours(parent, Wall)
			frontier = append(frontier, neighbors...)

			maze.SetGrid(row, col, Floor)
		}

		maze.Draw(10 * time.Millisecond) // Generate animation
		// Draw(maze, 10*time.Millisecond) // Generate animation
	}

	maze.SetGrid(maze.GetRows()-1, maze.GetCols()-1, End)

	return maze
}
