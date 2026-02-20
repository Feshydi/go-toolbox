package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []Middleware
}

//go:generate go run github.com/vektra/mockery/v2@v2.53.5 --name=Router --output=./mocks --outpkg=router_mocks --exported
type Router interface {
	http.Handler
	Use(middlewares ...func(http.Handler) http.Handler)
	Handle(method, pattern string, handler http.Handler)
}

type Handler struct {
	router Router
}

func NewHandler(router Router) *Handler {
	h := &Handler{
		router: router,
	}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) Use(middlewares ...Middleware) {
	for _, mw := range middlewares {
		h.router.Use(mw)
	}
}

func (h *Handler) RegisterRoutes(routes []Route) {
	for _, rt := range routes {
		handler := applyMiddlewares(rt.Handler, rt.Middlewares)
		h.router.Handle(rt.Method, rt.Pattern, handler)
	}
}

func applyMiddlewares(h http.Handler, mws []Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
