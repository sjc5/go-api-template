package middleware

import (
	"net/http"

	"github.com/sjc5/go-api-template/session"
	"github.com/sjc5/go-api-template/util"
)

func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := util.NewResponse(w)
		s := session.FromContext(r)
		if s == nil {
			res.Unauthorized("")
			return
		}
		next.ServeHTTP(w, r)
	})
}
