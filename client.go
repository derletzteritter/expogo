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
	DEFAULT_EXPO_HOST      = "https://exp.host/"
	DEFAULT_EXPO_PUSH_PATH = "--/api/v2/push/send"
)

var DEFAULT_HTTP_CLIENT = &http.Client{}

func NewExpoClient(config *ExpoConfig) *ExpoClient {
	// Set default values
	expoClient := &ExpoClient{
		host:       DEFAULT_EXPO_HOST,
		pushPath:   DEFAULT_EXPO_PUSH_PATH,
		httpClient: DEFAULT_HTTP_CLIENT,
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
