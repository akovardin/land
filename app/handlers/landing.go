package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/tools/template"

	"gohome.4gophers.ru/kovardin/land/views"
)

type Landing struct {
	registry *template.Registry
}

func NewLanding(registry *template.Registry) *Landing {
	return &Landing{
		registry: registry,
	}
}

func (l *Landing) Home(c echo.Context) error {
	// https://getbootstrap.com/docs/5.3/examples/
	name := c.PathParam("name")

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"home.html",
	).Render(map[string]any{
		"name": name,
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
