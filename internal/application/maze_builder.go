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

	err = GetConsoleSize(width, height)
	if err != nil {
		slog.Error("getting console size", slog.String("error", err.Error()))
		fmt.Println("Failed to get console size:", err)

		return
	}

	generator, err := selectGenerator()
	if err != nil {
		slog.Error("selecting generator", slog.String("error", err.Error()))
		fmt.Println("Failed to select generator:", err)

		return
	}

	startCell, endCell, err := initializeStartAndEndCells(width, height)
	if err != nil {
		slog.Error("initializing start and end cells", slog.String("error", err.Error()))
		fmt.Println("Failed to initialize start and end cells:", err)

		return
	}

	slog.Info(
		"generating maze", slog.Int("height", height),
		slog.Int("width", width), slog.Int("start_x", startCell.Row),
		slog.Int("start_y", startCell.Col), slog.Int("end_x", endCell.Row),
		slog.Int("end_y", endCell.Col), slog.String("generator", fmt.Sprintf("%T", generator)),
	)

	maze := buildMaze(generator, height, width, startCell, endCell)
	infrastructure.RenderMazeWithGridStepsWithDelay(&maze.GenerateSteps)

	solver, err := selectSolver()
	if err != nil {
		slog.Error("selecting solver", slog.String("error", err.Error()))
		fmt.Println("Failed to select solver:", err)

		return
	}

	slog.Info("solving maze algorithm", slog.String("solver", fmt.Sprintf("%T", solver)))

	found, path, money := solver.Solve(maze)
	infrastructure.RenderMazeWithGridStepsWithDelay(&path)

	if found {
		slog.Info("path found in maze. Money collected", slog.Int("money", money))
		fmt.Println("Path found! Money collected:", money)
	} else {
		slog.Info("no path found in maze. Money collected", slog.Int("money", money))
		fmt.Println("No path found. Money collected:", money)
	}
}

func buildMaze(generator domain.Generator, rowNums, colNums int, start, end *domain.Cell) *domain.Maze {
	maze := domain.NewMaze(rowNums, colNums)
	generatedMaze := generator.Generate(maze, start, end)
	generatedMaze.GenerateMoney()

	return generatedMaze
}

func initializeStartAndEndCells(width, height int) (startCell, endCell *domain.Cell, err error) {
	startX, startY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1, "start")
	if err != nil {
		slog.Error("getting start coordinates", slog.String("error", err.Error()))

		return nil, nil, fmt.Errorf("getting start coordinates: %w", err)
	}

	endX, endY, err := infrastructure.GetCoordinatesFromUser(width-1, height-1, "end")
	if err != nil {
		slog.Error("getting end coordinates", slog.String("error", err.Error()))

		return nil, nil, fmt.Errorf("getting end coordinates: %w", err)
	}

	startCell = domain.NewCell(startY-1, startX-1, nil)
	endCell = domain.NewCell(endY-1, endX-1, nil)

	return startCell, endCell, nil
}
