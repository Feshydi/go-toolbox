package build

import (
	"fmt"
	"go-toolbox/internal/config"
	example_service "go-toolbox/internal/service/example"
	example_handler "go-toolbox/internal/transport/http/handler/example"
	example_middleware "go-toolbox/internal/transport/http/middleware/example"
	"go-toolbox/internal/transport/http/router"
	"go-toolbox/pkg/logger"
	"net/http"
)

func Create(cfg *config.Config) *Application {
	logger := logger.New(cfg.Logger.Level)

	exampleService := example_service.New()
	exampleHandler := example_handler.New(exampleService, logger)
	exampleMiddleware := example_middleware.Middleware(logger)

	routes := router.Routes(exampleHandler, exampleMiddleware)
	handler := router.CreateHandler(routes)

	return &Application{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
		logger: logger,
	}
}
