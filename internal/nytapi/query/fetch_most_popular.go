package query

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thorstenpfister/gonyt/internal/nytapi"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
)

// FetchMostPopular models a query for fetching 'most popular' articles of a specified section and time period of the New York Times API.
type FetchMostPopular struct {
	Category string
	Period   int
}

// FetchMostPopularHandler is used to handle a FetchMostPopular query.
type FetchMostPopularHandler struct {
	Query FetchMostPopular
	Port  port.HTTPPort
}

// Handle handles the query for a most popular category for a given time period from the New York Times API.
func (h *FetchMostPopularHandler) Handle(ctx context.Context) (*[]nytapi.PopularArticle, error) {
	req, err := h.newFetchMostPopularHTTPRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.Port.Do(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := newFetchMostPopularAPIResponse(res)
	if err != nil {
		return nil, err
	}

	return &apiResponse.Results, nil
}

func (h *FetchMostPopularHandler) newFetchMostPopularHTTPRequest(ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("%v/mostpopular/v2/%v/%v.json?api-key=%v", h.Port.BaseURL, h.Query.Category, h.Query.Period, h.Port.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request for GET FetchMostPopularRequest with error: %v", err)
	}

	return req, nil
}

type fetchMostPopularAPIResponse struct {
	Status     string                  `json:"status,omitempty"`
	Copyright  string                  `json:"copyright,omitempty"`
	Section    string                  `json:"section,omitempty"`
	NumResults int                     `json:"num_results,omitempty"`
	Results    []nytapi.PopularArticle `json:"results,omitempty"`
}

func newFetchMostPopularAPIResponse(res *http.Response) (*fetchMostPopularAPIResponse, error) {
	if res.Body == nil {
		return nil, fmt.Errorf("no FetchMostPopularAPIResponse given")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of response with error: %v", err)
	}

	var response = new(fetchMostPopularAPIResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling the request body json into an FetchMostPopularResponse failed with error: %v", err)
	}

	return response, nil
}
