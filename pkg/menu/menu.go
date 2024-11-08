package menu

import (
	"fmt"

	"github.com/buger/goterm"
	"github.com/pkg/term"
)

const (
	up    byte = 65
	down  byte = 66
	enter byte = 13
)

var keys = map[byte]bool{
	up:   true,
	down: true,
}

type Menu struct {
	Prompt         string
	CursorPosition int
	MenuItems      []*Item
}

type Item struct {
	Text    string
	SubMenu *Menu
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:    prompt,
		MenuItems: make([]*Item, 0),
	}
}

func (m *Menu) AddItem(option string) *Menu {
	menuItem := &Item{
		Text: option,
	}

	m.MenuItems = append(m.MenuItems, menuItem)

	return m
}

func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			newline = ""
		}

		menuItemText := menuItem.Text
		cursor := "  "

		if index == m.CursorPosition {
			cursor = goterm.Color("> ", goterm.YELLOW)
			menuItemText = goterm.Color(menuItemText, goterm.YELLOW)
		}

		fmt.Printf("\r%s %s%s", cursor, menuItemText, newline)
	}
}

func (m *Menu) Display() (int, error) {
	if len(m.MenuItems) == 0 {
		return 0, nil
	}

	defer fmt.Printf("\033[?25h")

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	for {
		keyCode, err := getInput()
		if err != nil {
			return 0, fmt.Errorf("getting input: %w", err)
		}

		switch keyCode {
		case enter:
			fmt.Println("\r")
			return m.CursorPosition, nil
		case up:
			m.CursorPosition = (m.CursorPosition + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		case down:
			m.CursorPosition = (m.CursorPosition + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		}
	}
}

func getInput() (byte, error) {
	t, err := term.Open("/dev/tty")
	if err != nil {
		return 0, fmt.Errorf("opening terminal: %w", err)
	}

	err = term.RawMode(t)
	if err != nil {
		return 0, fmt.Errorf("setting raw mode: %w", err)
	}

	readBytes := make([]byte, 3)

	read, err := t.Read(readBytes)
	if err != nil {
		return 0, fmt.Errorf("reading from terminal: %w", err)
	}

	err = t.Restore()
	if err != nil {
		return 0, fmt.Errorf("restoring terminal: %w", err)
	}

	err = t.Close()
	if err != nil {
		return 0, fmt.Errorf("closing terminal: %w", err)
	}

	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2], nil
		}
	} else {
		return readBytes[0], nil
	}

	return 0, nil
}
