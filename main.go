package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eaudetcobello/gilo/app"
	"github.com/gdamore/tcell/v2"
)

func main() {
	logFile, err := os.OpenFile("editor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Editor starting...")

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
