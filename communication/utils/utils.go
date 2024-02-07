package utils

import (
	Types "github.com/nileshnk/media_upload_poc/communication/types"
)

func CheckErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrorWithCommonResponse(err error) Types.CommonResponse {
	if err != nil {
		return Types.CommonResponse{
			Success:    false,
			StatusCode: 500,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return Types.CommonResponse{
		Success: true,
		Message: "No Error",
	}
}
