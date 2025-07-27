package responses

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewErrorResponse(errs ...error) ErrorResponse {
	out := make([]string, len(errs))
	for i, e := range errs {
		out[i] = e.Error()
	}
	return ErrorResponse{Errors: out}
}
