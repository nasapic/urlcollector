package app

import (
	"fmt"
	"net/http"
	"sync"

	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/jsonapi"
)

type (
	// App description
	App struct {
		*base.App

		*Config

		JSONAPIEndpoint *jsonapi.Endpoint
	}
)

// NewApp initializes new App worker instance
func NewApp(name string, cfg *Config, log base.Logger) *App {
	app := App{
		App:             base.NewApp(name, log),
		Config:          cfg,
		JSONAPIEndpoint: jsonapi.NewEndpoint("json-api-endpoint"),
	}

	// Router
	app.JSONAPIRouter = app.NewJSONAPIRouter()

	return &app
}

// Init app
func (app *App) Init() error {
	return nil
}

// Start app
func (app *App) Start() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		app.StartJSONAPI()
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop the app
}

func (app *App) StartJSONAPI() error {
	p := fmt.Sprintf(":%d", app.Config.Server.JSONAPIPort)

	app.Log.Info("JSON API Server initializing...", "port", p)

	err := http.ListenAndServe(p, app.JSONAPIRouter)

	return err
}
