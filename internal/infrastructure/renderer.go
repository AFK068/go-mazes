package infrastructure

import (
	"fmt"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

const (
	slowTimeRender = 100 * time.Millisecond
	fastTimeRender = 30 * time.Millisecond
)

func RenderMazeWithGridStepsWithDelay(grids *[]domain.Grid) {
	if len(*grids) == 0 {
		return
	}

	// Select time delay based on the number of grids
	var timeDelay time.Duration
	if len(*grids) > 50 {
		timeDelay = fastTimeRender
	} else {
		timeDelay = slowTimeRender
	}

	for i := 0; i < len(*grids); i++ {
		printGrid(&(*grids)[i])
		time.Sleep(timeDelay)
		fmt.Print("\033[H\033[2J") // clear console
	}

	grid := (*grids)[len(*grids)-1]
	printGrid(&grid)
}

func printGrid(grid *domain.Grid) {
	rows := len(*grid)
	cols := len((*grid)[0])

	// Print top wall
	for i := 0; i < cols+1; i++ {
		fmt.Print(string(domain.Wall))
	}
	fmt.Println()

	for i := 0; i < rows; i++ {
		fmt.Print(string(domain.Wall)) // Print left wall
		for j := 0; j < cols; j++ {
			fmt.Print(string((*grid)[i][j]))
		}
		fmt.Println()
	}
}
