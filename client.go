package expogo

import "net/http"

type ExpoClient struct {
	host        string
	pushPath    string
	accessToken string
	httpClient  *http.Client
}

type ExpoConfig struct {
	Host        string
	PushPath    string
	AccessToken string
	HttpClient  *http.Client
}

const (
	DefaultExpoHost     = "https://exp.host/"
	DefaultExpoPushPath = "--/api/v2/push/send"
)

var DefaultHttpClient = &http.Client{}

func NewExpoClient(config *ExpoConfig) *ExpoClient {
	// Set default values
	expoClient := &ExpoClient{
		host:       DefaultExpoHost,
		pushPath:   DefaultExpoPushPath,
		httpClient: DefaultHttpClient,
	}

	// Override default values
	if config != nil {
		if config.Host != "" {
			expoClient.host = config.Host
		}
		if config.PushPath != "" {
			expoClient.pushPath = config.PushPath
		}
		if config.AccessToken != "" {
			expoClient.accessToken = config.AccessToken
		}
		if config.HttpClient != nil {
			expoClient.httpClient = config.HttpClient
		}
	}

	return expoClient
}
