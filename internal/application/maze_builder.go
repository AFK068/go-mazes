package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func InitializeMaze() error {
	width, height, err := infrastructure.GetAndRoundWidthAndHeightFromUser()
	if err != nil {
		return fmt.Errorf("getting width and height: %w", err)
	}

	width = width - width%2
	height = height - height%2

	// Controling that size maze is not bigger than console size
	err = GetConsoleSize(width, height)
	if err != nil {
		return fmt.Errorf("getting console size: %w", err)
	}

	startX, startY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1)
	if err != nil {
		return fmt.Errorf("getting start coordinates: %w", err)
	}

	endX, endY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1)
	if err != nil {
		return fmt.Errorf("getting end coordinates: %w", err)
	}

	startCell := domain.NewCell(startY-1, startX-1, nil)
	endCell := domain.NewCell(endY-1, endX-1, nil)

	generator, err := selectGenerator()
	if err != nil {
		return fmt.Errorf("selecting generator: %w", err)
	}

	generateMaze := domain.NewGenerateMaze(generator)
	maze := generateMaze.GenerateMaze(height, width, startCell, endCell)
	infrastructure.RenderMazeWithGridStepsWithDelay(maze.GetMazeGenerationStep())

	solver, err := selectSolver()
	if err != nil {
		return fmt.Errorf("selecting solver: %w", err)
	}

	found, path := solver.Solve(maze)
	infrastructure.RenderMazeWithGridStepsWithDelay(&path)

	if found {
		fmt.Println("Path found!")
	} else {
		fmt.Println("No path found.")
	}

	return nil
}
