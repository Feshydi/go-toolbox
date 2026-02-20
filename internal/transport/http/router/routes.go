package router

import (
	"go-toolbox/pkg/router"
	"net/http"
)

type Example interface {
	HandlerFunc() http.HandlerFunc
}

func Routes(example Example, exampleMiddleware router.Middleware) []router.Route {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/api/example-handler",
			Handler: example.HandlerFunc(),
			Middlewares: []router.Middleware{
				exampleMiddleware,
			},
		},
	}
}
