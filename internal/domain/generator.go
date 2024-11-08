package domain

type Generator interface {
	Generate(maze *Maze, start *Cell, end *Cell) *Maze
}
