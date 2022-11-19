package api

import (
	"context"
	"errors"
	"fmt"
)

type Logger interface {
	LogInfo(message string)
	LogError(message string, err error)
}

var ErrInvalidName = errors.New("invalid name")

// SayHello says hello to the given name.
func SayHello(_ context.Context, logger Logger, name string) (string, error) {
	hello, err := makeHello(name)
	if err != nil {
		err = fmt.Errorf(
			"SayHello failed to create hello from name %s. Internal Error: %w",
			name,
			err,
		)

		logger.LogError("API Error: ", err)

		return "", err
	}

	return hello, nil
}

// makeHello creates a hello string.
func makeHello(name string) (string, error) {
	if name == "" {
		return "", ErrInvalidName
	}

	if len(name) > 32 {
		return "", ErrInvalidName
	}

	return "Hello, " + name + "!", nil
}
