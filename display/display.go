package display

import (
	"log"

	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

func DrawEditor(screen tcell.Screen, editorState *state.EditorState) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)

	screen.Fill(' ', defStyle)
	log.Printf("Buffer has %d lines", len(editorState.Buffer().Data()))

	for row, rowRunes := range editorState.Buffer().Data() {
		log.Printf("Line %d has %d runes: %q", row, len(rowRunes), string(rowRunes))

		for col, r := range rowRunes {
			screen.SetContent(col, row, r, nil, defStyle)
		}
	}
}
