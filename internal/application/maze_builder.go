package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func InitializeMaze() {
	width, height, err := infrastructure.GetAndRoundWidthAndHeightFromUser()
	if err != nil {
		slog.Error("getting width and height maze size", slog.String("error", err.Error()))
		fmt.Println("Failed to get width and height from user:", err)
		return
	}

	width = width - width%2
	height = height - height%2

	// Controling that size maze is not bigger than console size
	err = GetConsoleSize(width, height)
	if err != nil {
		slog.Error("getting console size", slog.String("error", err.Error()))
		fmt.Println("Failed to get console size:", err)
		return
	}

	startX, startY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1)
	if err != nil {
		slog.Error("getting start coordinates", slog.String("error", err.Error()))
		fmt.Println("Failed to get start coordinates from user:", err)
		return
	}

	endX, endY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1)
	if err != nil {
		slog.Error("getting end coordinates", slog.String("error", err.Error()))
		fmt.Println("Failed to get end coordinates from user:", err)
		return
	}

	startCell := domain.NewCell(startY-1, startX-1, nil)
	endCell := domain.NewCell(endY-1, endX-1, nil)

	generator, err := selectGenerator()
	if err != nil {
		slog.Error("selecting generator", slog.String("error", err.Error()))
		fmt.Println("Failed to select generator:", err)
		return
	}

	slog.Info(
		"generating maze", slog.Int("height", height),
		slog.Int("width", width), slog.Int("start_x", startX),
		slog.Int("start_y", startY), slog.Int("end_x", endX),
		slog.Int("end_y", endY), slog.String("generator", fmt.Sprintf("%T", generator)),
	)

	generateMaze := domain.NewGenerateMaze(generator)
	maze := generateMaze.GenerateMaze(height, width, startCell, endCell)
	infrastructure.RenderMazeWithGridStepsWithDelay(maze.GetMazeGenerationStep())

	solver, err := selectSolver()
	if err != nil {
		slog.Error("selecting solver", slog.String("error", err.Error()))
		fmt.Println("Failed to select solver:", err)
		return
	}

	slog.Info("solving maze algorithm", slog.String("solver", fmt.Sprintf("%T", solver)))

	found, path := solver.Solve(maze)
	infrastructure.RenderMazeWithGridStepsWithDelay(&path)

	if found {
		slog.Info("path found in maze")
		fmt.Println("Path found!")
	} else {
		slog.Info("no path found in maze")
		fmt.Println("No path found.")
	}
}
