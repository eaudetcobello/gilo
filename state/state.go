package state

import "log"

type EditorState struct {
	screenWidth, screenHeight int
	quitFlag                  bool
	buffer                    *BufferState
}

type BufferState struct {
	data       [][]rune
	cursorLine int
	cursorCol  int
}


func (b *BufferState) Data() [][]rune {
	return b.data
}

func (b *BufferState) InsertRune(r rune) {
	log.Printf("inserting rune %q", r)
	
	line := b.data[b.cursorLine]
	line = append(line[:b.cursorCol], append([]rune{r}, line[b.cursorCol:]...)...)
	b.data[b.cursorLine] = line
	b.cursorCol++
}

func (b *BufferState) InsertNewline() {
	log.Printf("inserting newline")
	
	line := b.data[b.cursorLine]
	
	// split line at cursor
	beforeCursor := append([]rune{}, line[:b.cursorCol]...)
	afterCursor := append([]rune{}, line[b.cursorCol:]...)
	
	b.data[b.cursorLine] = beforeCursor
	
	// shift lines down
	newLine := b.cursorLine + 1
	b.data = append(b.data, nil)
	copy(b.data[newLine+1:], b.data[newLine:])
	b.data[newLine] = afterCursor
	
	b.cursorLine++
	b.cursorCol = 0 // TODO not gonna work with indentation
}

func (b *BufferState) SetData(data []string) {
	b.data = make([][]rune, len(data))
	for i, line := range data {
		b.data[i] = []rune(line)
	}
}

// CursorPos returns cursorLine, cursorCol.
func (b *BufferState) CursorPos() (int, int) {
	return b.cursorLine, b.cursorCol
}

func (b *BufferState) SetCursorPos(line, col int) {
	b.cursorLine = line
	b.cursorCol = col
}

func (b *BufferState) MoveCursorLeft() {
	if b.cursorCol > 0 {
		b.cursorCol--
	}
}

func (b *BufferState) MoveCursorRight() {
	// insert-mode: allows going one past the right-most rune
	if b.cursorCol < len(b.data[b.cursorLine]) {
		b.cursorCol++
	}
}

func (b *BufferState) MoveCursorDown() {
	if b.cursorLine < len(b.data)-1 {
		if b.cursorCol >= len(b.data[b.cursorLine+1]) {
			if len(b.data[b.cursorLine+1]) == 0 {
				b.cursorCol = 0
			} else {
				b.cursorCol = len(b.data[b.cursorLine+1])
			}
		}

		b.cursorLine++
	}
}

func (b *BufferState) MoveCursorUp() {
	if b.cursorLine > 0 {
		if len(b.data[b.cursorLine-1]) == 0 {
			b.cursorCol = 0
		} else if b.cursorCol > len(b.data[b.cursorLine-1]) {
			b.cursorCol = len(b.data[b.cursorLine-1])
		}

		b.cursorLine--
	}
}

func (b *BufferState) Backspace() {
	if (b.cursorLine == 0 && b.cursorCol == 0) {
		return
	}

	// join with previous line
	if b.cursorCol == 0 && b.cursorLine > 0 {
		b.cursorCol = len(b.data[b.cursorLine-1])
		b.data[b.cursorLine-1] = append(b.data[b.cursorLine-1], b.data[b.cursorLine]...)
		b.data = append(b.data[:b.cursorLine], b.data[b.cursorLine+1:]...)
		b.cursorLine--
		return
	}

	// stay on same line
	line := b.data[b.cursorLine]
	b.data[b.cursorLine] = append(line[:b.cursorCol-1], line[b.cursorCol:]...)
	b.cursorCol--
}

func NewEditorState(screenWidth, screenHeight int) *EditorState {
	initialBuff := &BufferState{
		data:       [][]rune{{}},
		cursorLine: 0,
		cursorCol:  0,
	}

	return &EditorState{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		buffer:       initialBuff,
	}
}

func (e *EditorState) Buffer() *BufferState {
	return e.buffer
}

func (e *EditorState) QuitFlag() bool {
	return e.quitFlag
}

func (e *EditorState) Quit() {
	e.quitFlag = true
}
