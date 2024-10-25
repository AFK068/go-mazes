package domain

import (
	"time"

	"golang.org/x/exp/rand"
)

type PrimGenerator struct{}

func (g *PrimGenerator) Generate(maze *Maze, startCell, endCell *Cell) *Maze {
	rand.Seed(uint64(time.Now().UnixNano()))

	maze.SetStart(startCell)

	frontier := []*Cell{}
	if startCell.GetRow()+1 < maze.GetRows() {
		frontier = append(frontier, NewCell(startCell.GetRow()+1, startCell.GetCol(), startCell))
	} else if startCell.GetRow()-1 >= 0 {
		frontier = append(frontier, NewCell(startCell.GetRow()-1, startCell.GetCol(), startCell))
	}
	if startCell.GetCol()+1 < maze.GetCols() {
		frontier = append(frontier, NewCell(startCell.GetRow(), startCell.GetCol()+1, startCell))
	} else if startCell.GetCol()-1 >= 0 {
		frontier = append(frontier, NewCell(startCell.GetRow(), startCell.GetCol()-1, startCell))
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

		if maze.IsValid(row, col) && maze.GetGrid()[row][col] == Wall {
			maze.SetGrid(child.GetRow(), child.GetCol(), Floor)
			maze.SetGrid(row, col, End)

			neighbors := maze.GetNeighbours(parent, Wall)
			frontier = append(frontier, neighbors...)

			maze.SetGrid(row, col, Floor)
		}

		maze.generateSteps = append(maze.generateSteps, maze.CopyGrid()) // Generate animation
	}

	maze.SetEnd(endCell)

	return maze
}
