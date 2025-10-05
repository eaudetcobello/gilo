package state

import "fmt"

type EditorState struct {
	screenWidth, screenHeight int
	topLine                   int
	quitFlag                  bool
	buffer                    *BufferState
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

func (e *EditorState) EnsureCursorVisible() {
	cy, _ := e.buffer.CursorPos()

	if cy >= e.topLine+e.screenHeight {
		e.topLine = cy - e.screenHeight + 1
	}

	if cy < e.topLine {
		e.topLine = cy
	}
}

func (e *EditorState) CursorScreenPos() (x, y int) {
	line, col := e.buffer.CursorPos()
	y = line - e.topLine
	x = col + e.GutterWidth()
	return x, y
}

func (e *EditorState) GutterWidth() int {
	return len(fmt.Sprintf("%d", len(e.Buffer().Data()))) + 1
}

func (e *EditorState) VisibleLines() [][]rune {
	data := e.buffer.data
	top := e.topLine
	bottom := min(top+e.screenHeight, len(data))
	return data[top:bottom]
}

func (e *EditorState) ScreenHeight() int {
	return e.screenHeight
}

func (e *EditorState) InsertNewline() {
	e.buffer.InsertNewline()
	e.EnsureCursorVisible()
}

func (e *EditorState) MoveCursorDown() {
	e.buffer.MoveCursorDown()
	e.EnsureCursorVisible()
}

func (e *EditorState) MoveCursorUp() {
	e.buffer.MoveCursorUp()
	e.EnsureCursorVisible()
}

func (e *EditorState) TopLine() int {
	return e.topLine
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
