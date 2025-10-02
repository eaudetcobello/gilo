package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eaudetcobello/gilo/app"
	"github.com/gdamore/tcell/v2"
)

func main() {
	logFile, err := os.OpenFile("editor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Editor starting...")

	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	if err := initEditor(filePath); err != nil {
		exitWithError(err)
	}
}

func initEditor(filepath string) error {
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
	sWidth, sHeight := screen.Size()
	log.Printf("Editor has dimensions (%d,%d)", sHeight, sWidth)

	// Load file if provided
	if filepath != "" {
		if err := editor.LoadFile(filepath); err != nil {
			return fmt.Errorf("failed to start editor: %w", err)
		}
	}

	editor.RunEventLoop()
	return nil
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
