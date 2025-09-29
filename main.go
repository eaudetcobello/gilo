package main

import (
	"fmt"
	"os"

	"github.com/eaudetcobello/gilo/app"
	"github.com/gdamore/tcell/v2"
)

func main() {
	if err := initEditor(); err != nil {
		exitWithError(err)
	}
}

func initEditor() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := screen.Init(); err != nil {
		return err
	}
	defer screen.Fini()

	// Screen setup
	screen.EnablePaste()
	screen.SetStyle(
		tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
	)

	editor := app.NewEditor(screen)
	editor.RunEventLoop()

	return nil
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
