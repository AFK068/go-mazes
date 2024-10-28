package integration_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMaze_IsWall(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	maze.SetGrid(1, 1, domain.Wall)

	assert.True(t, maze.IsWall(1, 1), "Expected cell (1, 1) to be a wall")
	assert.False(t, maze.IsWall(-1, -1), "Expected cell (-1, -1) not to be a wall")
}

func TestMaze_IsValid(t *testing.T) {
	maze := domain.NewMaze(8, 8)

	assert.True(t, maze.IsValid(0, 0), "Expected cell (0, 0) to be valid")
	assert.True(t, maze.IsValid(3, 3), "Expected cell (3, 3) to be valid")
	assert.False(t, maze.IsValid(9, 9), "Expected cell (4, 4) to be invalid")
	assert.False(t, maze.IsValid(-1, 0), "Expected cell (-1, 0) to be invalid")
}

func TestMaze_IsPathable(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	maze.SetGrid(0, 0, domain.Floor)

	assert.True(t, maze.IsPathable(0, 0), "Expected cell (0, 0) to be pathable")
	assert.False(t, maze.IsPathable(1, 1), "Expected cell (1, 1) not to be pathable")
	assert.False(t, maze.IsPathable(4, 4), "Expected cell (4, 4) to be invalid")
}

func TestMaze_SetGrid(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	maze.SetGrid(1, 1, domain.Floor)

	assert.Equal(t, domain.Floor, maze.GetGrid()[1][1], "Expected cell (1, 1) to be set to Floor")
	maze.SetGrid(4, 4, domain.Path)
	assert.NotEqual(t, domain.MainPath, maze.GetGrid()[4][4], "Expected cell (4, 4) to be invalid and not set to MainPath")
}

func TestMaze_CopyGrid(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	maze.SetGrid(1, 1, domain.Floor)
	maze.SetGrid(2, 2, domain.Wall)

	copyGrid := maze.CopyGrid()

	assert.Equal(t, maze.GetGrid(), copyGrid, "Expected copied grid to be equal to the original grid")

	maze.SetGrid(1, 1, domain.Wall)
	assert.NotEqual(t, maze.GetGrid(), copyGrid, "Expected copied grid to remain unchanged after modifying the original grid")
}

func TestNewMaze(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	assert.Equal(t, 8, maze.GetRows(), "Expected rows to be adjusted to minimum value of 8")
	assert.Equal(t, 8, maze.GetCols(), "Expected cols to be adjusted to minimum value of 8")

	maze = domain.NewMaze(9, 11)
	assert.Equal(t, 8, maze.GetRows(), "Expected rows to be adjusted to nearest even value")
	assert.Equal(t, 10, maze.GetCols(), "Expected cols to be adjusted to nearest even value")

	maze = domain.NewMaze(8, 8)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			assert.Equal(t, domain.Wall, maze.GetGrid()[i][j], "Expected all cells to be initialized as walls")
		}
	}
}

func TestMaze_SetStart(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	startCell := domain.NewCell(1, 1, nil)
	maze.SetStart(startCell)

	assert.Equal(t, startCell, maze.GetStart(), "Expected start cell to be set correctly")
	assert.Equal(t, domain.Start, maze.GetGrid()[1][1], "Expected start cell to be marked as Start")

	newStartCell := domain.NewCell(2, 2, nil)
	maze.SetStart(newStartCell)
	assert.Equal(t, newStartCell, maze.GetStart(), "Expected new start cell to be set correctly")
	assert.Equal(t, domain.Start, maze.GetGrid()[2][2], "Expected new start cell to be marked as Start")
	assert.Equal(t, domain.Floor, maze.GetGrid()[1][1], "Expected old start cell to be marked as Floor")
}

func TestMaze_SetEnd(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	endCell := domain.NewCell(1, 1, nil)
	maze.SetEnd(endCell)

	assert.Equal(t, endCell, maze.GetEnd(), "Expected end cell to be set correctly")
	assert.Equal(t, domain.End, maze.GetGrid()[1][1], "Expected end cell to be marked as End")

	newEndCell := domain.NewCell(2, 2, nil)
	maze.SetEnd(newEndCell)
	assert.Equal(t, newEndCell, maze.GetEnd(), "Expected new end cell to be set correctly")
	assert.Equal(t, domain.End, maze.GetGrid()[2][2], "Expected new end cell to be marked as End")
	assert.Equal(t, domain.Floor, maze.GetGrid()[1][1], "Expected old end cell to be marked as Floor")
}

func TestMaze_GetNeighbours(t *testing.T) {
	maze := domain.NewMaze(8, 8)
	maze.SetGrid(0, 1, domain.Floor)
	maze.SetGrid(1, 0, domain.Floor)
	maze.SetGrid(1, 2, domain.Floor)
	maze.SetGrid(2, 1, domain.Floor)

	cell := domain.NewCell(1, 1, nil)
	neighbours := maze.GetNeighbours(cell, domain.Floor)

	expectedNeighbours := []*domain.Cell{
		domain.NewCell(0, 1, cell),
		domain.NewCell(2, 1, cell),
		domain.NewCell(1, 0, cell),
		domain.NewCell(1, 2, cell),
	}

	assert.ElementsMatch(t, expectedNeighbours, neighbours, "Expected neighbours to match")
}
