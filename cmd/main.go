package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/benchttp/server"
	"github.com/benchttp/server/firestore"
	"github.com/benchttp/server/shutdown"
)

const (
	defaultConfigPath = ".env"
	defaultPort       = "9998"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go shutdown.ListenInterrupt(cancel)

	shutdownHandle, err := run(ctx, cancel)
	if err != nil {
		shutdownHandle.Call() // nolint
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Wait for interrupt.
	<-ctx.Done()

	if err := shutdownHandle.Call(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run executes the main program. It starts a server.Server inside
// a goroutine and returns a configured shutdown.Handle for gracefully
// shutting down the program.
//
// If the bootstrap of the server fails, run returns a non-nil error.
// Otherwise the caller must use the shutdown.Handle to stop the server.
func run(ctx context.Context, cancel context.CancelFunc) (*shutdown.Handle, error) {
	configPath := flag.String("config", defaultConfigPath, "Path to the configuration file (for example --config .env)")

	flag.Parse()

	config, err := readConfigFile(*configPath)
	if err != nil {
		return &shutdown.Handle{}, err
	}

	rs, err := firestore.NewReportService(context.Background(), config.project, config.collection)
	if err != nil {
		return &shutdown.Handle{}, err
	}

	srv := server.New(config.addr, rs)

	handle := shutdown.NewHandle(func() error {
		if err := srv.Shutdown(ctx); err != nil && err != context.Canceled {
			return err
		}

		if err := rs.Close(); err != nil {
			return err
		}

		return nil
	})

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintln(os.Stderr, err)
			cancel()
		}
	}()

	return handle, nil
}

type config struct {
	addr       string
	project    string
	collection string
}

func defaultConfig() config {
	return config{
		addr: ":" + defaultPort,
	}
}

func readConfigFile(file string) (config, error) {
	config := defaultConfig()

	if err := godotenv.Load(file); err != nil {
		return config, err
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.addr = ":" + port
	}

	config.project = os.Getenv("GOOGLE_PROJECT_ID")
	if config.project == "" {
		return config, errors.New("env variable GOOGLE_PROJECT_ID is not defined")
	}

	config.collection = os.Getenv("FIRESTORE_COLLECTION_ID")
	if config.collection == "" {
		return config, errors.New("env variable FIRESTORE_COLLECTION_ID is not defined")
	}

	return config, nil
}
