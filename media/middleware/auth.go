package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/nileshnk/media_upload_poc/media/dapr"
	Utils "github.com/nileshnk/media_upload_poc/media/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			break
		case "/upload":
			valid := validateToken(r)
			if !valid {
				Utils.HttpErrorResponse(w, errors.New("Invalid Token"), http.StatusUnauthorized)
				return
			}
			break
		default:
			w.WriteHeader(http.StatusNotFound)
		}

		next.ServeHTTP(w, r)
	})
}

func validateToken(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	valid := true
	if auth == "" {
		valid = false
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		valid = false
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if ok := dapr.InvokeAuthTokenValidate(r.Context(), token); !ok {
		valid = false
	}
	return valid
}
