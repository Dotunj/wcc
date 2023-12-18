package wc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	path = "../test.txt"
)

func Test_ByteCount(t *testing.T) {
	wc := &Wc{Path: path}

	byteCount, err := wc.ByteCount()
	assert.Equal(t, int64(342190), byteCount)
	assert.Nil(t, err)
}

func Test_LineCount(t *testing.T) {
	wc := &Wc{Path: path}

	lineCount, err := wc.LineCount()
	assert.Equal(t, 7145, lineCount)
	assert.Nil(t, err)
}

func Test_WordCount(t *testing.T) {
	wc := &Wc{Path: path}

	wordCount, err := wc.WordCount()
	assert.Equal(t, 58164, wordCount)
	assert.Nil(t, err)
}

func Test_CharacterCount(t *testing.T) {
	wc := &Wc{Path: path}

	characterCount, err := wc.CharacterCount()
	assert.Equal(t, 339292, characterCount)
	assert.Nil(t, err)
}
