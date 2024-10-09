package domain

import (
	"time"

	"golang.org/x/exp/rand"
)

type KruskalGenerator struct{}

func (g *KruskalGenerator) Generate(maze *Maze) *Maze {
	rand.Seed(uint64(time.Now().UnixNano()))

	nodes := make([]*Cell, 0)
	for i := 0; i < maze.GetRows(); i += 2 {
		for j := 0; j < maze.GetCols(); j += 2 {
			nodes = append(nodes, NewCell(i, j, nil))
		}
	}

	nodeRows := maze.GetRows() / 2
	nodeCols := maze.GetCols() / 2

	edges := make([]*Edge, 0)
	for r := 0; r < nodeRows; r++ {
		for c := 0; c < nodeCols; c++ {
			thisCell := c + (r * nodeCols)
			rightCell := (c + 1) + (r * nodeCols)
			downCell := c + ((r + 1) * nodeCols)

			if c < nodeCols-1 && nodes[thisCell].GetCol() != maze.GetCols()-1 {
				edges = append(edges, NewEdge(thisCell, rightCell))
			}
			if r < nodeRows-1 && nodes[thisCell].GetRow() != maze.GetRows()-1 {
				edges = append(edges, NewEdge(thisCell, downCell))
			}
		}
	}

	sets := make([]*UnionFindSet, 0)
	for i := 0; i < len(nodes); i++ {
		sets = append(sets, NewUnionFindSet(i))
	}

	totalEdges := 0
	for totalEdges < len(nodes)-1 {
		randIndex := rand.Intn(len(edges))
		nextEdge := edges[randIndex]
		edges = append(edges[:randIndex], edges[randIndex+1:]...)

		x := Find(sets, nextEdge.GetFirst())
		y := Find(sets, nextEdge.GetSecond())

		if x != y {
			firstRow := nodes[nextEdge.GetFirst()].GetRow()
			firstCol := nodes[nextEdge.GetFirst()].GetCol()
			secondRow := nodes[nextEdge.GetSecond()].GetRow()
			secondCol := nodes[nextEdge.GetSecond()].GetCol()
			midRow := (firstRow + secondRow) / 2
			midCol := (firstCol + secondCol) / 2

			maze.SetGrid(firstRow, firstCol, End)
			maze.SetGrid(secondRow, secondCol, End)
			maze.SetGrid(midRow, midCol, End)
			maze.Draw(10 * time.Millisecond) // Generate animation
			maze.SetGrid(firstRow, firstCol, Floor)
			maze.SetGrid(secondRow, secondCol, Floor)
			maze.SetGrid(midRow, midCol, Floor)
			Join(sets, x, y)
			totalEdges++
		}
	}

	maze.SetGrid(0, 0, Start)
	maze.SetGrid(maze.GetRows()-1, maze.GetCols()-1, End)
	maze.Draw(10 * time.Millisecond) // Generate animation

	return maze
}
