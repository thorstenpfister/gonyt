package port_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thorstenpfister/gonyt/internal/nytapi/apierror"
	"github.com/thorstenpfister/gonyt/internal/nytapi/port"
)

func Test_HTTPPort_HandlesSuccessfulResponseCodes_WithValue(t *testing.T) {
	var cases = []struct {
		mockedResponse http.Response
	}{
		{http.Response{StatusCode: 200}},
		{http.Response{StatusCode: 201}},
		{http.Response{StatusCode: 204}},
	}

	for _, tt := range cases {
		mockedHTTPClient := port.MockedHTTPClient{
			DoFunc: func(*http.Request) (*http.Response, error) {
				return &tt.mockedResponse, nil
			},
		}
		sut := port.HTTPPort{
			HTTPClient: &mockedHTTPClient,
			BaseURL:    "https://test.com",
		}

		req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
		res, err := sut.Do(req)

		require.Nil(t, err)
		assert.NotNil(t, res)
	}
}

func Test_HTTPPort_ErrorResponseCodesResultInAPIError_WithError(t *testing.T) {
	var cases = []struct {
		mockedResponse http.Response
	}{
		{http.Response{StatusCode: 1}},
		{http.Response{StatusCode: 400}},
		{http.Response{StatusCode: 401}},
		{http.Response{StatusCode: 403}},
		{http.Response{StatusCode: 404}},
		{http.Response{StatusCode: 405}},
		{http.Response{StatusCode: 406}},
		{http.Response{StatusCode: 408}},
		{http.Response{StatusCode: 409}},
		{http.Response{StatusCode: 410}},
		{http.Response{StatusCode: 411}},
		{http.Response{StatusCode: 413}},
		{http.Response{StatusCode: 414}},
		{http.Response{StatusCode: 415}},
		{http.Response{StatusCode: 417}},
		{http.Response{StatusCode: 429}},
		{http.Response{StatusCode: 500}},
		{http.Response{StatusCode: 502}},
		{http.Response{StatusCode: 503}},
		{http.Response{StatusCode: 504}},
	}

	for _, tt := range cases {
		mockedHTTPClient := port.MockedHTTPClient{
			DoFunc: func(*http.Request) (*http.Response, error) {
				return &tt.mockedResponse, nil
			},
		}
		sut := port.HTTPPort{
			HTTPClient: &mockedHTTPClient,
			BaseURL:    "https://test.com",
		}

		req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
		res, err := sut.Do(req)

		require.Nil(t, res)
		assert.NotNil(t, err)
		assert.IsType(t, apierror.APIError{}, err)
	}
}

func Test_HTTPPort_HandlesEmptyBody_WithValue(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200}, nil
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_HTTPPort_HandlesSuccessResponseWithRandomJSONBody_WithValue(t *testing.T) {
	json := `{"any":"value","will":"do"}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_HTTPPort_HandlesFailureResponseWithRandomJSONBody_WithError(t *testing.T) {
	json := `{"any":"value","will":"do"}`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 400, Body: body}, nil
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, res)
	assert.NotNil(t, err)
	assert.IsType(t, apierror.APIError{}, err)
}

func Test_HTTPPort_HTTPClientError_WithError(t *testing.T) {
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return nil, errors.New("test error")
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, res)
	assert.NotNil(t, err)
	assert.Implements(t, (*error)(nil), err)
}

func Test_HTTPPort_HandlesSuccessResponseWithInvalidJSONBody_WithNullValue(t *testing.T) {
	json := `not valid json`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: body}, nil
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, err)
	assert.NotNil(t, res)
}

func Test_HTTPPort_HandlesFailureResponseWithInvalidJSONBody_WithError(t *testing.T) {
	json := `not valid json`
	body := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mockedHTTPClient := port.MockedHTTPClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 400, Body: body}, nil
		},
	}
	sut := port.HTTPPort{
		HTTPClient: &mockedHTTPClient,
		BaseURL:    "https://test.com",
	}

	req := httptest.NewRequest(http.MethodGet, "https://test.com", nil)
	res, err := sut.Do(req)

	require.Nil(t, res)
	assert.NotNil(t, err)
	assert.Implements(t, (*error)(nil), err)
}
