package nytapi_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thorstenpfister/gonyt/internal/nytapi/apierror"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
	"github.com/thorstenpfister/gonyt/nytapi"
)

func Test_Client_ShouldHandleValid_FetchRequest_WithValues(t *testing.T) {
}

func Test_Client_ShouldHandleInvalid_FetchRequest_WithError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	// this is shit, recheck the public visibility
	section := nytapi.TopStoriesSection("arts")
	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(t, articles)
	require.Nil(t, updateTime)
	if assert.NotNil(t, err) {
		assert.IsType(t, apierror.APIError{}, err)
	}
}

func Test_Client_ShouldHandleInvalidTopStoriesSection_FetchRequest_withError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	// this is shit, recheck the public visibility
	section := nytapi.TopStoriesSection("This is not a valid section")
	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(t, articles)
	require.Nil(t, updateTime)
	assert.NotNil(t, err)
}
