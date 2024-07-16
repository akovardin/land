package handlers

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"go.uber.org/zap"
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
	// make conversions fired

	// create items

	// https://redirect.appmetrica.yandex.com/serve/821509867285037527?yclid={yclid}
	// https://redirect.appmetrica.yandex.com/serve/821509867285037527?yclid=444
	// https://land.kovardin.ru/l/downloader?yclid&rbclickid&referrer=appmetrica_tracking_id%3D821509867285037527%26ym_tracking_id%3D749562427864121681

	// /l/downloader/fire?client_id=&yclid=&install_timestamp=1721163558&appmetrica_device_id=5446207023320664663&click_id=&transaction_id=cpi14539061032550945951&match_type=&tracker=appmetrica_821509867285037527&rb_clickid=
	// https://yandex.com/support/direct/statistics/url-tags.html
	// https://appmetrica.yandex.com/docs/en/mobile-tracking/tracking-specification
	h.app.Logger().Info("fire", zap.Any("url", c.Request().URL.String()))

	fmt.Println(c.Request().URL.String())

	return nil
}
