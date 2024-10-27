package integration_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBFSSolver_success(t *testing.T) {
	mazes := []struct {
		setup         func() *domain.Maze
		expectedCoins []int
	}{
		{
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(0, 0, nil))
				maze.SetEnd(domain.NewCell(2, 0, nil))

				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Wall)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Wall)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 1, domain.Wall)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Floor)
				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)
				maze.SetGrid(3, 3, domain.Floor)

				maze.SetGrid(0, 2, domain.Money)
				maze.SetGrid(2, 3, domain.Money)

				return maze
			},
			expectedCoins: []int{2},
		},
		{
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(3, 3, nil))
				maze.SetEnd(domain.NewCell(2, 1, nil))

				maze.SetGrid(0, 0, domain.Floor)
				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Floor)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Floor)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 0, domain.Floor)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Floor)
				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)

				maze.SetGrid(0, 0, domain.Money)

				return maze
			},
			expectedCoins: []int{0},
		},
	}

	for _, tc := range mazes {
		maze := tc.setup()
		solver := &domain.BFSSolver{}

		found, _, coinsCollected := solver.Solve(maze)

		assert.True(t, found, "Expected to find a path, but no path was found")
		assert.Contains(
			t, tc.expectedCoins, coinsCollected,
			"Expected to collect one of %v coins, but collected %d",
			tc.expectedCoins, coinsCollected)
	}
}

func TestBFSSolver_failure(t *testing.T) {
	maze := domain.NewMaze(4, 4)

	maze.SetStart(domain.NewCell(0, 0, nil))
	maze.SetEnd(domain.NewCell(2, 0, nil))

	maze.SetGrid(0, 1, domain.Floor)
	maze.SetGrid(0, 2, domain.Floor)
	maze.SetGrid(0, 3, domain.Floor)

	maze.SetGrid(1, 0, domain.Wall)
	maze.SetGrid(1, 1, domain.Wall)
	maze.SetGrid(1, 2, domain.Wall)
	maze.SetGrid(1, 3, domain.Floor)

	maze.SetGrid(2, 1, domain.Wall)
	maze.SetGrid(2, 2, domain.Wall)
	maze.SetGrid(2, 3, domain.Floor)

	maze.SetGrid(3, 0, domain.Wall)
	maze.SetGrid(3, 1, domain.Wall)
	maze.SetGrid(3, 2, domain.Wall)
	maze.SetGrid(3, 3, domain.Floor)

	maze.SetGrid(0, 2, domain.Money)
	maze.SetGrid(3, 1, domain.Money)

	solver := &domain.BFSSolver{}

	found, _, coinsCollected := solver.Solve(maze)

	assert.False(t, found, "Expected not to find a path, but a path was found")

	expectedCoins := 1
	assert.Equal(t, expectedCoins, coinsCollected, "Expected to collect %d coins, but collected %d", expectedCoins, coinsCollected)

	assert.Equal(t, domain.Start, maze.GetGrid()[0][0], "Expected start cell to be %c, but got %c", domain.Start, maze.GetGrid()[0][0])
	assert.Equal(t, domain.End, maze.GetGrid()[2][0], "Expected end cell to be %c, but got %c", domain.End, maze.GetGrid()[2][0])
}

func TestDFSSolver_success(t *testing.T) {
	mazes := []struct {
		setup         func() *domain.Maze
		expectedCoins []int
	}{
		{
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(0, 0, nil))
				maze.SetEnd(domain.NewCell(2, 0, nil))

				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Wall)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Wall)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 1, domain.Wall)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Floor)
				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)
				maze.SetGrid(3, 3, domain.Floor)

				maze.SetGrid(0, 2, domain.Money)
				maze.SetGrid(2, 3, domain.Money)

				return maze
			},
			expectedCoins: []int{2},
		},
		{
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(3, 3, nil))
				maze.SetEnd(domain.NewCell(2, 1, nil))

				maze.SetGrid(0, 0, domain.Floor)
				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Floor)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Floor)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 0, domain.Floor)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Floor)
				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)

				maze.SetGrid(0, 0, domain.Money)

				return maze
			},
			expectedCoins: []int{0, 1},
		},
	}

	for _, tc := range mazes {
		maze := tc.setup()
		solver := &domain.DFSSolver{}

		found, _, coinsCollected := solver.Solve(maze)

		assert.True(t, found, "Expected to find a path, but no path was found")
		assert.Contains(t, tc.expectedCoins, coinsCollected,
			"Expected to collect one of %v coins, but collected %d",
			tc.expectedCoins, coinsCollected)
	}
}
func TestDFSSolver_failure(t *testing.T) {
	maze := domain.NewMaze(4, 4)

	maze.SetStart(domain.NewCell(0, 0, nil))
	maze.SetEnd(domain.NewCell(2, 0, nil))

	maze.SetGrid(0, 1, domain.Floor)
	maze.SetGrid(0, 2, domain.Floor)
	maze.SetGrid(0, 3, domain.Floor)

	maze.SetGrid(1, 0, domain.Wall)
	maze.SetGrid(1, 1, domain.Wall)
	maze.SetGrid(1, 2, domain.Wall)
	maze.SetGrid(1, 3, domain.Floor)

	maze.SetGrid(2, 1, domain.Wall)
	maze.SetGrid(2, 2, domain.Wall)
	maze.SetGrid(2, 3, domain.Floor)

	maze.SetGrid(3, 0, domain.Wall)
	maze.SetGrid(3, 1, domain.Wall)
	maze.SetGrid(3, 2, domain.Wall)
	maze.SetGrid(3, 3, domain.Floor)

	maze.SetGrid(0, 2, domain.Money)
	maze.SetGrid(3, 1, domain.Money)

	solver := &domain.DFSSolver{}

	found, _, coinsCollected := solver.Solve(maze)

	assert.False(t, found, "Expected not to find a path, but a path was found")

	expectedCoins := 1
	assert.Equal(t, expectedCoins, coinsCollected, "Expected to collect %d coins, but collected %d", expectedCoins, coinsCollected)

	assert.Equal(t, domain.Start, maze.GetGrid()[0][0], "Expected start cell to be %c, but got %c", domain.Start, maze.GetGrid()[0][0])
	assert.Equal(t, domain.End, maze.GetGrid()[2][0], "Expected end cell to be %c, but got %c", domain.End, maze.GetGrid()[2][0])
}

