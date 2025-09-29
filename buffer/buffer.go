package buffer

type GapBuffer struct {
	data     []rune
	gapStart int
	gapEnd   int
}

func NewGapBuffer() *GapBuffer {
	return &GapBuffer{
		data:     nil,
		gapStart: 0,
		gapEnd:   0,
	}
}
