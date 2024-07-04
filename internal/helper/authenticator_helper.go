package helper

import (
	"customer/internal/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func Authenticator(tokenAuth *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := jwtauth.VerifyRequest(tokenAuth, r, jwtauth.TokenFromHeader)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response.Response[any]{
					Status:     "Failed",
					StatusCode: http.StatusUnauthorized,
					Message:    "Unauthorized",
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
