package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lineText(es *EditorState, lineNum int) string {
	return string(es.Buffer().Data()[lineNum])
}

func TestEditorState(t *testing.T) {
	t.Run("initial state", func(t *testing.T) {
		want := &EditorState{
			screenWidth:  1,
			screenHeight: 1,
			buffer: &BufferState{
				data: [][]rune{{}},
			},
		}
		got := NewEditorState(1, 1)
		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})

	t.Run("quit flag", func(t *testing.T) {
		state := NewEditorState(1, 1)
		assert.False(t, state.QuitFlag())
		state.Quit()
		assert.True(t, state.QuitFlag())
	})
}

func TestInsertRune(t *testing.T) {
	t.Run("inserts rune at cursor", func(t *testing.T) {
		es := NewEditorState(1, 1)

		es.InsertRune('a')
		es.InsertRune('b')
		es.InsertRune('c')

		line := lineText(es, 0)
		assert.Equal(t, "abc", line)
	})

	t.Run("advances cursor after insert", func(t *testing.T) {
		es := NewEditorState(80, 24)

		es.InsertRune('a')

		assert.Equal(t, 1, es.buffer.cursorCol)
	})
}

func TestInsertNewline(t *testing.T) {
	t.Run("split line at cursor", func(t *testing.T) {
		es := NewEditorState(80, 24)
		es.InsertRune('a')
		es.InsertRune('b')
		es.InsertRune('c')

		es.buffer.cursorCol = 2
		es.buffer.cursorLine = 0

		es.InsertNewline()
		es.InsertRune('d')

		lines := es.Buffer().Data()

		assert.Len(t, lines, 2)
		assert.Equal(t, "ab", lineText(es, 0))
		assert.Equal(t, "dc", lineText(es, 1))
	})
}
