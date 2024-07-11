package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/tools/template"

	"gohome.4gophers.ru/kovardin/land/views"
)

type Landing struct {
	registry *template.Registry
	app      *pocketbase.PocketBase
}

func NewLanding(registry *template.Registry, app *pocketbase.PocketBase) *Landing {
	return &Landing{
		registry: registry,
		app:      app,
	}
}

func (l *Landing) Home(c echo.Context) error {
	slug := c.PathParam("name")
	landing, err := l.app.Dao().FindFirstRecordByFilter(
		"landings",
		"slug = {:slug}",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	home, err := l.app.Dao().FindFirstRecordByFilter(
		"home",
		"landing = {:landing}",
		dbx.Params{"landing": landing.Id},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	// todo: parse

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"home.html",
	).Render(map[string]any{
		"slug":            slug,
		"metaTitle":       landing.GetString("title"),
		"metaDescription": landing.GetString("description"),
		"metaKeywords":    landing.GetString("keywords"),
		"title":           home.GetString("title"),
		"description":     home.GetString("description"),

		// links
		"rustoreLink": home.GetString("rustore_link"),
		"androidLink": home.GetString("android_link"),
		"huaweiLink":  home.GetString("android_link"),

		// counters
		"yandexCounter": landing.GetString("yandex_counter"),
		"mtCounter":     landing.GetString("mt_counter"),
	})

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}

func (l *Landing) Terms(c echo.Context) error {
	name := c.PathParam("name")

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"terms.html",
	).Render(map[string]any{
		"name": name,
	})

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}

func (l *Landing) Privacy(c echo.Context) error {
	name := c.PathParam("name")

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"privacy.html",
	).Render(map[string]any{
		"name": name,
	})

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}
