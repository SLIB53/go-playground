package main

import (
	"context"
	"errors"
	"fmt"
	"local/hello/api"
	"local/hello/lib/slogger"
)

func main() {
	ctx := setupContext()

	var in string
	fmt.Scan(&in)

	sayHelloHandler(ctx, in)
}

// sayHelloHandler handles the SayHello API.
func sayHelloHandler(context context.Context, name string) {
	logger := context.Value(ContextKeyLogger).(*slogger.SLogger)

	if hello, err := api.SayHello(context, logger, name); err == nil {
		fmt.Println(hello)
	} else {
		if errors.Is(err, api.ErrInvalidName) {
			fmt.Println("Uh-oh! That's not a valid name.")
		}
	}
}
