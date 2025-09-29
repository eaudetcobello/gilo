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

func NewEditorState(screenWidth, screenHeight int) *EditorState {
	// TODO setup buffer see state.go
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

func (e *EditorState) InsertRune(r rune) {
	log.Printf("inserting rune %q", r)

	line := e.buffer.data[e.buffer.cursorLine]
	line = append(line[:e.buffer.cursorCol], append([]rune{r}, line[e.buffer.cursorCol:]...)...)
	e.buffer.data[e.buffer.cursorLine] = line
	e.buffer.cursorCol++
}

func (e *EditorState) InsertNewline() {
	log.Printf("inserting newline")

	line := e.buffer.data[e.buffer.cursorLine]

	// split line at cursor
	beforeCursor := append([]rune{}, line[:e.buffer.cursorCol]...)
	afterCursor := append([]rune{}, line[e.buffer.cursorCol:]...)

	e.buffer.data[e.buffer.cursorLine] = beforeCursor

	newLine := e.buffer.cursorLine + 1
	e.buffer.data = append(e.buffer.data, nil)
	copy(e.buffer.data[newLine+1:], e.buffer.data[newLine:])
	e.buffer.data[newLine] = afterCursor

	e.buffer.cursorLine++
	e.buffer.cursorCol = 0
}
