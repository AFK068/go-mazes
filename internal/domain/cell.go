package domain

type Cell struct {
	Row    int
	Col    int
	Parent *Cell
}

func NewCell(r, c int, prt *Cell) *Cell {
	return &Cell{
		Row:    r,
		Col:    c,
		Parent: prt,
	}
}

func (cell *Cell) GetChild() *Cell {
	return &Cell{
		Row: (cell.Row << 1) - cell.Parent.Row,
		Col: (cell.Col << 1) - cell.Parent.Col,
	}
}
