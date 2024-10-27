package domain

type Solver interface {
	Solve(maze *Maze) (bool, []Grid, int)
}
