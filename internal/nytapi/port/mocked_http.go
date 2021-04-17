package port

import "net/http"

// MockedHTTPClient provides an easy abstracted http.Client mock.
// The provided DoFunc() member can be used to inject custom behaviour during tests.
type MockedHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do performs as a mocked drop in replacement and satisfies the HTTPClient interface.
// It will execute the member DoFunc().
func (m *MockedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
