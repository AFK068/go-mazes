package main

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func main() {
	mazeKruskal := domain.KruskalGenerator{}
	//mazePrims := domain.PrimGenerator{}

	generateMazeKruskal := domain.NewGenerateMaze(&mazeKruskal)
	//generateMazePrims := domain.NewGenerateMaze(&mazePrims)

	maze := generateMazeKruskal.GenerateMaze(15, 50)
	//maze := generateMazePrims.GenerateMaze(15, 50)

	maze.PrintMaze()
}
