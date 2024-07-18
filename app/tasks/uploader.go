package tasks

import (
	"github.com/pocketbase/pocketbase"
)

type Uploader struct {
	app *pocketbase.PocketBase
}

func NewUploader(app *pocketbase.PocketBase) *Uploader {
	return &Uploader{
		app: app,
	}
}

func (u *Uploader) Do() {
	// upload fired conversions
	u.app.Logger().Info("upload conversions")
}
