package session

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
)

type Session struct {
	UserID    string
	SessionID string
}

func getSessionFromWhereverYouWant(r *http.Request) (*Session, error) {
	if rand.Float64() < 0.5 {
		return &Session{UserID: "123", SessionID: "456"}, nil
	}
	return nil, errors.New("no session")
}

type contentKey string

const sessionContextKey contentKey = "Session"

func FromContext(r *http.Request) *Session {
	session, ok := r.Context().Value(sessionContextKey).(*Session)
	if !ok {
		return nil
	}
	return session
}

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := getSessionFromWhereverYouWant(r)
		ctx := context.WithValue(r.Context(), sessionContextKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
