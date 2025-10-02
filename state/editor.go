package state

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

func (e *EditorState) adjustViewport() {
	cursorLine, _ := e.buffer.CursorPos()

	if cursorLine >= e.topLine+e.screenHeight {
		e.topLine = cursorLine - e.screenHeight + 1
	}

	if cursorLine < e.topLine {
		e.topLine = cursorLine
	}
}

func (e *EditorState) ScreenHeight() int {
	return e.screenHeight
}

func (e *EditorState) MoveCursorDown() {
	e.buffer.MoveCursorDown()
	e.adjustViewport()
}

func (e *EditorState) MoveCursorUp() {
	e.buffer.MoveCursorUp()
	e.adjustViewport()
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
