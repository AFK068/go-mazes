package application

import (
	"fmt"
	"log/slog"
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
			slog.Error("getting console size", slog.String("error", err.Error()))
			return fmt.Errorf("getting console size: %w", err)
		}

		size := strings.Split(strings.TrimSpace(string(out)), " ")
		if len(size) != 2 {
			slog.Error("unexpected output", slog.String("output", string(out)))
			return fmt.Errorf("unexpected output: %s", out)
		}

		height, err := strconv.Atoi(size[0])
		if err != nil {
			slog.Error("parsing height", slog.String("error", err.Error()))
			return fmt.Errorf("failing to parse height: %w", err)
		}

		width, err := strconv.Atoi(size[1])
		if err != nil {
			slog.Error("parsing width", slog.String("error", err.Error()))
			return fmt.Errorf("failing to parse width: %w", err)
		}

		if width >= expectedWidth && height >= expectedHeight {
			slog.Info("console size is big enough", slog.Int("width", width), slog.Int("height", height))
			return nil
		}

		if width != lastWidth || height != lastHeight {
			slog.Info("current console size", slog.Int("width", width), slog.Int("height", height))
			fmt.Printf(
				"Current console size: width = %d, height = %d. Waiting for size to be width = %d, height = %d...\n",
				width, height, expectedWidth, expectedHeight)

			lastWidth, lastHeight = width, height
		}

		time.Sleep(1 * time.Second)
	}
}
