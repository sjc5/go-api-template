package middleware

import (
	"net/http"

	"github.com/sjc5/go-api-template/internal/session"
	"github.com/sjc5/kit/pkg/response"
)

func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := response.New(w)
		s := session.FromContext(r)
		if s == nil {
			res.Unauthorized("")
			return
		}
		next.ServeHTTP(w, r)
	})
}
