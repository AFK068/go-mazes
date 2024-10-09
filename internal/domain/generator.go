package domain

type Generator interface {
	Generate(maze *Maze) *Maze
}

type GenerateMaze struct {
	generator Generator
}

func NewGenerateMaze(generator Generator) *GenerateMaze {
	return &GenerateMaze{generator}
}

func (g *GenerateMaze) GenerateMaze(rowNums, colNums int) *Maze {
	maze := NewMaze(rowNums, colNums)
	return g.generator.Generate(maze)
}

func Find(sets []*UnionFindSet, x int) int {
	if sets[x].parent != x {
		sets[x].parent = Find(sets, sets[x].parent)
	}
	return sets[x].parent
}

func Join(sets []*UnionFindSet, x, y int) {
	rootX := Find(sets, x)
	rootY := Find(sets, y)

	if rootX != rootY {
		if sets[rootX].rank > sets[rootY].rank {
			sets[rootY].parent = rootX
		} else if sets[rootX].rank < sets[rootY].rank {
			sets[rootX].parent = rootY
		} else {
			sets[rootY].parent = rootX
			sets[rootX].rank++
		}
	}
}
