package bot

import (
	"fmt"
	"github.com/savaki/jq"
	"io/ioutil"
	"net/http"
	"strings"
)

const OxfordBaseUrl = "https://od-api.oxforddictionaries.com/api/v1"

type OxfordConfiguration struct {
	AppID string
	AppKey string
}

type OxfordClient struct {
	config *OxfordConfiguration
}

func NewOxfordClient(config OxfordConfiguration) *OxfordClient {
	return &OxfordClient{
		config: &config,
	}
}

func (oxford OxfordClient) Meaning(word string) (string, error) {
	req, err := http.NewRequest("GET", OxfordBaseUrl + "/entries/en/" + word, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("app_id", oxford.config.AppID)
	req.Header.Add("app_key", oxford.config.AppKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	meaning := oxford.extractMeaning(body)

	return meaning, nil
}

func (oxford OxfordClient) extractMeaning(json []byte) string {
	op, _ := jq.Parse(".results.[0].lexicalEntries.[0].entries.[0].senses.[0].definitions.[0]")
	result, _ := op.Apply(json)
	meaning := strings.Replace(string(result), "\"", "", -1)
	return meaning
}

func (oxford OxfordClient) Synonym(word string) (string, error) {
	req, err := http.NewRequest("GET", OxfordBaseUrl + "/entries/en/" + word + "/synonyms", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("app_id", oxford.config.AppID)
	req.Header.Add("app_key", oxford.config.AppKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	synonyms := oxford.extractSynonym(body)

	return concatSynonyms(synonyms), nil
}

func (oxford OxfordClient) extractSynonym(json []byte) []string {
	op, _ := jq.Parse(".results.[0].lexicalEntries.[0].entries.[0].senses.[0].subsenses.[0].synonyms")
	tmp, _ := op.Apply(json)

	var result []string
	for i := 0; i < 5; i++ {
		op, _ = jq.Parse(fmt.Sprintf(".[%d].text", i))
		word, err := op.Apply(tmp)
		if err != nil {
			break
		}

		synonym := strings.Replace(string(word), "\"", "", -1)
		result = append(result, string(synonym))
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func concatSynonyms(synonyms []string) string {
	if len(synonyms) == 0 {
		return ""
	}

	if len(synonyms) == 1 {
		return synonyms[0]
	}

	result := strings.Join(synonyms[0:len(synonyms)-1], ", ")
	result = result + " and " + synonyms[len(synonyms)-1]
	return result
}