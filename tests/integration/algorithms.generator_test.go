package integration_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestPrimGenerator_Generate(t *testing.T) {
	rand.Seed(42)

	for i := 0; i < 1000; i++ {
		maze := domain.NewMaze(12, 12)
		startRow, startCol := rand.Intn(11), rand.Intn(11)
		endRow, endCol := rand.Intn(11), rand.Intn(11)

		startCell := domain.NewCell(startRow, startCol, nil)
		endCell := domain.NewCell(endRow, endCol, nil)

		generator := &domain.PrimGenerator{}
		generatedMaze := generator.Generate(maze, startCell, endCell)

		assert.Equal(t, startCell, generatedMaze.GetStart(), "Expected start cell to be set correctly")
		assert.Equal(t, endCell, generatedMaze.GetEnd(), "Expected end cell to be set correctly")

		solver := &domain.BFSSolver{}
		found, _, _ := solver.Solve(generatedMaze)

		assert.True(t, found, "Expected to find a path from start to end")
	}
}

func TestKruskalsGenerator_Generate(t *testing.T) {
	rand.Seed(42)

	for i := 0; i < 100; i++ {
		maze := domain.NewMaze(12, 12)
		startRow, startCol := rand.Intn(11), rand.Intn(11)
		endRow, endCol := rand.Intn(11), rand.Intn(11)

		startCell := domain.NewCell(startRow, startCol, nil)
		endCell := domain.NewCell(endRow, endCol, nil)

		generator := &domain.KruskalGenerator{}
		generatedMaze := generator.Generate(maze, startCell, endCell)

		assert.Equal(t, startCell, generatedMaze.GetStart(), "Expected start cell to be set correctly")
		assert.Equal(t, endCell, generatedMaze.GetEnd(), "Expected end cell to be set correctly")

		solver := &domain.BFSSolver{}
		found, _, _ := solver.Solve(generatedMaze)

		assert.True(t, found, "Expected to find a path from start to end")
	}
}
