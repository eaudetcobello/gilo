package display

import (
	"fmt"

	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

var (
	DefaultStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	LineNumStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorLightGray)
)

func drawLineNumber(screen tcell.Screen, screenRow, lineNum, gutterWidth int, style tcell.Style) {
	label := fmt.Sprintf("%*d ", gutterWidth-1, lineNum)
	for x, r := range label {
		screen.SetContent(x, screenRow, r, nil, style)
	}
}

func drawTextLine(screen tcell.Screen, lineNum int, runes []rune, gutterWidth int, style tcell.Style) {
	for x, r := range runes {
		screen.SetContent(x+gutterWidth, lineNum, r, nil, style)
	}
}

func DrawEditor(screen tcell.Screen, editorState *state.EditorState) {
	screen.Fill(' ', DefaultStyle)

	gutterWidth := editorState.GutterWidth()

	for row, rowRunes := range editorState.VisibleLines() {
		lineNum := editorState.TopLine() + row + 1
		drawLineNumber(screen, row, lineNum, gutterWidth, LineNumStyle)
		drawTextLine(screen, row, rowRunes, gutterWidth, DefaultStyle)
	}

	cx, cy := editorState.CursorScreenPos()
	screen.ShowCursor(cx, cy)
}
