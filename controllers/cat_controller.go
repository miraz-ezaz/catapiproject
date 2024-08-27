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

type CatController struct {
	beego.Controller
}

func fetchCatImages(apiKey string, ch chan<- string) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		ch <- ""
		return
	}
	req.Header.Set("x-api-key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making API request:", err)
		ch <- ""
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		ch <- ""
		return
	}
	log.Println("API response body:", string(body))
	ch <- string(body)
}

func (c *CatController) Get() {
	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fetchCatImages(apiKey, ch)
	}()

	wg.Wait()
	close(ch)

	images := <-ch
	if images == "" {
		log.Println("No images fetched")
		c.Data["Error"] = "Failed to fetch images"
		c.TplName = "cat.tpl"
		return
	}

	// Parse the JSON response to extract the image URL using a map
	var catImages []map[string]interface{}
	err := json.Unmarshal([]byte(images), &catImages)
	if err != nil || len(catImages) == 0 {
		log.Println("Error parsing JSON or no images found:", err)
		c.Data["Error"] = "Failed to parse images"
		c.TplName = "cat.tpl"
		return
	}

	// Access the first element and get the URL
	imageURL, ok := catImages[0]["url"].(string)
	if !ok {
		log.Println("Error: URL not found or invalid")
		c.Data["Error"] = "Failed to retrieve image URL"
		c.TplName = "cat.tpl"
		return
	}
	log.Println("Fetched image URL:", imageURL)

	// Pass the image URL to the view
	c.Data["ImageURL"] = imageURL
	c.TplName = "cat.tpl"
}
