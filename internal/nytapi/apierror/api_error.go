package apierror

import (
	"fmt"
	"net/http"
)

type errorType int

// Common HTTP errors.
//
// Unless otherwise noted, these are defined in RFC 7231 (https://tools.ietf.org/html/rfc7231)
const (
	UnknownError errorType = iota // Catch all API error for as of yet undefined cases.
	BadRequestError
	UnauthorizedError
	ForbiddenError
	NotFoundError
	MethodNotAllowedError
	NotAcceptableError
	RequestTimedOutError
	ConflictError
	GoneError
	LengthRequiredError
	PayloadTooLargeError
	URITooLongError
	UnsupportedMediaTypeError
	ExpectationFailedError
	TooManyRequestsError
	ServerError
	BadGatewayError
	ServiceUnavailableError
	GatewayTimeoutError
)

// APIError captures basic information for any API error.
type APIError struct {
	Type           errorType
	HTTPStatusCode int
}

// NewAPIError creates a new concrete apierror.APIError from a given *http.response.
func NewAPIError(res *http.Response) APIError {
	errorWithType := func(errType errorType) APIError {
		var statusCode = 0
		if res != nil {
			statusCode = res.StatusCode
		}

		return APIError{
			Type:           errType,
			HTTPStatusCode: statusCode,
		}
	}

	if res == nil {
		return errorWithType(UnknownError)
	}

	switch res.StatusCode {
	case 400:
		return errorWithType(BadRequestError)
	case 401:
		return errorWithType(UnauthorizedError)
	case 403:
		return errorWithType(ForbiddenError)
	case 404:
		return errorWithType(NotFoundError)
	case 405:
		return errorWithType(MethodNotAllowedError)
	case 406:
		return errorWithType(NotAcceptableError)
	case 408:
		return errorWithType(RequestTimedOutError)
	case 409:
		return errorWithType(ConflictError)
	case 410:
		return errorWithType(GoneError)
	case 411:
		return errorWithType(LengthRequiredError)
	case 413:
		return errorWithType(PayloadTooLargeError)
	case 414:
		return errorWithType(URITooLongError)
	case 415:
		return errorWithType(UnsupportedMediaTypeError)
	case 417:
		return errorWithType(ExpectationFailedError)
	case 429:
		return errorWithType(TooManyRequestsError)
	case 500:
		return errorWithType(ServerError)
	case 502:
		return errorWithType(BadGatewayError)
	case 503:
		return errorWithType(ServiceUnavailableError)
	case 504:
		return errorWithType(GatewayTimeoutError)
	}

	return errorWithType(UnknownError)
}

func (err APIError) Error() string {
	return fmt.Sprintf("%v %v", err.baseMessage(), err.baseInformation())
}

func (err APIError) baseInformation() string {
	return fmt.Sprintf("HTTP Status: %d", err.HTTPStatusCode)
}

func (err APIError) baseMessage() string {
	switch err.Type {
	case BadRequestError:
		return "Bad request. Invalid syntax when requesting resource from server."
	case UnauthorizedError:
		return "Unauthorized. Trying to access API endpoints with an invalid request signature or access token."
	case ForbiddenError:
		return "Forbidden. Trying to obtain an access token with incorrect client credentials."
	case NotFoundError:
		return "Not found. Trying to access a non-existent endpoint or resource."
	case MethodNotAllowedError:
		return "Method not allowed. Trying to access an endpoint that exists using a method that is not supported by the target resource."
	case NotAcceptableError:
		return "Not acceptable. Trying to access content with an incorrect content type specific in the request header."
	case RequestTimedOutError:
		return "Request timed out. The server did not receive a complete request message within the time that it was prepared to wait."
	case ConflictError:
		return "Conflict. The resource has already been created. It is safe to ignore this error message and continue processing. Returned for DELETE calls when an incorrect version has been specified."
	case GoneError:
		return "Gone. Access to the target resource is no longer available (most likely permanent)."
	case LengthRequiredError:
		return "Length required. The server refuses to accept the request without a defined Content-Length."
	case PayloadTooLargeError:
		return "Payload too large. The server is refusing to process a request because the request payload is larger than the server is willing or able to process."
	case URITooLongError:
		return "URI too long. The server is refusing to service the request because the request-target is longer than the server is willing to interpret."
	case UnsupportedMediaTypeError:
		return "Unsupported media type. The origin server is refusing to service the request because the payload is in a format not supported by this method on the target resource."
	case ExpectationFailedError:
		return "Expecation failed. The expectation given in the request's Expect header field could not be met by at least one of the inbound servers."
	case TooManyRequestsError:
		return "Too many requests. Rate limit for requests per second has been exceeded, please back-off immediately, then retry later."
	case ServerError:
		return "Server error. Internal error occurred or the request timed out. This is safe to retry after waiting a short amount of time."
	case BadGatewayError:
		return "Bad Gateway. Temporary internal networking problem. This is safe to retry after waiting a short amount of time."
	case ServiceUnavailableError:
		return "Service unavailable. Service is temporarily overloaded. This is safe to retry after waiting a short amount of time."
	case GatewayTimeoutError:
		return "Gateway timeout. Temporary internal networking problem. This is safe to retry after waiting a short amount of time."
	}

	return "Unknown error. Please get in touch with your support representative."
}
