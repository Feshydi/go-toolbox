package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi *chi.Mux
}

func NewRouter() *Router {
	return &Router{chi: chi.NewRouter()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chi.ServeHTTP(w, req)
}

func (r *Router) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		r.chi.Use(func(next http.Handler) http.Handler {
			return mw(next)
		})
	}
}

func (r *Router) Handle(method, pattern string, handler http.Handler) {
	r.chi.Method(method, pattern, handler)
}
