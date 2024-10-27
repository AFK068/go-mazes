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
	genratedMaze := g.generator.Generate(maze, start, end)

	// Generate money and add maze with money to generate steps.
	genratedMaze.GenerateMoney()
	maze.generateSteps = append(maze.generateSteps, maze.CopyGrid())

	return genratedMaze
}
