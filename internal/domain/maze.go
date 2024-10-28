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
	rows          int
	cols          int
	grid          Grid
	generateSteps []Grid
	start         *Cell
	end           *Cell
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

		mz.grid[i] = tempRow
	}

	return mz
}

func (mz *Maze) GetRows() int {
	return mz.rows
}

func (mz *Maze) GetCols() int {
	return mz.cols
}

func (mz *Maze) GetStart() *Cell {
	return mz.start
}

func (mz *Maze) GetEnd() *Cell {
	return mz.end
}

func (mz *Maze) GetGrid() Grid {
	return mz.grid
}

func (mz *Maze) GetMazeGenerationStep() *[]Grid {
	return &mz.generateSteps
}

func (mz *Maze) GetIndex(cell *Cell) int {
	return cell.GetRow()*mz.cols + cell.GetCol()
}

func (mz *Maze) IsWall(r, c int) bool {
	return mz.IsValid(r, c) && mz.grid[r][c] == Wall
}

func (mz *Maze) IsValid(r, c int) bool {
	return r >= 0 && r < mz.rows && c >= 0 && c < mz.cols
}

func (mz *Maze) IsPathable(r, c int) bool {
	return mz.IsValid(r, c) && mz.grid[r][c] != Wall
}

func (mz *Maze) SetGrid(r, w int, val rune) {
	if mz.IsValid(r, w) {
		mz.grid[r][w] = val
	}
}

func (mz *Maze) CopyGrid() Grid {
	copyGrid := make([][]rune, len(mz.GetGrid()))
	for i := 0; i < len(mz.GetGrid()); i++ {
		copyGrid[i] = make([]rune, len(mz.GetGrid()[i]))
		copy(copyGrid[i], mz.GetGrid()[i])
	}

	return copyGrid
}

func (mz *Maze) SetStart(cell *Cell) {
	r, w := cell.GetRow(), cell.GetCol()
	if mz.IsValid(r, w) {
		if mz.start != nil {
			mz.grid[mz.start.row][mz.start.col] = Floor
		}

		mz.start = cell
		mz.grid[r][w] = Start
	}
}

func (mz *Maze) SetEnd(cell *Cell) {
	r, w := cell.GetRow(), cell.GetCol()
	if mz.IsValid(r, w) {
		if mz.end != nil {
			mz.grid[mz.end.row][mz.end.col] = Floor
		}

		mz.end = cell
		mz.grid[r][w] = End
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
		row := cell.row + offset[0]
		col := cell.col + offset[1]

		if ((cellType == Floor) && mz.IsPathable(row, col)) || ((cellType == Wall) && mz.IsWall(row, col)) {
			neighbours = append(neighbours, NewCell(row, col, cell))
		}
	}

	return neighbours
}

func (mz *Maze) GenerateMoney() {
	for i := 0; i < mz.GetRows(); i++ {
		for j := 0; j < mz.GetCols(); j++ {
			if mz.GetGrid()[i][j] == Floor {
				randBInt, _ := rand.Int(rand.Reader, big.NewInt(10))
				number := randBInt.Int64()

				if number < 1 {
					mz.SetGrid(i, j, Money)
				}
			}
		}
	}
}
