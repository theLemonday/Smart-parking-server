package presenter

type Response struct {
	Type string `json:"type"`
}

type ErrorResponse struct {
	Response
	Message string `json:"msg,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func NewErrorResponseWithCode(err error, errCode int) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{Type: "error"},
		Message:  err.Error(),
		Code:     errCode,
	}
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{Type: "error"},
		Message:  err.Error(),
	}
}
