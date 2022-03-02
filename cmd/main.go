package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/benchttp/server"
	"github.com/benchttp/server/firestore"
)

const (
	defaultConfigPath = ".env"
	defaultPort       = "9998"
)

func main() {
	closeHandle, err := run()

	if err != nil {
		closeHandle.close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := closeHandle.close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// shutdownHandle
type shutdownHandle struct {
	closeFunc func() error
}

func (c *shutdownHandle) close() error {
	return c.closeFunc()
}

func run() (*shutdownHandle, error) {
	c := &shutdownHandle{}

	configPath := flag.String("config", defaultConfigPath, "")

	flag.Parse()

	config, err := readConfigFile(*configPath)
	if err != nil {
		return c, err
	}

	rs, err := firestore.NewReportService(context.Background(), config.project, config.collection)
	if err != nil {
		return c, err
	}

	srv := server.New(config.addr, rs)

	c.closeFunc = func() error {
		if err := srv.Close(); err != nil {
			return err
		}

		if err := rs.Close(); err != nil {
			return err
		}

		return nil
	}

	return c, srv.Start()
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
