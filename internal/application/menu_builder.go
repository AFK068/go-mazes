package application

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/menu"
)

func displayMenu(title string, items []string) (int, error) {
	mainMenu := menu.NewMenu(title)
	for _, item := range items {
		mainMenu.AddItem(item)
	}

	selectedIndex, err := mainMenu.Display()
	if err != nil {
		slog.Error("displaying menu", slog.String("error", err.Error()))
		return 0, fmt.Errorf("displaying menu: %w", err)
	}

	if selectedIndex >= 0 && selectedIndex < len(items) {
		return selectedIndex, nil
	}

	slog.Error("invalid selection index in menu", slog.Int("selected_index", selectedIndex))

	return -1, fmt.Errorf("invalid selection")
}

// Displays a menu to select a maze generation algorithm and return the selected generator.
func selectGenerator() (domain.Generator, error) {
	generators := []domain.Generator{
		&domain.KruskalGenerator{},
		&domain.PrimGenerator{},
	}

	selectedIndex, err := displayMenu("Select maze generation algorithm", []string{"Kruskal's algorithm", "Prim's algorithm"})
	if err != nil {
		slog.Error("selecting generator", slog.String("error", err.Error()))
		return nil, fmt.Errorf("selecting generator: %w", err)
	}

	return generators[selectedIndex], nil
}

// Displays a menu to select a maze solving algorithm and return the selected solver.
func selectSolver() (domain.Solver, error) {
	solvers := []domain.Solver{
		&domain.DFSSolver{},
		&domain.BFSSolver{},
		&domain.WallFollowerSolver{},
	}

	selectedIndex, err := displayMenu("Select maze solving algorithm", []string{"Depth-first search", "Breadth-first search", "Wall follower"})
	if err != nil {
		slog.Error("selecting solver", slog.String("error", err.Error()))
		return nil, fmt.Errorf("selecting solver: %w", err)
	}

	return solvers[selectedIndex], nil
}
