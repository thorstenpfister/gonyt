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

type FetchTopStories struct {
	Section nytapi.TopStoriesSection
}

type FetchTopStoriesHandler struct {
	Query FetchTopStories
	Port  port.HTTPPort
}

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
		return nil, fmt.Errorf("Failed to create HTTP request for GET FetchTopStoriesRequest with error: %v", err)
	}

	return req, nil
}

type FetchTopStoriesAPIResponse struct {
	Status      string           `json:"status,omitempty"`
	Copyright   string           `json:"copyright,omitempty"`
	Section     string           `json:"section,omitempty"`
	LastUpdated time.Time        `json:"last_updated,omitempty"`
	NumResults  int              `json:"num_results,omitempty"`
	Results     []nytapi.Article `json:"results,omitempty"`
}

func newFetchTopStoriesAPIResponse(res *http.Response) (*FetchTopStoriesAPIResponse, error) {
	if res.Body == nil {
		return nil, fmt.Errorf("No FetchTopStoriesAPIResponse given.")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read body of response with error: %v", err)
	}

	var response = new(FetchTopStoriesAPIResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Unmarshaling the request body json into an FetchAccountResponse failed with error: %v", err)
	}

	return response, nil
}
