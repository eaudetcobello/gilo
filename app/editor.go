package app

import (
	"log"

	"github.com/eaudetcobello/gilo/display"
	"github.com/eaudetcobello/gilo/state"
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	screen        tcell.Screen
	termEventChan chan tcell.Event
	quitChan      chan struct{}
	state         *state.EditorState
	filepath string
}

func NewEditor(screen tcell.Screen) *Editor {
	screenWidth, screenHeight := screen.Size()
	termEventChan := make(chan tcell.Event, 1)
	quitChan := make(chan struct{}, 1)

	editorState := state.NewEditorState(
		screenWidth,
		screenHeight,
	)

	editor := &Editor{
		screen:        screen,
		termEventChan: termEventChan,
		quitChan:      quitChan,
		state:         editorState,
	}

	return editor
}

func (e *Editor) RunEventLoop() {
	e.redraw(true)
	go e.screen.ChannelEvents(e.termEventChan, e.quitChan)
	e.runMainEventLoop()
}

func (e *Editor) redraw(sync bool) {
	display.DrawEditor(e.screen, e.state)

	if sync {
		e.screen.Sync()
	} else {
		e.screen.Show()
	}
}

func (e *Editor) runMainEventLoop() {
	for {
		select {
		case event := <-e.termEventChan:
			switch ev := event.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					e.state.Quit()

					break
				} else if ev.Key() == tcell.KeyEnter {
					log.Printf("got keyenter")
					e.state.Buffer().InsertNewline()
				} else if ev.Key() == tcell.KeyLeft {
					log.Printf("got keyleft")
					e.state.Buffer().MoveCursorLeft()
				} else if ev.Key() == tcell.KeyRight {
					log.Printf("got keyright")
					e.state.Buffer().MoveCursorRight()
				} else if ev.Key() == tcell.KeyDown {
					log.Printf("got keydown")
					e.state.MoveCursorDown()
				} else if ev.Key() == tcell.KeyUp {
					log.Printf("got keyup")
					e.state.MoveCursorUp()
				} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
					log.Printf("got keybs")
					e.state.Buffer().Backspace()
				} else if ev.Key() == tcell.KeyRune {
					log.Printf("got rune")
					e.state.Buffer().InsertRune(ev.Rune())
				}
			}

			cline, ccol := e.state.Buffer().CursorPos()
			log.Printf("Cursor at (%d,%d)", cline, ccol)
			log.Printf("Buffer has %d lines", len(e.state.Buffer().Data()))

			e.redraw(false)
		}

		if e.state.QuitFlag() {
			log.Printf("Quit flag set, exiting event loop...\n")

			return
		}
	}
}

func (e *Editor) LoadFile(path string) error {
	if err := e.state.Buffer().LoadFromFile(path); err != nil {
		return err
	}
	e.filepath = path
	return nil
}

// func (e *Editor) SaveFile() error {
// 	if e.filepath == "" {
// 		return fmt.Errorf("no filename")
// 	}
// 	return e.state.Buffer().SaveToFile(e.filepath)
// }
