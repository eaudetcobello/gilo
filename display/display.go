package display

import (
	"fmt"

	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

func DrawEditor(screen tcell.Screen, editorState *state.EditorState) {
	data := editorState.Buffer().Data()

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	lineNumStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorLightGray)

	// clear screen
	screen.Fill(' ', defStyle)

	gutterWidth := len(fmt.Sprintf("%d", len(data))) + 1

	// draw text
	topLine := editorState.TopLine()
	visibleLines := data[topLine:min(topLine+editorState.ScreenHeight(), len(data))]
	for row, rowRunes := range visibleLines {
		lineNum := fmt.Sprintf("%*d ", gutterWidth-1, row+topLine+1)
		for col, r := range lineNum {
			screen.SetContent(col, row, r, nil, lineNumStyle)
		}

		for col, r := range rowRunes {
			screen.SetContent(col+gutterWidth, row, r, nil, defStyle)
		}
	}

	// draw cursor
	cLine, cCol := editorState.Buffer().CursorPos()
	screen.ShowCursor(cCol+gutterWidth, cLine-topLine)
}
