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

// BookReview as delivered by the New York Times API.
type BookReview = nytapi.BookReview

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

// FetchBookReviews is used to fetch book reviews from the New York Times API.
func (c *Client) FetchBookReviews(ctx context.Context, category BookReviewsCategory, searchTerm string) (*[]BookReview, error) {
	if err := category.IsValid(); err != nil {
		return nil, err
	}

	fetchBookReviews := query.FetchBookReviews{
		Category: string(category),
		Term:     searchTerm,
	}
	handler := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  c.port,
	}

	bookReviews, err := handler.Handle(ctx)
	if err != nil {
		return nil, err
	}

	return bookReviews, nil
}
