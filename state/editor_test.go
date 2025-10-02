package state_test

import (
	"fmt"
	"testing"

	"github.com/eaudetcobello/gilo/state"
	"github.com/stretchr/testify/assert"
)

func TestEditorState(t *testing.T) {
	t.Parallel()

	t.Run("initial state", func(t *testing.T) {
		t.Parallel()

		got := state.NewEditorState(1, 1)
		assert.NotNil(t, got)
	})

	t.Run("quit flag", func(t *testing.T) {
		t.Parallel()

		state := state.NewEditorState(1, 1)
		assert.False(t, state.QuitFlag())
		state.Quit()
		assert.True(t, state.QuitFlag())
	})
}

func TestTopLine(t *testing.T) {
	t.Run("initial topline", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		lines := make([]string, 100)
		for i := range lines {
			lines[i] = fmt.Sprintf("Lorem ipsum, %d", i)
		}
		es.Buffer().SetData(lines)

		assert.Equal(t, 0, es.TopLine())
	})

	t.Run("move cursor down", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		lines := make([]string, 100)
		for i := range lines {
			lines[i] = fmt.Sprintf("Lorem ipsum, %d", i)
		}
		es.Buffer().SetData(lines)

		es.Buffer().SetCursorPos(23, 0)
		es.MoveCursorDown()
		assert.Equal(t, 1, es.TopLine())

		es.MoveCursorDown()
		assert.Equal(t, 2, es.TopLine())
	})

	t.Run("move cursor up", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 9)

		lines := make([]string, 15)
		for i := range lines {
			lines[i] = fmt.Sprintf("Lorem ipsum, %d", i)
		}

		es.Buffer().SetData(lines)

		// scroll 1 line past screen height, top line becomes 1
		es.Buffer().SetCursorPos(es.ScreenHeight() - 1, 0)
		es.MoveCursorDown()
		assert.Equal(t, 1, es.TopLine())

		// go down 2 lines past screen height, top line becomes 2
		es.MoveCursorDown()
		assert.Equal(t, 2, es.TopLine())

		// go up 9 times, from line 10
		for range es.ScreenHeight() {
			es.MoveCursorUp()
		}

		// top line is now 1 (10 - 9 = 1)
		assert.Equal(t, 1, es.TopLine())
	})
}
