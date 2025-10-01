package state_test

import (
	"testing"

	"github.com/eaudetcobello/gilo/state"
	"github.com/stretchr/testify/assert"
)

func lineText(bs *state.BufferState, lineNum int) string {
	return string(bs.Data()[lineNum])
}

func assertCursorAt(t *testing.T, b *state.BufferState, line, col int) {
	t.Helper()

	cline, ccol := b.CursorPos()

	assert.Equal(t, line, cline)
	assert.Equal(t, col, ccol)
}

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

func TestInsertRune(t *testing.T) {
	t.Parallel()

	t.Run("inserts rune at cursor", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(1, 1)

		es.Buffer().SetData([]string{
			"abc",
		})

		line := lineText(es.Buffer(), 0)
		assert.Equal(t, "abc", line)
	})

	t.Run("advances cursor after insert", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().InsertRune('a')

		assertCursorAt(t, es.Buffer(), 0, 1)
	})
}

func TestInsertNewline(t *testing.T) {
	t.Parallel()

	t.Run("split line at cursor", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().SetData([]string{
			"abc",
		})

		es.Buffer().SetCursorPos(0, 2)

		es.Buffer().InsertNewline()
		es.Buffer().InsertRune('d')

		lines := es.Buffer().Data()

		assert.Len(t, lines, 2)
		assert.Equal(t, "ab", lineText(es.Buffer(), 0))
		assert.Equal(t, "dc", lineText(es.Buffer(), 1))
	})
}

func TestMoveCursorLeft(t *testing.T) {
	t.Parallel()

	t.Run("move cursor left at beginning of line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().MoveCursorLeft()
		assertCursorAt(t, es.Buffer(), 0, 0)
	})

	t.Run("move cursor left in line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().SetData([]string{
			"abcd",
		})

		es.Buffer().SetCursorPos(0, 2)
		es.Buffer().MoveCursorLeft()
		assertCursorAt(t, es.Buffer(), 0, 1)
	})
}

func TestMoveCursorRight(t *testing.T) {
	t.Parallel()

	t.Run("move cursor right empty line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().MoveCursorRight()
		assertCursorAt(t, es.Buffer(), 0, 0)
	})

	t.Run("move cursor right", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().SetData([]string{
			"abcd",
		})

		es.Buffer().SetCursorPos(0, 2)
		es.Buffer().MoveCursorRight()
		assertCursorAt(t, es.Buffer(), 0, 3)

		// assert can go further than right-most rune
		es.Buffer().MoveCursorRight()
		assertCursorAt(t, es.Buffer(), 0, 4)
	})
}

func TestMoveCursorUpDown(t *testing.T) {
	t.Parallel()

	t.Run("move cursor down when buffer has no data", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().MoveCursorDown()
		assertCursorAt(t, es.Buffer(), 0, 0)
	})

	t.Run("move cursor up when buffer has no data", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)

		es.Buffer().MoveCursorUp()
		assertCursorAt(t, es.Buffer(), 0, 0)
	})

	t.Run("move cursor up/down", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData(
			[]string{
				"Test",
				"Test",
				"Test",
			},
		)

		es.Buffer().MoveCursorDown()
		assertCursorAt(t, es.Buffer(), 1, 0)

		es.Buffer().MoveCursorDown()
		assertCursorAt(t, es.Buffer(), 2, 0)

		es.Buffer().MoveCursorUp()
		assertCursorAt(t, es.Buffer(), 1, 0)

		es.Buffer().MoveCursorUp()
		assertCursorAt(t, es.Buffer(), 0, 0)
	})

	t.Run("move down from longer line to shorter line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData(
			[]string{
				"Longer line",
				"Shorter",
			},
		)

		es.Buffer().SetCursorPos(0, len(es.Buffer().Data()[0]))
		es.Buffer().MoveCursorDown()
		assertCursorAt(t, es.Buffer(), 1, len(es.Buffer().Data()[1]))
	})

	t.Run("move up from longer line to shorter line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData(
			[]string{
				"Shorter",
				"Longer line",
			},
		)

		es.Buffer().SetCursorPos(1, len(es.Buffer().Data()[1]))
		es.Buffer().MoveCursorUp()
		assertCursorAt(t, es.Buffer(), 0, len(es.Buffer().Data()[0]))
	})

	t.Run("move up/down with blank line", func (t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData(
			[]string{
				"Line",
				"",
				"Another line",
			},
		)

		// up
		es.Buffer().SetCursorPos(2, len(es.Buffer().Data()[2])-1)
		es.Buffer().MoveCursorUp()
		assertCursorAt(t, es.Buffer(), 1, 0)

		// down
		es.Buffer().SetCursorPos(0, len(es.Buffer().Data()[0])-1)
		es.Buffer().MoveCursorDown()
		assertCursorAt(t, es.Buffer(), 1, 0)
	})
}

func TestBackspace(t *testing.T) {
	t.Parallel()

	t.Run("backspace empty line", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData([]string{})

		es.Buffer().Backspace()

		assertCursorAt(t, es.Buffer(), 0, 0)
		assert.Equal(t, es.Buffer().Data(), [][]rune{})
	})

	t.Run("backspace in line removes rune before cursor", func(t *testing.T) {
		t.Parallel()

		es := state.NewEditorState(80, 24)
		es.Buffer().SetData([]string{
			"Line text abc",
		})

		es.Buffer().SetCursorPos(0, len(es.Buffer().Data()[0])-1)

		es.Buffer().Backspace()

		assertCursorAt(t, es.Buffer(), 0, len(es.Buffer().Data()[0])-1)

		assert.Equal(t, "Line text ac", lineText(es.Buffer(), 0))

		es.Buffer().SetCursorPos(0, 3)

		es.Buffer().Backspace()

		assert.Equal(t, "Lie text ac", lineText(es.Buffer(), 0))
	})

	t.Run("backspace at start of line moves cursor to end of prev line", func(t *testing.T) {
		t.Parallel()
	})
}