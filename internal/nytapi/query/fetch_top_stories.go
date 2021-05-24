package query

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thorstenpfister/gonyt/internal/nytapi"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
)

// FetchTopStories models a query for fetching 'Top stories' of a specified section of the New York Times API.
type FetchTopStories struct {
	Section string
}

// FetchTopStoriesHandler is used to handle a FetchTopStories query.
type FetchTopStoriesHandler struct {
	Query FetchTopStories
	Port  port.HTTPPort
}

// Handle handles the query for a 'Top stories' section from the New York Times API.
func (h *FetchTopStoriesHandler) Handle(ctx context.Context) (*[]nytapi.Article, *time.Time, error) {
	req, err := h.newFetchTopStoriesHTTPRequest(ctx)
	if err != nil {
		return nil, nil, err
	}

	res, err := h.Port.Do(req)
	if err != nil {
		return nil, nil, err
	}

	apiResponse, err := newFetchTopStoriesAPIResponse(res)
	if err != nil {
		return nil, nil, err
	}

	return &apiResponse.Results, &apiResponse.LastUpdated, nil
}

func (h *FetchTopStoriesHandler) newFetchTopStoriesHTTPRequest(ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("%v/topstories/v2/%v.json?api-key=%v", h.Port.BaseURL, h.Query.Section, h.Port.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request for GET FetchTopStoriesRequest with error: %v", err)
	}

	return req, nil
}

type fetchTopStoriesAPIResponse struct {
	Status      string           `json:"status,omitempty"`
	Copyright   string           `json:"copyright,omitempty"`
	Section     string           `json:"section,omitempty"`
	LastUpdated time.Time        `json:"last_updated,omitempty"`
	NumResults  int              `json:"num_results,omitempty"`
	Results     []nytapi.Article `json:"results,omitempty"`
}

func newFetchTopStoriesAPIResponse(res *http.Response) (*fetchTopStoriesAPIResponse, error) {
	if res.Body == nil {
		return nil, fmt.Errorf("no FetchTopStoriesAPIResponse given")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of response with error: %v", err)
	}

	var response = new(fetchTopStoriesAPIResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling the request body json into an FetchAccountResponse failed with error: %v", err)
	}

	return response, nil
}
