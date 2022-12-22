package chel

import (
	"net/http"

	"github.com/ryanminichiello/chel/config"
)

type Client struct {
	client        *http.Client
	TeamsService  *TeamService
	PeopleService *PeopleService
	APIVersion    string
}

// New returns a default NHL API Client
func New() (*Client, error) {
	defaultHTTPClient := http.DefaultClient
	apiConfig, err := config.InitConfig()
	if err != nil {
		return &Client{}, err
	}

	c := &Client{
		client: defaultHTTPClient,
		TeamsService: &TeamService{
			defaultHTTPClient,
			apiConfig,
		},
		PeopleService: &PeopleService{
			defaultHTTPClient,
			apiConfig,
		},
		APIVersion: "1.0",
	}
	return c, nil
}
