package middleware

import (
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/rs/cors"
	"github.com/sjc5/go-api-template/internal/platform"
	"github.com/sjc5/go-api-template/internal/session"
	"github.com/sjc5/kit/pkg/middleware/secureheaders"
)

const OneMegabyteSize = 1048576

func ApplyGlobal(r *chi.Mux) {
	r.Use(
		// Some basic middleware appropriate to apply early
		chimiddleware.RequestID,
		chimiddleware.Logger,
		chimiddleware.Recoverer,
		chimiddleware.RequestSize(OneMegabyteSize),

		// Security middleware
		httprate.LimitByRealIP(1, 1*time.Second),
		cors.New(cors.Options{AllowedOrigins: platform.GetEnv().AllowedOrigins}).Handler,
		secureheaders.Middleware,

		// Some more basic middleware appropriate to apply later
		chimiddleware.Compress(5),
		chimiddleware.Heartbeat("/"),

		// Add session (if any) to request context
		session.ContextMiddleware,
	)
}
