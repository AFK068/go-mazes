package domain

type Generator interface {
	Generate(maze *Maze, start *Cell, end *Cell) *Maze
}

type GenerateMaze struct {
	generator Generator
}

func NewGenerateMaze(generator Generator) *GenerateMaze {
	return &GenerateMaze{generator}
}

func (g *GenerateMaze) GenerateMaze(rowNums, colNums int, start *Cell, end *Cell) *Maze {
	maze := NewMaze(rowNums, colNums)
	return g.generator.Generate(maze, start, end)
}
