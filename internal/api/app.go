package api

import (
	"net/http"
	"sample/pkg/models"
	"time"

	"github.com/leozz37/shutall"
)

type Application struct {
	Database *models.Database
}

func (app *Application) RunServer() {
	svr := &http.Server{
		Addr:         ":8080",
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutall.Run(svr, 2*time.Second)
}
