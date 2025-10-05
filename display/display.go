package display

import (
	"fmt"

	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

var (
	DefaultStyle   = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	LineNumStyle   = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray)
	StatusBarStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
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

func drawStatusLine(screen tcell.Screen, lineNum, width, cx, cy, gutterWidth int, fileName string, style tcell.Style) {
	text := fmt.Sprintf("%s (%d, %d)", fileName, cy, cx)
	for x := range gutterWidth {
		screen.SetContent(x, lineNum, ' ', nil, style)
	}

	for x, r := range text {
		screen.SetContent(x+gutterWidth, lineNum, r, nil, style)
	}

	for x := len(text); x < width+gutterWidth; x++ {
		screen.SetContent(x+gutterWidth, lineNum, ' ', nil, style)
	}
}

func DrawEditor(screen tcell.Screen, es *state.EditorState) {
	screen.Fill(' ', DefaultStyle)

	gutterWidth := es.GutterWidth()

	for row, rowRunes := range es.VisibleLines() {
		drawLineNumber(screen, row, es.TopLine()+row+1, gutterWidth, LineNumStyle)
		drawTextLine(screen, row, rowRunes, gutterWidth, DefaultStyle)
	}

	cx, cy := es.CursorScreenPos()
	screen.ShowCursor(cx, cy)

	row, col := es.Buffer().CursorPos()
	drawStatusLine(screen, es.ScreenHeight()-1, es.ScreenWidth(), col+1, row+1, gutterWidth, es.Filename(), StatusBarStyle)
}
