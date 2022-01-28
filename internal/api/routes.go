package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func (app *Application) Routes() http.Handler {
	router := gin.Default()

	// Setup Prometheus
	prom := ginprometheus.NewPrometheus("gin")
	prom.Use(router)

	router.GET("/recomendation/:city", app.AlbumRecomendation)

	return router
}
