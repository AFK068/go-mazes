package main

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
)

func main() {
	err := application.InitializeMaze()
	fmt.Println(err)
}
