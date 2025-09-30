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

// CursorPos returns cursorLine, cursorCol.
func (b *BufferState) CursorPos() (int, int) {
	return b.cursorLine, b.cursorCol
}

func (b *BufferState) SetCursorPos(line, col int) {
	b.cursorLine = line
	b.cursorCol = col
}

func (b *BufferState) RuneAtCursor() (rune, bool) {
	if b.cursorLine >= len(b.data) {
		return 0, false
	}

	line := b.data[b.cursorLine]
	if b.cursorCol >= len(line) {
		return 0, false
	}

	return ([]rune(line))[b.cursorCol], true
}

func (b *BufferState) MoveCursorLeft() {
	if b.cursorCol > 0 {
		b.cursorCol--
	}
}

func (b *BufferState) MoveCursorRight() {
	if b.cursorCol < len(b.data[b.cursorLine]) {
		b.cursorCol++
	}
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
