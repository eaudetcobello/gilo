package display

import (
	"fmt"
	"log"

	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

func DrawEditor(screen tcell.Screen, editorState *state.EditorState) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	lineNumStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorLightGray)

	// clear screen
	screen.Fill(' ', defStyle)
	log.Printf("Buffer has %d lines", len(editorState.Buffer().Data()))

	lineCount := len(editorState.Buffer().Data())
	gutterWidth := len(fmt.Sprintf("%d", lineCount)) + 1

	// draw text
	for row, rowRunes := range editorState.Buffer().Data() {
		log.Printf("Line %d has %d runes: %q", row, len(rowRunes), string(rowRunes))

		lineNum := fmt.Sprintf("%*d ", gutterWidth-1, row+1)
		for col, r := range lineNum {
			screen.SetContent(col, row, r, nil, lineNumStyle)
		}

		for col, r := range rowRunes {
			screen.SetContent(col + gutterWidth, row, r, nil, defStyle)
		}
	}

	// draw cursor
	cLine, cCol := editorState.Buffer().CursorPos()
	screen.ShowCursor(cCol + gutterWidth, cLine)
}
