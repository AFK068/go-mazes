package infrastructure

import (
	"fmt"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func DrawMaze(mz *domain.Maze, delay time.Duration) {
	mazeGenerationStep := *mz.GetMazeGenerationStep()
	for i := 0; i < len(mazeGenerationStep); i++ {
		printGrid(&mazeGenerationStep[i])
		time.Sleep(delay)
		fmt.Print("\033[H\033[2J")
	}

	grid := mz.GetGrid()
	printGrid(&grid)
}

func printGrid(grid *domain.Grid) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			fmt.Print(string((*grid)[i][j]))
		}
		fmt.Println()
	}
}
