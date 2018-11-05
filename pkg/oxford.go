package bot

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

func (oxford OxfordClient) Meaning(word string) string {
	return ""
}

func (oxford OxfordClient) Synonym(word string) string {
	return ""
}