package application

import (
	"fmt"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

func InitializeMaze() error {
	width, height, err := infrastructure.GetWidthAndHeightFromUser()
	if err != nil {
		return fmt.Errorf("getting width and height: %w", err)
	}

	// Controling that size maze is not bigger than console size
	err = GetConsoleSize(width, height)
	if err != nil {
		return fmt.Errorf("getting console size: %w", err)
	}

	startX, startY, err := infrastructure.GetCoordinatesFromUser(width, height)
	if err != nil {
		return fmt.Errorf("getting start coordinates: %w", err)
	}

	endX, endY, err := infrastructure.GetCoordinatesFromUser(width, height)
	if err != nil {
		return fmt.Errorf("getting end coordinates: %w", err)
	}

	startCell := domain.NewCell(startY, startX, nil)
	endCell := domain.NewCell(endY, endX, nil)

	generateMazeKruskal := domain.NewGenerateMaze(&domain.KruskalGenerator{})
	generateMazePrims := domain.NewGenerateMaze(&domain.PrimGenerator{})

	actionsMaze := []func() domain.Maze{
		func() domain.Maze {
			maze := generateMazeKruskal.GenerateMaze(height, width, startCell, endCell)
			infrastructure.DrawMaze(maze, 10*time.Millisecond)
			return *maze
		},
		func() domain.Maze {
			maze := generateMazePrims.GenerateMaze(height, width, startCell, endCell)
			infrastructure.DrawMaze(maze, 10*time.Millisecond)
			return *maze
		},
	}

	menu := pkg.NewMenu("Select algorithm")
	menu.AddItem("Kruskal's algorithm")
	menu.AddItem("Prim's algorithm")

	selectedIndex, _ := menu.Display()
	if selectedIndex >= 0 && selectedIndex < len(actionsMaze) {
		actionsMaze[selectedIndex]()
	} else {
		fmt.Println("Invalid selection")
	}

	fmt.Println(actionsMaze)
	return nil
}
