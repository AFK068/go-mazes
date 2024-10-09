package domain

type Cell struct {
	row    int
	col    int
	parent *Cell
}

func (cell *Cell) GetRow() int {
	return cell.row
}

func (cell *Cell) GetCol() int {
	return cell.col
}

func (cell *Cell) GetParent() *Cell {
	return cell.parent
}

func NewCell(r, c int, prt *Cell) *Cell {
	return &Cell{
		row:    r,
		col:    c,
		parent: prt,
	}
}

func (cell *Cell) GetChild() *Cell {
	return &Cell{
		row: (cell.row << 1) - cell.parent.row,
		col: (cell.col << 1) - cell.parent.col,
	}
}
