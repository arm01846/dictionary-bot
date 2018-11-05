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