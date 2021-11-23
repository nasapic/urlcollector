package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/base"
)

type (
	textResponse string
)

func (app *App) NewWebRouter() *base.Router {
	rt := app.NewJSONAPIRouter()

	return rt
}

func (app *App) NewJSONAPIRouter() *base.Router {
	rt := base.NewRouter("json-api-home-router")

	// NOTE: Make these values configurables
	rt.SetHourlyRate(30)
	rt.SetDailyRate(50)

	app.addJSONAPICollectorRouter(rt)

	return rt
}

func (t textResponse) write(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(t))
}

func (app *App) addJSONAPICollectorRouter(parent chi.Router) chi.Router {
	return parent.Route("/pictures", func(child chi.Router) {
		child.Get("/", app.JSONAPIEndpoint.SearchURLs)
	})
}
