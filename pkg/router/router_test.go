package router

import (
	router_mocks "go-toolbox/pkg/router/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func wrapRecorder(order *[]string, label string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			*order = append(*order, label)
			next.ServeHTTP(w, r)
		})
	}
}

func TestHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	mockRouter := router_mocks.NewRouter(t)
	handler := NewHandler(mockRouter)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	mockRouter.
		On("ServeHTTP", rec, req).
		Once()

	handler.ServeHTTP(rec, req)

	mockRouter.AssertExpectations(t)
}

func TestHandler_Use(t *testing.T) {
	t.Parallel()

	mockRouter := router_mocks.NewRouter(t)
	handler := NewHandler(mockRouter)

	mw1 := func(next http.Handler) http.Handler { return next }
	mw2 := func(next http.Handler) http.Handler { return next }

	mockRouter.
		On("Use", mock.Anything).
		Twice()

	handler.Use(mw1, mw2)

	mockRouter.AssertExpectations(t)
}

func TestHandler_ApplyMiddlewares_Order(t *testing.T) {
	t.Parallel()

	var order []string

	mw1 := wrapRecorder(&order, "mw1")
	mw2 := wrapRecorder(&order, "mw2")

	h := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	h = applyMiddlewares(h, []Middleware{mw1, mw2})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, []string{"mw1", "mw2"}, order)
}

func TestHandler_RegisterRoutes(t *testing.T) {
	t.Parallel()

	var order []string

	mw1 := wrapRecorder(&order, "mw1")
	mw2 := wrapRecorder(&order, "mw2")

	handlerCalled := false
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		order = append(order, "handler")
		handlerCalled = true
	}

	route := Route{
		Method:      http.MethodGet,
		Pattern:     "/",
		Handler:     handlerFunc,
		Middlewares: []Middleware{mw1, mw2},
	}

	mockRouter := router_mocks.NewRouter(t)
	mockRouter.On("Handle", mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(2).(http.Handler)
			req := httptest.NewRequest(route.Method, route.Pattern, nil)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
		}).
		Once()

	handler := NewHandler(mockRouter)
	handler.RegisterRoutes([]Route{route})

	assert.Equal(t, []string{"mw1", "mw2", "handler"}, order)
	assert.True(t, handlerCalled)

	mockRouter.AssertExpectations(t)
}
