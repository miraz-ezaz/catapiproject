package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	beego.Controller
}

type FavoriteImage struct {
	URL string `json:"url"`
}

// Fetches the favorite images from The Cat API
func fetchFavoriteImages(apiKey string, ch chan<- []FavoriteImage, done chan<- bool) {
	defer close(ch)
	defer close(done)

	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://api.thecatapi.com/v1/favourites"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating GET request:", err)
		done <- true
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making GET request:", err)
		done <- true
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading GET response body:", err)
		done <- true
		return
	}

	var response []struct {
		Image struct {
			URL string `json:"url"`
		} `json:"image"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		done <- true
		return
	}

	var favoriteImages []FavoriteImage
	for _, item := range response {
		favoriteImages = append(favoriteImages, FavoriteImage{URL: item.Image.URL})
	}

	ch <- favoriteImages
	done <- true
}

func (c *FavoritesController) Get() {
	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")

	favoriteImagesCh := make(chan []FavoriteImage)
	favoriteImagesDone := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fetchFavoriteImages(apiKey, favoriteImagesCh, favoriteImagesDone)
		wg.Wait()
	}()

	favoriteImages := <-favoriteImagesCh

	c.Data["FavoriteImages"] = favoriteImages
	c.TplName = "favorites.tpl"
}
