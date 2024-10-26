package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

func displayMenu(title string, items []string) (int, error) {
	menu := pkg.NewMenu(title)
	for _, item := range items {
		menu.AddItem(item)
	}

	selectedIndex, err := menu.Display()
	if err != nil {
		return 0, fmt.Errorf("displaying menu: %w", err)
	}

	if selectedIndex >= 0 && selectedIndex < len(items) {
		return selectedIndex, nil
	} else {
		return -1, fmt.Errorf("invalid selection")
	}
}

// Displays a menu to select a maze generation algorithm and return the selected generator
func selectGenerator() (domain.Generator, error) {
	generators := []domain.Generator{
		&domain.KruskalGenerator{},
		&domain.PrimGenerator{},
	}

	selectedIndex, err := displayMenu("Select maze generation algorithm", []string{"Kruskal's algorithm", "Prim's algorithm"})
	if err != nil {
		return nil, fmt.Errorf("selecting generator: %w", err)
	}

	return generators[selectedIndex], nil
}

// Displays a menu to select a maze solving algorithm and return the selected solver
func selectSolver() (domain.Solver, error) {
	solvers := []domain.Solver{
		&domain.DFSSolver{},
		&domain.BFSSolver{},
		&domain.WallFollowerSolver{},
	}

	selectedIndex, err := displayMenu("Select maze solving algorithm", []string{"Depth-first search", "Breadth-first search", "Wall follower"})
	if err != nil {
		return nil, fmt.Errorf("selecting solver: %w", err)
	}

	return solvers[selectedIndex], nil
}