func TestWallFollowerSolver_success(t *testing.T) {
	maze := domain.NewMaze(4, 4)

	maze.SetStart(domain.NewCell(0, 0, nil))
	maze.SetEnd(domain.NewCell(2, 0, nil))

	maze.SetGrid(0, 1, domain.Floor)
	maze.SetGrid(0, 2, domain.Floor)
	maze.SetGrid(0, 3, domain.Floor)

	maze.SetGrid(1, 0, domain.Wall)
	maze.SetGrid(1, 1, domain.Wall)
	maze.SetGrid(1, 2, domain.Wall)
	maze.SetGrid(1, 3, domain.Floor)

	maze.SetGrid(2, 1, domain.Wall)
	maze.SetGrid(2, 2, domain.Wall)
	maze.SetGrid(2, 3, domain.Floor)

	maze.SetGrid(3, 0, domain.Floor)
	maze.SetGrid(3, 1, domain.Floor)
	maze.SetGrid(3, 2, domain.Floor)
	maze.SetGrid(3, 3, domain.Floor)

	maze.SetGrid(0, 2, domain.Money)
	maze.SetGrid(2, 3, domain.Money)

	solver := &domain.WallFollowerSolver{}

	found, _, coinsCollected := solver.Solve(maze)

	assert.True(t, found, "Expected to find a path, but no path was found")

	expectedCoins := 2
	assert.Equal(t, expectedCoins, coinsCollected, "Expected to collect %d coins, but collected %d", expectedCoins, coinsCollected)

	assert.Equal(t, domain.Start, maze.GetGrid()[0][0], "Expected start cell to be %c, but got %c", domain.Start, maze.GetGrid()[0][0])
	assert.Equal(t, domain.End, maze.GetGrid()[2][0], "Expected end cell to be %c, but got %c", domain.End, maze.GetGrid()[2][0])
}

func TestWallFollowerSolver_failure(t *testing.T) {
	mazes := []struct {
		setup         func() *domain.Maze
		expectedCoins []int
	}{
		{
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(0, 0, nil))
				maze.SetEnd(domain.NewCell(2, 0, nil))

				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Wall)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Wall)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 1, domain.Wall)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Wall)
				maze.SetGrid(3, 1, domain.Wall)
				maze.SetGrid(3, 2, domain.Wall)
				maze.SetGrid(3, 3, domain.Floor)

				maze.SetGrid(0, 2, domain.Money)
				maze.SetGrid(3, 1, domain.Money)

				return maze
			},
			expectedCoins: []int{1},
		},
		{
			// Test by infinity loop.
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(3, 3, nil))
				maze.SetEnd(domain.NewCell(2, 1, nil))

				maze.SetGrid(0, 0, domain.Floor)
				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Floor)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Floor)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 2, domain.Floor)
				maze.SetGrid(1, 3, domain.Floor)

				maze.SetGrid(2, 0, domain.Floor)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 0, domain.Floor)
				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)

				return maze
			},
			expectedCoins: []int{0},
		}, {
			setup: func() *domain.Maze {
				maze := domain.NewMaze(4, 4)

				maze.SetStart(domain.NewCell(1, 2, nil))
				maze.SetEnd(domain.NewCell(3, 0, nil))

				maze.SetGrid(0, 0, domain.Floor)
				maze.SetGrid(0, 1, domain.Floor)
				maze.SetGrid(0, 2, domain.Wall)
				maze.SetGrid(0, 3, domain.Floor)

				maze.SetGrid(1, 0, domain.Floor)
				maze.SetGrid(1, 1, domain.Wall)
				maze.SetGrid(1, 3, domain.Wall)

				maze.SetGrid(2, 0, domain.Floor)
				maze.SetGrid(2, 1, domain.Floor)
				maze.SetGrid(2, 2, domain.Wall)
				maze.SetGrid(2, 3, domain.Floor)

				maze.SetGrid(3, 1, domain.Floor)
				maze.SetGrid(3, 2, domain.Floor)
				maze.SetGrid(3, 3, domain.Floor)

				return maze
			},
			expectedCoins: []int{0},
		},
	}

	for _, tc := range mazes {
		maze := tc.setup()
		solver := &domain.WallFollowerSolver{}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		found, _, coinsCollected := solver.Solve(maze)
		err := w.Close()
		require.NoError(t, err)

		os.Stdout = old

		var buf bytes.Buffer
		_, err = io.Copy(&buf, r)
		require.NoError(t, err)

		assert.False(t, found, "Expected not to find a path, but a path was found")
		assert.Contains(t, tc.expectedCoins, coinsCollected,
			"Expected to collect one of %v coins, but collected %d",
			tc.expectedCoins, coinsCollected)

		if tc.expectedCoins[0] == 0 {
			assert.Contains(t, buf.String(),
				"Exceeded maximum steps, possible infinite loop.",
				"Expected to see 'Exceeded maximum steps, possible infinite loop.' in output")
		}
	}
}
