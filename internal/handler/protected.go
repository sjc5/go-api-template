package handler

import (
	"fmt"
	"net/http"

	"github.com/sjc5/go-api-template/internal/session"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	session := session.FromContext(r)
	msg := fmt.Sprintf("Protected content\n\nUser ID: %s", session.UserID)
	w.Write([]byte(msg))
}
