package router

import (
	"go-toolbox/internal/transport/http/router/chi"
	"go-toolbox/pkg/router"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func CreateHandler(routes []router.Route) http.Handler {
	h := router.NewHandler(chi.NewRouter())

	h.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
	)

	h.RegisterRoutes(routes)

	return h
}
