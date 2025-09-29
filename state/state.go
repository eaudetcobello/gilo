package state

import "github.com/eaudetcobello/gilo/buffer"

type EditorState struct {
	screenWidth, screenHeight int
	quitFlag                  bool
	buffer                    *BufferState
}

type BufferState struct {
	gapBuffer *buffer.GapBuffer
}

func NewEditorState(screenWidth, screenHeight int) *EditorState {
	// TODO setup buffer see state.go
	return &EditorState{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		buffer: &BufferState{
			gapBuffer: buffer.NewGapBuffer(),
		},
	}
}

func (e *EditorState) QuitFlag() bool {
	return e.quitFlag
}

func (e *EditorState) Quit() {
	e.quitFlag = true
}
