package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/template"
	"go.uber.org/fx"

	"gohome.4gophers.ru/kovardin/land/app/handlers"
	"gohome.4gophers.ru/kovardin/land/static"
)

func main() {
	fx.New(
		fx.Provide(pocketbase.New),
		fx.Provide(template.NewRegistry),
		fx.Provide(handlers.NewLanding),
		fx.Invoke(routing),
	).Run()
}

func routing(app *pocketbase.PocketBase, lc fx.Lifecycle, registry *template.Registry, landing *handlers.Landing) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// prints "Hello!" every 2 minutes
		scheduler.MustAdd("hello", "*/2 * * * *", func() {
			log.Println("Hello!")
		})

		scheduler.Start()

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.GET("/l/:name", landing.Home)
		e.Router.GET("/l/:name/terms", landing.Terms)
		e.Router.GET("/l/:name/privacy", landing.Privacy)

		e.Router.GET("/static/*", func(c echo.Context) error {
			p := c.PathParam("*")

			path, err := url.PathUnescape(p)
			if err != nil {
				return fmt.Errorf("failed to unescape path variable: %w", err)
			}

			err = c.FileFS(path, static.FS)
			if err != nil && errors.Is(err, echo.ErrNotFound) {
				return c.FileFS("index.html", static.FS)
			}

			return err
		})

		return nil

	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Start()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
