package domain

type Grid [][]rune

const (
	Wall  rune = '⬛'
	Path  rune = '🟪'
	Floor rune = '⬜'
	Start rune = '🟦'
	End   rune = '🟥'
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

	if rows < 8 {
		rows = 8
	} else if rows > 60 {
		rows = 60
	}

	if cols < 8 {
		cols = 8
	} else if cols > 200 {
		cols = 200
	}

	mz := &Maze{
		rows,
		cols,
		make([][]rune, rows),
		make([]Grid, 0),
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

func (mz *Maze) GetGrid() Grid {
	return mz.grid
}

func (mz *Maze) GetMazeGenerationStep() *[]Grid {
	return &mz.generateSteps
}

func (mz *Maze) GetIndex(cell *Cell) int {
	return cell.GetRow()*mz.cols + cell.GetCol()
}

func (mz *Maze) SetGrid(r, w int, val rune) {
	mz.grid[r][w] = val
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

func (mz *Maze) CopyGrid() Grid {
	copyGrid := make([][]rune, len(mz.GetGrid()))
	for i := 0; i < len(mz.GetGrid()); i++ {
		copyGrid[i] = make([]rune, len(mz.GetGrid()[i]))
		copy(copyGrid[i], mz.GetGrid()[i])
	}

	return copyGrid
}
