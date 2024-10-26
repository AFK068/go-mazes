package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAndRoundWidthAndHeightFromUser() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	var width, height int

	for {
		fmt.Print("Enter the width (minimum 8, maximum 200): ")
		widthStr, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, fmt.Errorf("reading width: %w", err)
		}
		widthStr = strings.TrimSpace(widthStr)
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			return 0, 0, fmt.Errorf("converting width to int: %w", err)
		}
		if width >= 8 && width <= 250 {
			break
		}
		fmt.Println("Invalid width. Please enter an integer between 8 and 250.")
	}

	for {
		fmt.Print("Enter the height (minimum 8, maximum 60): ")
		heightStr, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, fmt.Errorf("reading height: %w", err)
		}
		heightStr = strings.TrimSpace(heightStr)
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			return 0, 0, fmt.Errorf("converting height to int: %w", err)
		}
		if height >= 8 && height <= 70 {
			break
		}
		fmt.Println("Invalid height. Please enter an integer between 8 and 70.")
	}

	return width, height, nil
}

func GetCoordinatesFromUser(maxValueX, maxValueY int) (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	var x, y int

	for {
		fmt.Printf("Enter the x coordinate (in the range from 1 to %d): ", maxValueX)
		xStr, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, fmt.Errorf("reading width: %w", err)
		}
		xStr = strings.TrimSpace(xStr)
		x, err = strconv.Atoi(xStr)
		if err != nil {
			return 0, 0, fmt.Errorf("converting width to int: %w", err)
		}
		if x >= 1 && x <= maxValueX {
			break
		}
		fmt.Printf("Invalid x coordinate value. Please enter an integer between 0 and %d", maxValueX)
	}

	for {
		fmt.Printf("Enter the y coordinate (in the range from 1 to %d): ", maxValueY)
		yStr, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, fmt.Errorf("reading height: %w", err)
		}
		yStr = strings.TrimSpace(yStr)
		y, err = strconv.Atoi(yStr)
		if err != nil {
			return 0, 0, fmt.Errorf("converting height to int: %w", err)
		}
		if y >= 1 && y <= maxValueY {
			break
		}
		fmt.Printf("Invalid y coordinate value. Please enter an integer between 0 and %d", maxValueY)
	}

	return x, y, nil
}
