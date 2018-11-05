package bot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectFirstWord(t *testing.T) {
	t.Run("without words should return empty string", func(t * testing.T) {
		assert.Equal(t, "", detectFirstWord(""))
	})

	t.Run("with single word should return that word back", func(t *testing.T) {
		assert.Equal(t, "word", detectFirstWord("word"))
	})

	t.Run("with more than one word should return the first word", func(t *testing.T) {
		assert.Equal(t, "first", detectFirstWord("first second"))
	})
}