package apierror_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorstenpfister/gonyt/internal/nytapi/apierror"
)

func Test_APIError_DefaultsHTTPResponseNil_WithValue(t *testing.T) {
	sut := apierror.NewAPIError(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, sut.Type, apierror.UnknownError)
}

func Test_APIError_HandlesAPIErrorReponseNil_WithValue(t *testing.T) {
	sut := apierror.NewAPIError(&http.Response{StatusCode: 400})

	assert.NotNil(t, sut)
	assert.Equal(t, sut.Type, apierror.BadRequestError)
}

func Test_APIError_ShouldBeReflectingStatusCode_WithValue(t *testing.T) {
	var cases = []struct {
		inputResponse        http.Response
		expectedAPIErrorType int
	}{
		{http.Response{StatusCode: 1}, int(apierror.UnknownError)},
		{http.Response{StatusCode: 400}, int(apierror.BadRequestError)},
		{http.Response{StatusCode: 401}, int(apierror.UnauthorizedError)},
		{http.Response{StatusCode: 403}, int(apierror.ForbiddenError)},
		{http.Response{StatusCode: 404}, int(apierror.NotFoundError)},
		{http.Response{StatusCode: 405}, int(apierror.MethodNotAllowedError)},
		{http.Response{StatusCode: 406}, int(apierror.NotAcceptableError)},
		{http.Response{StatusCode: 408}, int(apierror.RequestTimedOutError)},
		{http.Response{StatusCode: 409}, int(apierror.ConflictError)},
		{http.Response{StatusCode: 410}, int(apierror.GoneError)},
		{http.Response{StatusCode: 411}, int(apierror.LengthRequiredError)},
		{http.Response{StatusCode: 413}, int(apierror.PayloadTooLargeError)},
		{http.Response{StatusCode: 414}, int(apierror.URITooLongError)},
		{http.Response{StatusCode: 415}, int(apierror.UnsupportedMediaTypeError)},
		{http.Response{StatusCode: 417}, int(apierror.ExpectationFailedError)},
		{http.Response{StatusCode: 429}, int(apierror.TooManyRequestsError)},
		{http.Response{StatusCode: 500}, int(apierror.ServerError)},
		{http.Response{StatusCode: 502}, int(apierror.BadGatewayError)},
		{http.Response{StatusCode: 503}, int(apierror.ServiceUnavailableError)},
		{http.Response{StatusCode: 504}, int(apierror.GatewayTimeoutError)},
	}

	for _, tt := range cases {
		sut := apierror.NewAPIError(&tt.inputResponse)

		assert.Equal(t, int(sut.Type), tt.expectedAPIErrorType)
	}
}

func Test_APIError_ShouldHaveErrorMessage(t *testing.T) {
	var cases = []struct {
		inputResponse http.Response
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
		sut := apierror.NewAPIError(&tt.inputResponse)

		if assert.Implements(t, (*error)(nil), sut) {
			assert.NotEmpty(t, sut.Error())
		}
	}
}
