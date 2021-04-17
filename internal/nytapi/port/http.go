package port

import (
	"fmt"
	"net/http"

	"github.com/thorstenpfister/gonyt/internal/nytapi/apierror"
)

// HTTPClient represents an interface compatible with net/http.Do() to facility injection of mocks.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPPort represents a port holding a http client satisfying the HTTPClient interface,
// a BaseURL for all requests and a Host used in request headers.
type HTTPPort struct {
	HTTPClient HTTPClient
	BaseURL    string
	APIKey     string
}

// Do intiates the execution of a http.Request and results in a http.Response in case of success
// or in an error specifying the source of failure.
func (p *HTTPPort) Do(req *http.Request) (*http.Response, error) {
	res, err := p.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request execution failed with error: %v", err)
	}

	if !p.successfulRequest(res) {
		apiError := apierror.NewAPIError(res)
		return nil, apiError
	}

	return res, nil
}

func (p *HTTPPort) successfulRequest(res *http.Response) bool {
	if res.StatusCode == 200 || res.StatusCode == 201 || res.StatusCode == 204 {
		return true
	}
	return false
}
