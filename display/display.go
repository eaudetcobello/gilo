package display

import (
	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

func DrawEditor(screen tcell.Screen, editorState *state.EditorState) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)

	screen.Fill('x', defStyle)
}
