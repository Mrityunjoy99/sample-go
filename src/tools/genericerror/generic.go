package genericerror

import "github.com/Mrityunjoy99/sample-go/src/common/constant"


type ErrorDetails map[string]interface{}

type GenericError interface {
	error
	GetCode() constant.ErrorCode
	GetMessage() string
	GetDetails() ErrorDetails
	GetErr() error
}

type genericError struct {
	code    constant.ErrorCode
	message string
	details ErrorDetails
	err     error
}

func (e *genericError) Error() string {
	return e.message
}

func (e *genericError) GetCode() constant.ErrorCode {
	return e.code
}

func (e *genericError) GetMessage() string {
	return e.message
}

func (e *genericError) GetDetails() ErrorDetails {
	return e.details
}

func (e *genericError) GetErr() error {
	return e.err
}

func NewGenericError(code constant.ErrorCode, message string, details ErrorDetails, err error) GenericError {
	return &genericError{
		code:    code,
		message: message,
		details: details,
		err:     err,
	}
}

func NewInternalErrByErr(err error) GenericError {
	return &genericError{
		code:    constant.ErrorCodeInternalServerError,
		message: err.Error(),
		details: nil,
		err:     err,
	}
}
