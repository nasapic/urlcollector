package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/nasapic/urlcollector/internal/app"
)

type contextKey string

const (
	appName = "urlcollector"
)

var (
	a app.App
)

func main() {
	cfg := app.LoadConfig()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	initExitMonitor(ctx, cancel)

	// App
	a, err := app.NewApp(appName, cfg)
	if err != nil {
		exit(err)
	}

	// Init service
	a.Init()

	// Start service
	a.Start()

	log.Fatalf("%s stoped: %s", appName, err)
}

func exit(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func initExitMonitor(ctx context.Context, cancel context.CancelFunc) {
	go checkSigterm(cancel)
	go checkCancel(ctx)
}

func checkSigterm(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()
}

func checkCancel(ctx context.Context) {
	<-ctx.Done()
	a.Stop()
	os.Exit(1)
}
