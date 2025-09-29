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
				}
			}
		}

		if e.state.QuitFlag() {
			log.Printf("Quit flag set, exiting event loop...\n")

			return
		}
	}
}
