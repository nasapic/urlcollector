package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/nasapic/base"
)

type (
	textResponse string
)

func (app *App) NewWebRouter() *base.Router {
	rt := app.NewJSONAPIRouter()

	app.addJSONAPICollectorRouter(rt)

	return rt
}

func (app *App) NewJSONAPIRouter() *base.Router {
	rt := base.NewRouter("json-api-home-router")
	return rt
}

func (t textResponse) write(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(t))
}

func (app *App) addJSONAPICollectorRouter(parent chi.Router) chi.Router {
	return parent.Route("/pictures", func(child chi.Router) {
		child.Get("/?from={from}&to={to}", app.JSONAPIEndpoint.SearchURLs)
	})
}
