package music

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Album struct {
	Tracks []struct {
		Artists []struct {
			Name string `json:"name"`
		}
		Album struct {
			Name string `json:"name"`
		}
		External_urls struct {
			URL string `json:"spotify"`
		}
	}
}

func GetAlbumRecomendationByTemp(genre string) (Album, error) {
	bearerToken := "Bearer " + os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyURL := os.Getenv("SPOTIFY_URL") + fmt.Sprintf("limit=1&market=ES&seed_genres=%s", genre)

	req, err := http.NewRequest("GET", spotifyURL, nil)
	req.Header.Set("Authorization", bearerToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Album{}, nil
	}
	defer resp.Body.Close()

	var album Album
	json.NewDecoder(resp.Body).Decode(&album)

	return album, nil
}
