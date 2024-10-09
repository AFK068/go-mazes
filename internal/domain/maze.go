package domain

import (
	"fmt"
	"time"
)

const (
	Wall  rune = '█'
	Floor rune = ' '
	Start rune = 'S'
	End   rune = 'E'
)

type Maze struct {
	rows          int
	cols          int
	grid          [][]rune
	generateSteps MazeGenerationStep
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
		*NewMazeGenerationStep(),
		nil,
		nil,
	}

	for i := 0; i < rows; i++ {
		temp_row := make([]rune, cols)
		for j := 0; j < cols; j++ {
			temp_row[j] = Wall
		}

		mz.grid[i] = temp_row
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

func (mz *Maze) GetGrid() [][]rune {
	return mz.grid
}

func (mz *Maze) GetMazeGenerationStep() *MazeGenerationStep {
	return &mz.generateSteps
}

func (mz *Maze) SetGrid(r, w int, val rune) {
	mz.grid[r][w] = val
}

func (mz *Maze) SetStart(r, w int) {
	if mz.IsPathable(r, w) {
		if mz.start != nil {
			mz.grid[mz.start.row][mz.start.col] = Floor
		}

		mz.start = NewCell(r, w, nil)
		mz.grid[r][w] = Start
	}
}

func (mz *Maze) SetEnd(r, w int) {
	if mz.IsPathable(r, w) {
		if mz.end != nil {
			mz.grid[mz.end.row][mz.end.col] = Floor
		}

		mz.end = NewCell(r, w, nil)
		mz.grid[r][w] = End
	}
}

func (mz *Maze) GetNeighbours(cell *Cell, сellType rune) []*Cell {
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

		if ((сellType == Floor) && mz.IsPathable(row, col)) || ((сellType == Wall) && mz.IsWall(row, col)) {
			neighbours = append(neighbours, NewCell(row, col, cell))
		}
	}

	return neighbours
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

func (mz *Maze) Draw(delay time.Duration) {
	grid := mz.GetGrid()
	for i := 0; i < mz.GetRows(); i++ {
		for j := 0; j < mz.GetCols(); j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
	time.Sleep(delay)
	fmt.Print("\033[H\033[2J")
}

func (mz *Maze) PrintMaze() {
	for j := 0; j < mz.GetCols()+2; j++ {
		fmt.Print(string("_"))
	}
	println()

	grid := mz.GetGrid()
	for i := 0; i < len(grid); i++ {
		fmt.Print(string("|"))
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(string(grid[i][j]))
		}

		fmt.Println()
	}
}
