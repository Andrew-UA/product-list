package responses

type ErrorResponse struct {
	ResponseJson
	Errors []error `json:"errors"`
}

// TODO:
func NewErrorResponse(errs ...error) *ErrorResponse {
	response := &ErrorResponse{
		Errors: errs,
	}

	return response
}
