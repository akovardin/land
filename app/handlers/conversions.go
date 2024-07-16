package handlers

import "github.com/labstack/echo/v5"

type Conversions struct {
}

func NewConversions() *Conversions {
	return &Conversions{}
}

func (h *Conversions) Fire(c echo.Context) error {
	// make conversions fired
	return nil
}
