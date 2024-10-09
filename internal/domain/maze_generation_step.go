package domain

type MazeGenerationStep struct {
	steps [][]rune
}

func NewMazeGenerationStep() *MazeGenerationStep {
	return &MazeGenerationStep{make([][]rune, 0)}
}
