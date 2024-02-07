package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	Types "github.com/nileshnk/media_upload_poc/auth/types"
)

func DecodePayload(payloadBytes []byte) (Types.CommonResponse, error) {
	var parsedData Types.CommonResponse

	errJson2 := json.Unmarshal(payloadBytes, &parsedData)
	if errJson2 != nil {
		fmt.Println(errJson2)
		return parsedData, errJson2
	}
	return parsedData, nil
}

func HttpErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	// w.WriteHeader()

	json.NewEncoder(w).Encode(Types.CommonResponse{
		Success:    false,
		Message:    err.Error(),
		StatusCode: statusCode,
		Data:       nil,
	})
}
