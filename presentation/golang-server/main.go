package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/BlueberryBuns/presentation/golang-server/server"
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

	logger := log.New(stdout, "golang-server: ", log.LstdFlags)
	srv := server.NewServer(logger)
	config := server.GetDefaultConfig()
	httpServer := &http.Server{Addr: net.JoinHostPort(config.Address, config.Port), Handler: srv}
	go func() {
		fmt.Println("Server is running on: ", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, getenv, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
