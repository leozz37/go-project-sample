package api

import (
	"log"

	"sample/internal/metrics"
	"sample/internal/services/music"
	"sample/internal/services/weather"

	"github.com/gin-gonic/gin"
)

func (app *Application) AlbumRecomendation(c *gin.Context) {
	city := c.Param("city")

	temp, err := app.Database.GetTemperature(city)
	if err != nil {
		log.Fatal(err)
	}

	if temp == 0 {
		temp, err = weather.GetWeatherByCityName(city)
		if err != nil {
			log.Fatal(err)
		}
		app.Database.SetTemperature(city, temp)
		metrics.CacheSet.Inc()
	} else {
		metrics.SuccessCacheGet.Inc()
	}

	genre := getMusicGenreByTemp(temp)
	album, err := music.GetAlbumRecomendationByTemp(genre)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"temperature": temp,
		"name":        album.Tracks[0].Album.Name,
		"artist":      album.Tracks[0].Artists[0].Name,
		"spotify":     album.Tracks[0].External_urls.URL,
	})
}

func getMusicGenreByTemp(temp float64) string {
	var genre string
	if temp >= 30 {
		genre = "rock"
	} else if temp < 30 && temp >= 15 {
		genre = "country"
	} else if temp < 14 && temp >= 10 {
		genre = "classical"
	} else {
		genre = "jazz"
	}

	return genre
}
