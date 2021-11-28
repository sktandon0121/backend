package utils

type CommonErrors struct {
	ResponseCode int    `json:"responseCode"`
	ErrorMessage string `json:"errorMessage"`
	Error        error  `json:"error"`
}

type CommonErrorResponse interface {
	AccessDenied(err error, message string) *CommonErrors
	BadRequest(err error, message string) *CommonErrors
	InternalServerError(err error, message string) *CommonErrors
}
type errSvc struct{}

func NewCommonErrorResponse() CommonErrorResponse {
	return &errSvc{}
}

func (e *errSvc) AccessDenied(err error, message string) *CommonErrors {
	return &CommonErrors{ResponseCode: 401, ErrorMessage: message, Error: err}
}

func (e *errSvc) BadRequest(err error, message string) *CommonErrors {
	return &CommonErrors{ResponseCode: 400, ErrorMessage: message, Error: err}
}

func (e *errSvc) InternalServerError(err error, message string) *CommonErrors {
	return &CommonErrors{ResponseCode: 500, ErrorMessage: message, Error: err}
}
