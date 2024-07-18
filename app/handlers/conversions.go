package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

const (
	networkYandex = "yandex"
	networkVK     = "vk"
)

type Conversions struct {
	app *pocketbase.PocketBase
}

func NewConversions(app *pocketbase.PocketBase) *Conversions {
	return &Conversions{
		app: app,
	}
}

func (h *Conversions) Fire(c echo.Context) error {
	h.app.Logger().Info("conversion fire request", "url", c.Request().URL.String())

	slug := c.PathParam("name")
	landing, err := h.app.Dao().FindFirstRecordByFilter(
		"landings",
		"slug = {:slug} && enabled = true",
		dbx.Params{"slug": slug},
	)
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	conversions, err := h.app.Dao().FindCollectionByNameOrId("conversions")
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	record := models.NewRecord(conversions)

	yclid := c.QueryParam("yclid")
	rbclickid := c.QueryParam("rb_clickid")
	key := yclid + rbclickid

	if key == "" {
		return nil
	}

	network := networkVK
	if yclid != "" {
		network = networkYandex
	}

	record.Set("yclid", yclid)
	record.Set("rb_clickid", rbclickid)
	record.Set("key", key)
	record.Set("uploaded", false)
	record.Set("network", network)
	record.Set("landing", landing.Id)

	if err := h.app.Dao().SaveRecord(record); err != nil {
		h.app.Logger().Error("error on save conversions", "error", err)
		return apis.NewApiError(http.StatusInternalServerError, "error on save", err)
	}

	return nil
}
