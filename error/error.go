package error

import (
	"fmt"
)

type CustomError struct {
	Err     error
	Message string
	Code    string
}

// Error Code
const (
	InvalidRequest      = "error/"
	InvalidResponse     = "InvalidResponse"
	InvalidConfig       = "InvalidConfig"
	InvalidAuth         = "InvalidAuth"
	InvalidEmail        = "InvalidEmail"
	InvalidFile         = "InvalidFile"
	InvalidFileUpload   = "InvalidFileUpload"
	InvalidFileDownload = "InvalidFileDownload"
	InvalidFileDelete   = "InvalidFileDelete"
	InvalidFileList     = "InvalidFileList"
	InvalidFileRename   = "InvalidFileRename"
	InvalidFileMove     = "InvalidFileMove"
)

var ErrorCodeMap = map[string]string{
	"0001": "Invalid Request",
	"0002": "Invalid Response",
}

func (e *CustomError) Error() string {

	return fmt.Sprintf("%s: %s", e.Message, e.Err)
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func New(message string, err error) *CustomError {

	return &CustomError{
		Err:     err,
		Message: message,
	}
}
