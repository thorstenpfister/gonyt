package nytapi_test

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
	"github.com/thorstenpfister/gonyt/nytapi"
)

func Test_Client_ShouldHandleValid_FetchTopStories_WithValues(t *testing.T) {
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

	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	section := nytapi.Arts
	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(t, err)
	assert.NotNil(t, articles)
	assert.NotNil(t, updateTime)
}

func Test_Client_ShouldHandleInvalid_FetchTopStories_WithError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	section := nytapi.Arts
	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(t, articles)
	require.Nil(t, updateTime)
	if assert.NotNil(t, err) {
		assert.IsType(t, apierror.APIError{}, err)
	}
}

func Test_Client_ShouldHandleInvalidTopStoriesSection_FetchTopStories_withError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	section := nytapi.TopStoriesSection("This is not a valid section")
	articles, updateTime, err := sut.FetchTopStories(ctx, section)

	require.Nil(t, articles)
	require.Nil(t, updateTime)
	assert.NotNil(t, err)
}

func Test_Client_ShouldHandleValid_FetchBookReviews_WithValues(t *testing.T) {
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

	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.Author
	term := "Michelle Obama"
	bookReviews, err := sut.FetchBookReviews(ctx, category, term)

	require.Nil(t, err)
	assert.NotNil(t, bookReviews)
}

func Test_Client_ShouldHandleInvalid_FetchBookReviews_WithError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.Author
	term := "Michelle Obama"
	bookReviews, err := sut.FetchBookReviews(ctx, category, term)

	require.Nil(t, bookReviews)
	if assert.NotNil(t, err) {
		assert.IsType(t, apierror.APIError{}, err)
	}
}

func Test_Client_ShouldHandleInvalidBookReviewsCategory_FetchBookReviews_withError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.BookReviewsCategory("This is not a valid section")
	term := "Michelle Obama"
	bookReviews, err := sut.FetchBookReviews(ctx, category, term)

	require.Nil(t, bookReviews)
	assert.NotNil(t, err)
}

func Test_Client_ShouldHandleValid_FetchMostPopularArticles_WithValues(t *testing.T) {
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

	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.Emailed
	period := nytapi.Day
	articles, err := sut.FetchMostPopularArticles(ctx, category, period)

	require.Nil(t, err)
	assert.NotNil(t, articles)
}

func Test_Client_ShouldHandleInvalid_FetchMostPopularArticles_WithError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.Emailed
	period := nytapi.Day
	articles, err := sut.FetchMostPopularArticles(ctx, category, period)

	require.Nil(t, articles)
	if assert.NotNil(t, err) {
		assert.IsType(t, apierror.APIError{}, err)
	}
}

func Test_Client_ShouldHandleInvalidMostPopularCategory_FetchMostPopularArticles_withError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.MostPopularCategory("not-a-valid-category")
	period := nytapi.Day
	articles, err := sut.FetchMostPopularArticles(ctx, category, period)

	require.Nil(t, articles)
	assert.NotNil(t, err)
}

func Test_Client_ShouldHandleInvalidMostPopularPeriod_FetchMostPopularArticles_withError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 404}, nil
		},
	}
	apiKey := "mockedApiKey"
	sut := nytapi.NewClient(&mockedHTTPClient, apiKey)

	ctx := context.Background()
	category := nytapi.Emailed
	period := nytapi.MostPopularPeriod(999999)
	articles, err := sut.FetchMostPopularArticles(ctx, category, period)

	require.Nil(t, articles)
	assert.NotNil(t, err)
}
