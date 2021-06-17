package query_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thorstenpfister/gonyt/internal/nytapi/apierror"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
	"github.com/thorstenpfister/gonyt/internal/nytapi/query"
)

func Test_FetchMostPopularHandler_HandlesSuccessResponseWithBody_WithValue(t *testing.T) {
	json := `{
				"status": "OK",
				"copyright": "Copyright (c) 2021 The New York Times Company.  All Rights Reserved.",
				"num_results": 20,
				"results": [
					{
					"uri": "nyt://article/ea6e6f6e-f40a-5be8-99b4-a4ffa79531d3",
					"url": "https://www.nytimes.com/2021/06/09/well/move/exercise-blood-test.html",
					"id": 100000007804694,
					"asset_id": 100000007804694,
					"source": "New York Times",
					"published_date": "2021-06-09",
					"updated": "2021-06-16 15:04:02",
					"section": "Well",
					"subsection": "Move",
					"nytdsection": "well",
					"adx_keywords": "Exercise;Weight;Longevity;Research;Blood;Nature Metabolism (Journal)",
					"column": null,
					"byline": "By Gretchen Reynolds",
					"type": "Article",
					"title": "The Best Type of Exercise? A Blood Test Holds Clues",
					"abstract": "Researchers are studying the proteins in blood to learn why some of us respond to certain forms of exercise better than others.",
					"des_facet": [
						"Exercise",
						"Weight",
						"Longevity",
						"Research",
						"Blood"
					],
					"org_facet": [
						"Nature Metabolism (Journal)"
					],
					"per_facet": [],
					"geo_facet": [],
					"media": [
						{
						"type": "image",
						"subtype": "photo",
						"caption": "",
						"copyright": "Neil Hall/EPA, via Shutterstock",
						"approved_for_syndication": 1,
						"media-metadata": [
							{
							"url": "https://static01.nyt.com/images/2021/06/15/well/physed-responder2/physed-responder2-thumbStandard.jpg",
							"format": "Standard Thumbnail",
							"height": 75,
							"width": 75
							},
							{
							"url": "https://static01.nyt.com/images/2021/06/15/well/physed-responder2/physed-responder2-mediumThreeByTwo210.jpg",
							"format": "mediumThreeByTwo210",
							"height": 140,
							"width": 210
							},
							{
							"url": "https://static01.nyt.com/images/2021/06/15/well/physed-responder2/physed-responder2-mediumThreeByTwo440.jpg",
							"format": "mediumThreeByTwo440",
							"height": 293,
							"width": 440
							}
						]
						}
					],
					"eta_id": 0
					}
				]
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchMostPopular := query.FetchMostPopular{
		Category: "emailed",
		Period:   1,
	}
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	mockedPort := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test-is-mocked.com",
		APIKey:     "1234567890",
	}
	sut := query.FetchMostPopularHandler{
		Query: fetchMostPopular,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, err := sut.Handle(ctx)

	require.Nil(t, err)
	assert.NotNil(t, articles)
}

func Test_FetchMostPopularHandler_HandlesSuccessResponseWithInvalidBody_WithError(t *testing.T) {
	json := `totally not valid`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchMostPopular := query.FetchMostPopular{
		Category: "emailed",
		Period:   1,
	}
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	mockedPort := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test-is-mocked.com",
		APIKey:     "1234567890",
	}
	sut := query.FetchMostPopularHandler{
		Query: fetchMostPopular,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, err := sut.Handle(ctx)

	require.Nil(t, articles)
	assert.NotNil(t, err)
}
func Test_FetchMostPopularHandler_HandlesSuccessResponseWithoutBody_WithError(t *testing.T) {
	fetchMostPopular := query.FetchMostPopular{
		Category: "emailed",
		Period:   1,
	}
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200}, nil
		},
	}
	mockedPort := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test-is-mocked.com",
		APIKey:     "1234567890",
	}
	sut := query.FetchMostPopularHandler{
		Query: fetchMostPopular,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, err := sut.Handle(ctx)

	require.Nil(t, articles)
	assert.NotNil(t, err)
}

func Test_FetchMostPopularHandler_HandlesFailureResponse_WithError(t *testing.T) {
	fetchMostPopular := query.FetchMostPopular{
		Category: "emailed",
		Period:   1,
	}
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 400}, nil
		},
	}
	mockedPort := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test-is-mocked.com",
		APIKey:     "1234567890",
	}
	sut := query.FetchMostPopularHandler{
		Query: fetchMostPopular,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, err := sut.Handle(ctx)

	require.Nil(t, articles)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}

func Test_FetchMostPopularHandler_HandlesFailureResponseBody_WithError(t *testing.T) {
	json := `{
				"unexpected": "message"
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchMostPopular := query.FetchMostPopular{
		Category: "emailed",
		Period:   1,
	}
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404, Body: body}, nil
		},
	}
	mockedPort := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test-is-mocked.com",
		APIKey:     "1234567890",
	}
	sut := query.FetchMostPopularHandler{
		Query: fetchMostPopular,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, err := sut.Handle(ctx)

	require.Nil(t, articles)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}
