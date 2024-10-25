package application

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetConsoleSize(expectedWidth, expectedHeight int) error {
	var lastWidth, lastHeight int
	for {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("getting console size: %w", err)
		}

		size := strings.Split(strings.TrimSpace(string(out)), " ")
		if len(size) != 2 {
			return fmt.Errorf("unexpected output: %s", out)
		}

		height, err := strconv.Atoi(size[0])
		if err != nil {
			return fmt.Errorf("failing to parse height: %w", err)
		}

		width, err := strconv.Atoi(size[1])
		if err != nil {
			return fmt.Errorf("failing to parse width: %w", err)
		}

		if width >= expectedWidth && height >= expectedHeight {
			return nil
		}

		if width != lastWidth || height != lastHeight {
			fmt.Printf("Current console size: width = %d, height = %d. Waiting for size to be width = %d, height = %d...\n", width, height, expectedWidth, expectedHeight)
			lastWidth, lastHeight = width, height
		}
		time.Sleep(1 * time.Second)
	}
}
