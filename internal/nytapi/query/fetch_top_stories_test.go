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

func Test_FetchTopStoriesHandler_HandlesSuccessResponseWithBody_WithValue(t *testing.T) {
	json := `{
				"status": "OK",
				"copyright": "Copyright (c) 2021 The New York Times Company. All Rights Reserved.",
				"section": "Arts",
				"last_updated": "2021-04-17T12:29:15-04:00",
				"num_results": 1,
				"results": [
					{
					"section": "theater",
					"subsection": "",
					"title": "After Bullying Reports, Scott Rudin Will Step Away From Broadway",
					"abstract": "The powerful producer of “Hello, Dolly!” and “The Book of Mormon” regrets “the pain my behavior caused” and says others will directly run his shows.",
					"url": "https://www.nytimes.com/2021/04/17/theater/scott-rudin-steps-away-from-broadway.html",
					"uri": "nyt://article/4036fdba-5b79-5bb6-af12-7315e4f38501",
					"byline": "By Michael Paulson",
					"item_type": "Article",
					"updated_date": "2021-04-17T13:56:26-04:00",
					"created_date": "2021-04-17T11:23:12-04:00",
					"published_date": "2021-04-17T11:23:12-04:00",
					"material_type_facet": "",
					"kicker": "",
					"des_facet": [
						"Theater",
						"Workplace Hazards and Violations"
					],
					"org_facet": [],
					"per_facet": [
						"Rudin, Scott"
					],
					"geo_facet": [],
					"multimedia": [
						{
						"url": "https://static01.nyt.com/images/2021/04/17/arts/17rudin/merlin_186176976_ccdce95b-3694-4427-84f4-a52b3f5109ce-superJumbo.jpg",
						"format": "superJumbo",
						"height": 2048,
						"width": 1584,
						"type": "image",
						"subtype": "photo",
						"caption": "Scott Rudin accepting the 2015 Tony Award for best revival of a play as a lead producer of  “Skylight.” ",
						"copyright": "Charles Sykes/Invision, via Charles Sykes/Invision/Ap"
						},
						{
						"url": "https://static01.nyt.com/images/2021/04/17/arts/17rudin/17rudin-thumbStandard.jpg",
						"format": "Standard Thumbnail",
						"height": 75,
						"width": 75,
						"type": "image",
						"subtype": "photo",
						"caption": "Scott Rudin accepting the 2015 Tony Award for best revival of a play as a lead producer of  “Skylight.” ",
						"copyright": "Charles Sykes/Invision, via Charles Sykes/Invision/Ap"
						}
					],
					"short_url": "https://nyti.ms/3dw0D1v"
					}
				]
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchTopStories := query.FetchTopStories{
		Section: "arts",
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
	sut := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, lastUpdated, err := sut.Handle(ctx)

	require.Nil(t, err)
	assert.NotNil(t, articles)
	assert.NotNil(t, lastUpdated)
}

func Test_FetchTopStoriesHandler_HandlesSuccessResponseWithInvalidBody_WithError(t *testing.T) {
	json := `totally not valid`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchTopStories := query.FetchTopStories{
		Section: "arts",
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
	sut := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, lastUpdated, err := sut.Handle(ctx)

	require.Nil(t, articles)
	require.Nil(t, lastUpdated)
	assert.NotNil(t, err)
}
func Test_FetchTopStoriesHandler_HandlesSuccessResponseWithoutBody_WithError(t *testing.T) {
	fetchTopStories := query.FetchTopStories{
		Section: "arts",
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
	sut := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, lastUpdated, err := sut.Handle(ctx)

	require.Nil(t, articles)
	require.Nil(t, lastUpdated)
	assert.NotNil(t, err)
}

func Test_FetchTopStoriesHandler_HandlesFailureResponse_WithError(t *testing.T) {
	fetchTopStories := query.FetchTopStories{
		Section: "arts",
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
	sut := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, lastUpdated, err := sut.Handle(ctx)

	require.Nil(t, articles)
	require.Nil(t, lastUpdated)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}

func Test_FetchTopStoriesHandler_HandlesFailureResponseBody_WithError(t *testing.T) {
	json := `{
				"unexpected": "message"
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchTopStories := query.FetchTopStories{
		Section: "arts",
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
	sut := query.FetchTopStoriesHandler{
		Query: fetchTopStories,
		Port:  mockedPort,
	}
	ctx := context.Background()

	articles, lastUpdated, err := sut.Handle(ctx)

	require.Nil(t, articles)
	require.Nil(t, lastUpdated)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}
