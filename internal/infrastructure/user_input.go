package infrastructure

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/manifoldco/promptui"
)

func GetAndRoundWidthAndHeightFromUser() (width, height int, err error) {
	validate := func(input string) error {
		value, err := strconv.Atoi(input)
		if err != nil {
			slog.Error("converting width to int", slog.String("error", err.Error()))
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		if value < 8 || value > 250 {
			slog.Error("invalid width value", slog.Int("value", value))
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		return nil
	}

	promptWidth := promptui.Prompt{
		Label:    "Enter the width (minimum 8, maximum 250)",
		Validate: validate,
	}

	widthStr, err := promptWidth.Run()
	if err != nil {
		slog.Error("prompting width", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("promting width: %w", err)
	}

	width, err = strconv.Atoi(widthStr)
	if err != nil {
		slog.Error("converting width to int", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("converting width to int: %w", err)
	}

	validate = func(input string) error {
		value, err := strconv.Atoi(input)
		if err != nil {
			slog.Error("converting height to int", slog.String("error", err.Error()))
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		if value < 8 || value > 70 {
			slog.Error("invalid height", slog.String("height", input))

			return &domain.InvalidInput{Message: "Invalid input."}
		}

		return nil
	}

	promptHeight := promptui.Prompt{
		Label:    "Enter the height (minimum 8, maximum 70)",
		Validate: validate,
	}

	heightStr, err := promptHeight.Run()
	if err != nil {
		slog.Error("prompting height", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("promting height: %w", err)
	}

	height, err = strconv.Atoi(heightStr)
	if err != nil {
		slog.Error("converting height to int", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("converting height to int: %w", err)
	}

	return width, height, nil
}

func GetCoordinatesFromUser(maxValueX, maxValueY int) (x, y int, err error) {
	validate := func(input string) error {
		value, err := strconv.Atoi(input)
		if err != nil {
			slog.Error("converting input to int", slog.String("error", err.Error()))
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		if value < 1 || value > maxValueX {
			slog.Error("invalid input value", slog.String("input", input))
			return &domain.InvalidInput{Message: "Invalid input."}
		}

		return nil
	}

	promptX := promptui.Prompt{
		Label:    fmt.Sprintf("Enter the x coordinate (in the range from 1 to %d)", maxValueX),
		Validate: validate,
	}

	xStr, err := promptX.Run()
	if err != nil {
		slog.Error("prompting x coordinate", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("prompting x coordinate: %w", err)
	}

	x, err = strconv.Atoi(xStr)
	if err != nil {
		slog.Error("converting x to int", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("converting x to int: %w", err)
	}

	validate = func(input string) error {
		value, err := strconv.Atoi(input)
		if err != nil {
			slog.Error("converting input to int", slog.String("error", err.Error()))

			return &domain.InvalidInput{Message: "Invalid input."}
		}

		if value < 1 || value > maxValueY {
			slog.Error("invalid input value", slog.String("input", input))

			return &domain.InvalidInput{Message: "Invalid input."}
		}

		return nil
	}

	promptY := promptui.Prompt{
		Label:    fmt.Sprintf("Enter the y coordinate (in the range from 1 to %d)", maxValueY),
		Validate: validate,
	}

	yStr, err := promptY.Run()
	if err != nil {
		slog.Error("prompting y coordinate", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("prompting y coordinate: %w", err)
	}

	y, err = strconv.Atoi(yStr)
	if err != nil {
		slog.Error("converting y to int", slog.String("error", err.Error()))
		return 0, 0, fmt.Errorf("converting y to int: %w", err)
	}

	return x, y, nil
}
