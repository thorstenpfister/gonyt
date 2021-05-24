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

func Test_FetchBookReviewsHandler_HandlesSuccessResponseWithBody_WithValue(t *testing.T) {
	json := `{
				"status": "OK",
				"copyright": "Copyright (c) 2021 The New York Times Company.  All Rights Reserved.",
				"num_results": 1,
				"results": [
					{
					"url": "https:\/\/www.nytimes.com\/2018\/12\/06\/books\/review\/michelle-obama-becoming-memoir.html",
					"publication_dt": "2018-12-06",
					"byline": "Isabel Wilkerson",
					"book_title": "Becoming",
					"book_author": "Michelle Obama",
					"summary": "The former first lady\u2019s long-awaited new memoir recounts with insight, candor and wit her family\u2019s trajectory from the Jim Crow South to Chicago\u2019s South Side and her own improbable journey from there to the White House.",
					"uuid": "00000000-0000-0000-0000-000000000000",
					"uri": "nyt:\/\/book\/00000000-0000-0000-0000-000000000000",
					"isbn13": [
						"9781524763138"
					]
					}
				]
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchBookReviews := query.FetchBookReviews{
		Category: "article",
		Term:     "Michelle Obama",
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
	sut := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  mockedPort,
	}
	ctx := context.Background()

	bookReviews, err := sut.Handle(ctx)

	require.Nil(t, err)
	assert.NotNil(t, bookReviews)
}

func Test_FetchBookReviewsHandler_HandlesSuccessResponseWithInvalidBody_WithError(t *testing.T) {
	json := `totally not valid`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchBookReviews := query.FetchBookReviews{
		Category: "article",
		Term:     "Michelle Obama",
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
	sut := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  mockedPort,
	}
	ctx := context.Background()

	bookReviews, err := sut.Handle(ctx)

	require.Nil(t, bookReviews)
	assert.NotNil(t, err)
}
func Test_FetchBookReviewsHandler_HandlesSuccessResponseWithoutBody_WithError(t *testing.T) {
	fetchBookReviews := query.FetchBookReviews{
		Category: "article",
		Term:     "Michelle Obama",
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
	sut := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  mockedPort,
	}
	ctx := context.Background()

	bookReviews, err := sut.Handle(ctx)

	require.Nil(t, bookReviews)
	assert.NotNil(t, err)
}

func Test_FetchBookReviewsHandler_HandlesFailureResponse_WithError(t *testing.T) {
	fetchBookReviews := query.FetchBookReviews{
		Category: "article",
		Term:     "Michelle Obama",
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
	sut := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  mockedPort,
	}
	ctx := context.Background()

	bookReviews, err := sut.Handle(ctx)

	require.Nil(t, bookReviews)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}

func Test_FetchBookReviewsHandler_HandlesFailureResponseBody_WithError(t *testing.T) {
	json := `{
				"unexpected": "message"
			}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	fetchBookReviews := query.FetchBookReviews{
		Category: "article",
		Term:     "Michelle Obama",
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
	sut := query.FetchBookReviewsHandler{
		Query: fetchBookReviews,
		Port:  mockedPort,
	}
	ctx := context.Background()

	bookReviews, err := sut.Handle(ctx)

	require.Nil(t, bookReviews)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}
