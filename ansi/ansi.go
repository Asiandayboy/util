package ansi

import "fmt"

func MoveCursor(row, col int) {
	if row < 0 {
		row = 0
	}
	if col < 0 {
		col = 0
	}

	fmt.Printf("\x1b[%d;%dH", row, col)
}

func ClearEntireScreen() {
	fmt.Print("\x1b[2J")
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func EraseEntireLine() {
	fmt.Print("\x1b[2K")
}

func EraseFromCursorToEndOfLine() {
	fmt.Print("\x1b[0K")
}

func GetArrowKeyPress(buffer []byte) string {
	if !(buffer[0] == '\x1b' && buffer[1] == '[') {
		return ""
	}

	key := buffer[2]

	switch key {
	case 'A':
		return "UP"
	case 'B':
		return "DOWN"
	case 'C':
		return "RIGHT"
	case 'D':
		return "LEFT"
	}

	return ""
}

func SetTerminalWindowTitle(title string) {
	fmt.Print("\x1b]2;" + title + "\x07")
}
