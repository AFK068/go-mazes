package domain

import (
	"crypto/rand"
	"math/big"
)

type Grid [][]rune

const (
	Wall     rune = '⬛'
	MainPath rune = '🟧'
	Visited  rune = '🟩'
	Path     rune = '🟪'
	Floor    rune = '⬜'
	Start    rune = '🟦'
	End      rune = '🟥'
	Money    rune = '💰'
)

type Maze struct {
	Rows          int
	Cols          int
	Grid          Grid
	GenerateSteps []Grid
	Start         *Cell
	End           *Cell
}

func NewMaze(r, w int) *Maze {
	rows := r - r%2
	cols := w - w%2

	mz := &Maze{
		rows,
		cols,
		make([][]rune, rows),
		make([]Grid, 0),
		nil,
		nil,
	}

	for i := 0; i < rows; i++ {
		tempRow := make([]rune, cols)
		for j := 0; j < cols; j++ {
			tempRow[j] = Wall
		}

		mz.Grid[i] = tempRow
	}

	return mz
}

func (mz *Maze) GetIndex(cell *Cell) int {
	return cell.Row*mz.Cols + cell.Col
}

func (mz *Maze) IsWall(r, c int) bool {
	return mz.IsValid(r, c) && mz.Grid[r][c] == Wall
}

func (mz *Maze) IsValid(r, c int) bool {
	return r >= 0 && r < mz.Rows && c >= 0 && c < mz.Cols
}

func (mz *Maze) IsPathable(r, c int) bool {
	return mz.IsValid(r, c) && mz.Grid[r][c] != Wall
}

func (mz *Maze) SetGrid(r, w int, val rune) {
	if mz.IsValid(r, w) {
		mz.Grid[r][w] = val
	}
}

func (mz *Maze) CopyGrid() Grid {
	copyGrid := make([][]rune, len(mz.Grid))
	for i := 0; i < len(mz.Grid); i++ {
		copyGrid[i] = make([]rune, len(mz.Grid[i]))
		copy(copyGrid[i], mz.Grid[i])
	}

	return copyGrid
}

func (mz *Maze) SetStart(cell *Cell) {
	r, w := cell.Row, cell.Col
	if mz.IsValid(r, w) {
		if mz.Start != nil {
			mz.Grid[mz.Start.Row][mz.Start.Col] = Floor
		}

		mz.Start = cell
		mz.Grid[r][w] = Start
	}
}

func (mz *Maze) SetEnd(cell *Cell) {
	r, w := cell.Row, cell.Col
	if mz.IsValid(r, w) {
		if mz.End != nil {
			mz.Grid[mz.End.Row][mz.End.Col] = Floor
		}

		mz.End = cell
		mz.Grid[r][w] = End
	}
}

func (mz *Maze) GetNeighbours(cell *Cell, cellType rune) []*Cell {
	neighbours := make([]*Cell, 0)

	var offsets = [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for _, offset := range offsets {
		row := cell.Row + offset[0]
		col := cell.Col + offset[1]

		if ((cellType == Floor) && mz.IsPathable(row, col)) || ((cellType == Wall) && mz.IsWall(row, col)) {
			neighbours = append(neighbours, NewCell(row, col, cell))
		}
	}

	return neighbours
}

// Сheck if there are possible moves from the current position.
func (mz *Maze) NextMovePossible(cell *Cell, visited map[int]bool) bool {
	neighbors := mz.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[mz.GetIndex(neighbor)] {
			return true
		}
	}

	return false
}

// Take an unvisited neighbor.
func (mz *Maze) NextFesableMove(cell *Cell, visited map[int]bool) *Cell {
	neighbors := mz.GetNeighbours(cell, Floor)
	for _, neighbor := range neighbors {
		if !visited[mz.GetIndex(neighbor)] {
			return neighbor
		}
	}

	return nil
}

func (mz *Maze) GenerateMoney() {
	for i := 0; i < mz.Rows; i++ {
		for j := 0; j < mz.Cols; j++ {
			if mz.Grid[i][j] == Floor {
				randBInt, _ := rand.Int(rand.Reader, big.NewInt(10))
				number := randBInt.Int64()

				if number < 1 {
					mz.SetGrid(i, j, Money)
				}
			}
		}
	}

	mz.GenerateSteps = append(mz.GenerateSteps, mz.CopyGrid())
}
