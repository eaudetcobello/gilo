package state_test

import (
	"testing"

	"github.com/eaudetcobello/gilo/state"
	"github.com/stretchr/testify/assert"
)

func lineText(bs *state.BufferState, lineNum int) string {
	return string(bs.Data()[lineNum])
}

func TestEditorState(t *testing.T) {
	t.Run("initial state", func(t *testing.T) {
		got := state.NewEditorState(1, 1)
		assert.NotNil(t, got)
	})

	t.Run("quit flag", func(t *testing.T) {
		state := state.NewEditorState(1, 1)
		assert.False(t, state.QuitFlag())
		state.Quit()
		assert.True(t, state.QuitFlag())
	})
}

func TestInsertRune(t *testing.T) {
	t.Run("inserts rune at cursor", func(t *testing.T) {
		es := state.NewEditorState(1, 1)

		es.Buffer().InsertRune('a')
		es.Buffer().InsertRune('b')
		es.Buffer().InsertRune('c')

		line := lineText(es.Buffer(), 0)
		assert.Equal(t, "abc", line)
	})

	t.Run("advances cursor after insert", func(t *testing.T) {
		es := state.NewEditorState(80, 24)

		es.Buffer().InsertRune('a')

		cRow, cCol := es.Buffer().CursorPos()
		assert.Equal(t, 0, cRow)
		assert.Equal(t, 1, cCol)
	})
}

func TestInsertNewline(t *testing.T) {
	t.Run("split line at cursor", func(t *testing.T) {
		es := state.NewEditorState(80, 24)
		es.Buffer().InsertRune('a')
		es.Buffer().InsertRune('b')
		es.Buffer().InsertRune('c')

		es.Buffer().SetCursorPos(0, 2)

		es.Buffer().InsertNewline()
		es.Buffer().InsertRune('d')

		lines := es.Buffer().Data()

		assert.Len(t, lines, 2)
		assert.Equal(t, "ab", lineText(es.Buffer(), 0))
		assert.Equal(t, "dc", lineText(es.Buffer(), 1))
	})
}
