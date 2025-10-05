package state

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type BufferState struct {
	data [][]rune
	cy   int
	cx   int
}

func (b *BufferState) Data() [][]rune {
	return b.data
}

func (b *BufferState) InsertRune(r rune) {
	log.Printf("inserting rune %q", r)

	line := b.data[b.cy]
	line = append(line[:b.cx], append([]rune{r}, line[b.cx:]...)...)
	b.data[b.cy] = line
	b.cx++
}

func (b *BufferState) InsertNewline() {
	log.Printf("inserting newline")

	line := b.data[b.cy]

	// split line at cursor
	beforeCursor := append([]rune{}, line[:b.cx]...)
	afterCursor := append([]rune{}, line[b.cx:]...)

	b.data[b.cy] = beforeCursor

	// shift lines down
	newLine := b.cy + 1
	b.data = append(b.data, nil)
	copy(b.data[newLine+1:], b.data[newLine:])
	b.data[newLine] = afterCursor

	b.cy++
	b.cx = 0 // TODO not gonna work with indentation
}

func (b *BufferState) SetData(data []string) {
	b.data = make([][]rune, len(data))
	for i, line := range data {
		b.data[i] = []rune(line)
	}
}

// CursorPos returns cursorLine, cursorCol.
func (b *BufferState) CursorPos() (int, int) {
	return b.cy, b.cx
}

func (b *BufferState) SetCursorPos(line, col int) {
	b.cy = line
	b.cx = col
}

func (b *BufferState) MoveCursorLeft() {
	if b.cx > 0 {
		b.cx--
	}
}

func (b *BufferState) MoveCursorRight() {
	// insert-mode: allows going one past the right-most rune
	if b.cx < len(b.data[b.cy]) {
		b.cx++
	}
}

func (b *BufferState) MoveCursorDown() {
	if b.cy < len(b.data)-1 {
		if b.cx >= len(b.data[b.cy+1]) {
			if len(b.data[b.cy+1]) == 0 {
				b.cx = 0
			} else {
				b.cx = len(b.data[b.cy+1])
			}
		}

		b.cy++
	}
}

func (b *BufferState) MoveCursorUp() {
	if b.cy > 0 {
		if len(b.data[b.cy-1]) == 0 {
			b.cx = 0
		} else if b.cx > len(b.data[b.cy-1]) {
			b.cx = len(b.data[b.cy-1])
		}

		b.cy--
	}
}

func (b *BufferState) Backspace() {
	if b.cy == 0 && b.cx == 0 {
		return
	}

	// join with previous line
	if b.cx == 0 && b.cy > 0 {
		b.cx = len(b.data[b.cy-1])
		b.data[b.cy-1] = append(b.data[b.cy-1], b.data[b.cy]...)
		b.data = append(b.data[:b.cy], b.data[b.cy+1:]...)
		b.cy--
		return
	}

	// stay on same line
	line := b.data[b.cy]
	b.data[b.cy] = append(line[:b.cx-1], line[b.cx:]...)
	b.cx--
}

func (b *BufferState) LoadFromFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	content := strings.TrimRight(string(file), "\n")
	b.SetData(strings.Split(content, "\n"))

	return nil
}
