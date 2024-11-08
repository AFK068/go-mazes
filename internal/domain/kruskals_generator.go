package domain

import (
	"github.com/es-debug/backend-academy-2024-go-template/pkg/union"
	"golang.org/x/exp/rand"
)

type KruskalGenerator struct{}

func (g *KruskalGenerator) Generate(maze *Maze, startCell, endCell *Cell) *Maze {
	nodes := make([]*Cell, 0)

	for i := 0; i < maze.Rows; i += 2 {
		for j := 0; j < maze.Cols; j += 2 {
			nodes = append(nodes, NewCell(i, j, nil))
		}
	}

	nodeRows := maze.Rows / 2
	nodeCols := maze.Cols / 2
	edges := make([]*Edge, 0)

	for r := 0; r < nodeRows; r++ {
		for c := 0; c < nodeCols; c++ {
			thisCell := c + (r * nodeCols)
			rightCell := (c + 1) + (r * nodeCols)
			downCell := c + ((r + 1) * nodeCols)

			if c < nodeCols-1 && nodes[thisCell].Col != maze.Cols-1 {
				edges = append(edges, NewEdge(thisCell, rightCell))
			}

			if r < nodeRows-1 && nodes[thisCell].Row != maze.Rows-1 {
				edges = append(edges, NewEdge(thisCell, downCell))
			}
		}
	}

	sets := make([]*union.FindSet, 0)
	for i := 0; i < len(nodes); i++ {
		sets = append(sets, union.NewUnionFindSet(i))
	}

	totalEdges := 0
	for totalEdges < len(nodes)-1 {
		randIndex := rand.Intn(len(edges))
		nextEdge := edges[randIndex]
		edges = append(edges[:randIndex], edges[randIndex+1:]...)

		x := union.Find(sets, nextEdge.First)
		y := union.Find(sets, nextEdge.Second)

		if x != y {
			firstRow := nodes[nextEdge.First].Row
			firstCol := nodes[nextEdge.First].Col
			secondRow := nodes[nextEdge.Second].Row
			secondCol := nodes[nextEdge.Second].Col
			midRow := (firstRow + secondRow) / 2
			midCol := (firstCol + secondCol) / 2

			maze.SetGrid(firstRow, firstCol, End)
			maze.SetGrid(secondRow, secondCol, End)
			maze.SetGrid(midRow, midCol, End)
			maze.GenerateSteps = append(maze.GenerateSteps, maze.CopyGrid()) // Generate animation
			maze.SetGrid(firstRow, firstCol, Floor)
			maze.SetGrid(secondRow, secondCol, Floor)
			maze.SetGrid(midRow, midCol, Floor)
			union.Join(sets, x, y)

			totalEdges++
		}
	}

	maze.SetStart(startCell)
	maze.SetEnd(endCell)
	maze.GenerateSteps = append(maze.GenerateSteps, maze.CopyGrid()) // Generate animation

	return maze
}
