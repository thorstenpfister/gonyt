package nytapi_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/thorstenpfister/gonyt/nytapi"
)

var apiKey = "__PLEASE_PROVIDE_YOUR_OWN_KEY__"

type ClientE2ETestSuite struct {
	suite.Suite
	apiKey     string
	httpClient http.Client
}

func (suite *ClientE2ETestSuite) SetupTest() {
	if apiKey == "__PLEASE_PROVIDE_YOUR_OWN_KEY__" {
		panic(fmt.Sprintf("Please provide your own API key for e2e test runs. Currently it is: %v", apiKey))
	}

	suite.apiKey = apiKey
	suite.httpClient = http.Client{
		Timeout: time.Minute,
	}
}

func TestClientE2ETestSuite(t *testing.T) {
	suite.Run(t, new(ClientE2ETestSuite))
}

func (suite *ClientE2ETestSuite) Test_E2E_Client_ShouldHandleValid_FetchTopStories_WithValues() {
	sut := nytapi.NewClient(&suite.httpClient, suite.apiKey)

	ctx := context.Background()
	section := nytapi.Arts

	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(suite.T(), err)
	assert.NotNil(suite.T(), articles)
	assert.NotNil(suite.T(), updateTime)
}

func (suite *ClientE2ETestSuite) Test_E2E_Client_ShouldHandleValid_FetchBookReviews_WithValues() {
	var cases = []struct {
		category   nytapi.BookReviewsCategory
		searchTerm string
	}{
		{nytapi.Author, "Michelle Obama"},
		{nytapi.Isbn, "9781524763138"},
		{nytapi.Title, "Becoming"},
	}

	ctx := context.Background()
	sut := nytapi.NewClient(&suite.httpClient, suite.apiKey)

	for _, tt := range cases {
		bookReviews, err := sut.FetchBookReviews(ctx, tt.category, tt.searchTerm)

		require.Nil(suite.T(), err)
		assert.NotNil(suite.T(), bookReviews)
	}
}

func (suite *ClientE2ETestSuite) Test_E2E_Client_ShouldHandleValid_FetchMostPopularArticles_WithValues() {
	sut := nytapi.NewClient(&suite.httpClient, suite.apiKey)

	ctx := context.Background()
	category := nytapi.Emailed
	period := nytapi.Day

	articles, err := sut.FetchMostPopularArticles(ctx, category, period)

	require.Nil(suite.T(), err)
	assert.NotNil(suite.T(), articles)
}
