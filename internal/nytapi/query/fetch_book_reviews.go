package query

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/thorstenpfister/gonyt/internal/nytapi"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
)

// FetchTopStories models a query for fetching book reviews of a specified category of the New York Times API.
type FetchBookReviews struct {
	Category string
	Term     string
}

// FetchBookReviewsHandler is used to handle a FetchBookReviews query.
type FetchBookReviewsHandler struct {
	Query FetchBookReviews
	Port  port.HTTPPort
}

// Handle handles the query for book reviews from the New York Times API.
func (h *FetchBookReviewsHandler) Handle(ctx context.Context) (*[]nytapi.BookReview, error) {
	req, err := h.newFetchBookReviewsHTTPRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.Port.Do(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := newFetchBookReviewsAPIResponse(res)
	if err != nil {
		return nil, err
	}

	return &apiResponse.Results, nil
}

func (h *FetchBookReviewsHandler) newFetchBookReviewsHTTPRequest(ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("%v/books/v3/reviews.json?%v=%v&api-key=%v", h.Port.BaseURL, h.Query.Category, url.QueryEscape(h.Query.Term), h.Port.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request for GET FetchBookReviewsRequest with error: %v", err)
	}

	return req, nil
}

type fetchBookReviewsAPIResponse struct {
	Status     string              `json:"status,omitempty"`
	Copyright  string              `json:"copyright,omitempty"`
	NumResults int                 `json:"num_results,omitempty"`
	Results    []nytapi.BookReview `json:"results,omitempty"`
}

func newFetchBookReviewsAPIResponse(res *http.Response) (*fetchBookReviewsAPIResponse, error) {
	if res.Body == nil {
		return nil, fmt.Errorf("no FetchBookReviewsAPIResponse given")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of response with error: %v", err)
	}

	var response = new(fetchBookReviewsAPIResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling the request body json into an fetchBookReviews failed with error: %v", err)
	}

	return response, nil
}
