package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func getenv(key, source string) string {
	envFunctions := map[string]func(key string) string{
		"local":            func(key string) string { return os.Getenv(key) },
		"ssm":              func(key string) string { /* get from ssm */ return key },
		"remoteRepository": func(key string) string { /* get from remote */ return key },
	}
	// Alternative implementation
	// switch source {
	// case "local":
	// 	return os.Getenv(key)
	// case "ssm":
	// 	// get from ssm
	// 	return key
	// case "remoteRepository":
	// 	// get from remote
	// 	return key
	// default:
	// 	panic(fmt.Sprintf("Unknown source: %s", source))
	// }
	return envFunctions[source](key)
}

func run(
	ctx context.Context,
	args []string,
	getenv func(string, string) string,
	stdin io.Reader,
	stdout io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, getenv, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
