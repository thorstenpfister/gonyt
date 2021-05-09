package nytapi

import (
	"context"
	"time"

	"github.com/thorstenpfister/gonyt/internal/nytapi"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
	"github.com/thorstenpfister/gonyt/internal/nytapi/query"
)

// Article as delivered by the New York Times API.
type Article = nytapi.Article

// Client for querying the New York Times API.
type Client struct {
	port port.HTTPPort
}

// NewClient provides a client for querying the New York Times API, providing your own HTTP client and API key.
func NewClient(httpClient port.HTTPClient, apiKey string) Client {
	client := Client{
		port: port.HTTPPort{
			HTTPClient: httpClient,
			BaseURL:    "https://api.nytimes.com/svc",
			APIKey:     apiKey,
		},
	}
	return client
}

// FetchTopStories is used to fetch the 'Top stories' from the New York Times API.
func (c *Client) FetchTopStories(ctx context.Context, section TopStoriesSection) (*[]Article, *time.Time, error) {
	if err := section.IsValid(); err != nil {
		return nil, nil, err
	}

	fetchTopStories := query.FetchTopStories{
		Section: string(section),
	}
	handler := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  c.port,
	}

	articles, lastUpdated, err := handler.Handle(ctx)
	if err != nil {
		return nil, nil, err
	}

	return articles, lastUpdated, nil
}
