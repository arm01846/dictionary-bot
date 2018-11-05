package bot

import (
	"io/ioutil"
	"net/http"
)

const OxfordBaseUrl = "https://od-api.oxforddictionaries.com/api/v1"

type OxfordConfiguration struct {
	AppID string
	AppKey string
}

type OxfordClient struct {
	config *OxfordConfiguration
}

func NewOxfordClient(config *OxfordConfiguration) *OxfordClient {
	return &OxfordClient{
		config: config,
	}
}

func (oxford OxfordClient) Meaning(word string) (string, error) {
	resp, err := http.Get(OxfordBaseUrl + "/entries/en/" + word)
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
	return ""
}

func (oxford OxfordClient) Synonym(word string) ([]string, error) {
	resp, err := http.Get(OxfordBaseUrl + "/entries/en/" + word + "/synonym")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	synonym := oxford.extractSynonym(body)

	return synonym, nil
}

func (oxford OxfordClient) extractSynonym(json []byte) []string {
	return nil
}