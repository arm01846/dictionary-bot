package bot

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)


func Test_OxfordClient_extractMeaning(t *testing.T) {
	oxfordClient := NewOxfordClient(OxfordConfiguration{})

	t.Run("line", func(t *testing.T) {
		file, _ := os.Open("sample/line_meaning.json")
		input, _ := ioutil.ReadAll(file)
		assert.Equal(t, "a long, narrow mark or band", oxfordClient.extractMeaning(input))
	})
}

func Test_OxfordClient_extractSynonym(t *testing.T) {
	oxfordClient := NewOxfordClient(OxfordConfiguration{})

	t.Run("line", func(t *testing.T) {
		file, _ := os.Open("sample/line_synonym.json")
		input, _ := ioutil.ReadAll(file)
		assert.Equal(t, []string{
			"underline",
			"underscore",
			"stroke",
			"slash",
			"virgule",
		}, oxfordClient.extractSynonym(input))
	})

	t.Run("pen", func(t *testing.T) {
		file, _ := os.Open("sample/pen_synonym.json")
		input, _ := ioutil.ReadAll(file)
		assert.Nil(t, oxfordClient.extractSynonym(input))
	})
}

func Test_concatSynonyms(t *testing.T) {
	t.Run("no words, should return empty string", func(t *testing.T) {
		assert.Equal(t, "", concatSynonyms(nil))
		assert.Equal(t, "", concatSynonyms([]string{}))
	})

	t.Run("one word, should do nothing", func(t *testing.T) {
		assert.Equal(t, "word", concatSynonyms([]string{"word"}))
	})

	t.Run("more than one word, should join with ',' and 'and' for the last word", func(t *testing.T) {
		assert.Equal(t, "a and b", concatSynonyms([]string{"a", "b"}))
		assert.Equal(t, "a, b and c", concatSynonyms([]string{"a", "b", "c"}))
		assert.Equal(t, "a, b, c and d", concatSynonyms([]string{"a", "b", "c", "d"}))
		assert.Equal(t, "a, b, c, d and e", concatSynonyms([]string{"a", "b", "c", "d", "e"}))
	})

}