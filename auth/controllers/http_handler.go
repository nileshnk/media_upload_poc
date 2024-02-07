package controllers

import (
	"encoding/json"
	"net/http"

	Types "github.com/nileshnk/media_upload_poc/auth/types"
	Utils "github.com/nileshnk/media_upload_poc/auth/utils"
)

func ValidateTokenFromRequestBody(w http.ResponseWriter, r *http.Request) {
	var token Types.Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		Utils.HttpErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	valid, tokenData, errValidate := ValidateAccessToken(token.Token)
	if !valid {
		Utils.HttpErrorResponse(w, errValidate, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(Types.CommonResponse{
		Success:    true,
		Message:    "Token Valid",
		StatusCode: http.StatusOK,
		Data:       tokenData,
	})
}
