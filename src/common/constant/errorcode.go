package constant

type ErrorCode string

const (
	ErrorCodeBadRequest          ErrorCode = "BAD_REQUEST"
	ErrorCodeResourceNotFound    ErrorCode = "RESOURCE_NOT_FOUND"
	ErrorCodeInternalServerError ErrorCode = "INTERNAL_SERVER"
)
