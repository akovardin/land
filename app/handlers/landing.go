package handlers

import (
	"net/http"
	"net/url"
	"path"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/template"
	"go.uber.org/zap"

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
		"slug = {:slug} && enabled = true",
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

	features, err := l.app.Dao().FindRecordsByFilter(
		"features",
		"landing.slug = {:slug} && enabled = true",
		"-created",
		4,
		0,
		dbx.Params{"slug": slug},
	)
	if err != nil {
		features = []*models.Record{}
	}

	rustoreLink := home.GetString("rustore_link")

	parts, _ := url.Parse(rustoreLink)
	bundleName := path.Base(parts.Path)

	img := home.GetString("image")

	// var rustoreLink = "intent://apps.rustore.ru/app/com.roblox.client#Intent;scheme=rustore;package=ru.vk.store;S.browser_fallback_url=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.roblox.client%26hl%3Den;end";

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"landing/home.html",
	).Render(map[string]any{
		"slug":            slug,
		"metaTitle":       landing.GetString("title"),
		"metaDescription": landing.GetString("description"),
		"metaKeywords":    landing.GetString("keywords"),
		"title":           home.GetString("title"),
		"description":     home.GetString("description"),
		"cta":             home.GetString("cta"),
		"image":           "/api/files/home/" + home.Id + "/" + img,

		// links
		"rustoreLink": rustoreLink,
		"bundleName":  bundleName,
		"androidLink": home.GetString("android_link"),
		"huaweiLink":  home.GetString("huawei_link"),
		"directLink":  home.GetString("direct_link"),

		// counters
		"yandexCounter": landing.GetString("yandex_counter"),
		"mtCounter":     landing.GetString("mt_counter"),

		// features
		"features": features,
	})

	l.app.Logger().Info("features", zap.Any("f", features))

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}

func (l *Landing) Terms(c echo.Context) error {
	slug := c.PathParam("name")
	landing, err := l.app.Dao().FindFirstRecordByFilter(
		"landings",
		"slug = {:slug} && enabled = true",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	privacy, err := l.app.Dao().FindFirstRecordByFilter(
		"terms",
		"landing.slug = {:slug}",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	appurl := l.app.Settings().Meta.AppUrl + "/l/" + slug + "terms/"

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"landing/terms.html",
	).Render(map[string]any{
		"slug":      slug,
		"metaTitle": landing.GetString("title"),
		"appname":   landing.GetString("appname"),
		"developer": privacy.GetString("developer"),
		"appurl":    appurl,
	})

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}

func (l *Landing) Privacy(c echo.Context) error {
	slug := c.PathParam("name")
	landing, err := l.app.Dao().FindFirstRecordByFilter(
		"landings",
		"slug = {:slug} && enabled = true",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	privacy, err := l.app.Dao().FindFirstRecordByFilter(
		"privacy",
		"landing.slug = {:slug}",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	appurl := l.app.Settings().Meta.AppUrl + "/l/" + slug + "privacy/"

	html, err := l.registry.LoadFS(views.FS,
		"layout.html",
		"landing/privacy.html",
	).Render(map[string]any{
		"slug":             slug,
		"metaTitle":        landing.GetString("title"),
		"appname":          landing.GetString("appname"),
		"developer":        privacy.GetString("developer"),
		"collectFio":       privacy.GetBool("collectFio"),
		"collectPhone":     privacy.GetBool("collectPhone"),
		"collectAnalytics": privacy.GetBool("collectAnalytics"),
		"collectEmail":     privacy.GetBool("collectEmail"),
		"email":            privacy.GetString("email"),
		"appurl":           appurl,
	})

	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	return c.HTML(http.StatusOK, html)
}
