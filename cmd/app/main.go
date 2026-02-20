package main

import (
	"context"
	"go-toolbox/internal/build"
	"go-toolbox/internal/config"
	"go-toolbox/pkg/shutdown"
	"time"
)

func main() {
	cfg := config.MustLoad()

	app := build.Create(cfg)

	app.Start()

	shutdown.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Stop(ctx)
}
