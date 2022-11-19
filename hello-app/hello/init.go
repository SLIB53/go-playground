package main

import (
	"context"
	"local/hello/lib/slogger"
)

type ContextKey string

var ContextKeyLogger ContextKey = "logger"

// setupContext intializes the root application context.
func setupContext() context.Context {
	ctx := context.Background()

	logger := setupLogger()

	return context.WithValue(ctx, ContextKeyLogger, logger)
}

// setupLogger initializes the root application logger.
func setupLogger() *slogger.SLogger {
	return slogger.New()
}
