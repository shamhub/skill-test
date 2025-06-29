package error

import "fmt"

type FailedRequest struct {
	URL string
	Err error
}

func (f FailedRequest) Error() string {
	return fmt.Sprintf("error in making a service request. URL: %v error: %v", f.URL, f.Err)
}

type UnSupportedResponse struct {
	ResponseType string
}

func (u UnSupportedResponse) Error() string {
	return u.ResponseType
}
