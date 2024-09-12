package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/wundergraph/cosmo/connect-go/wg/cosmo/platform/v1/platformv1connect"
	"github.com/wundergraph/cosmo/terraform-provider-cosmo/internal/utils"
)

type PlatformClient struct {
	Client      platformv1connect.PlatformServiceClient
	CosmoApiKey string
}

func NewClient(apiKey, apiUrl string) (*PlatformClient, error) {
	cosmoApiKey, ok := os.LookupEnv(utils.EnvCosmoApiKey)
	if !ok {
		return nil, fmt.Errorf("COSMO_API_KEY environment variable not set")
	}

	cosmoApiUrl, ok := os.LookupEnv(utils.EnvCosmoApiUrl)
	if !ok {
		cosmoApiUrl = "https://cosmo-cp.wundergraph.com"
	}

	httpClient := http.Client{}
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	client := platformv1connect.NewPlatformServiceClient(&httpClient, cosmoApiUrl)

	return &PlatformClient{
		Client:      client,
		CosmoApiKey: cosmoApiKey,
	}, nil
}
