package controllers

import (
	"bytes"
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

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type FavResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type VoteResponse struct {
	Message     string `json:"message"`
	ID          int    `json:"id"`
	ImageID     string `json:"image_id"`
	SubID       string `json:"sub_id"`
	Value       int    `json:"value"`
	CountryCode string `json:"country_code"`
}

func fetchCatImages(apiKey string, ch chan<- CatImage, done chan<- bool) {
	client := &http.Client{Timeout: 10 * time.Second}
	log.Println("Creating new HTTP request...")
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		done <- true
		return
	}

	log.Println("Setting API key header...")
	req.Header.Set("x-api-key", apiKey)
	log.Println("Sending request to The Cat API...")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making API request:", err)
		done <- true
		return
	}
	defer resp.Body.Close()

	log.Println("Reading response body...")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		done <- true
		return
	}
	log.Println("API response body:", string(body))

	var catImages []CatImage
	err = json.Unmarshal(body, &catImages)
	if err != nil || len(catImages) == 0 {
		log.Println("Error unmarshaling JSON or no images found:", err)
		done <- true
		return
	}

	ch <- catImages[0]
	done <- true
}

func makeFavRequest(apiKey, imageID string, ch chan<- FavResponse, done chan<- bool) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://api.thecatapi.com/v1/favourites"

	payload := map[string]string{
		"image_id": imageID,
		"sub_id":   "my-user-1234", // Hardcoded sub_id
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating POST request:", err)
		done <- true
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	log.Println("Sending POST request to Favourites API...")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making POST request:", err)
		done <- true
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading POST response body:", err)
		done <- true
		return
	}
	log.Println("Favourites API response body:", string(body))

	var favResp FavResponse
	err = json.Unmarshal(body, &favResp)
	if err != nil || favResp.Message != "SUCCESS" {
		log.Println("Failed to add to favourites:", err)
		done <- true
		return
	}

	ch <- favResp
	done <- true
}

func makeVoteRequest(apiKey, imageID string, value int, ch chan<- VoteResponse, done chan<- bool) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://api.thecatapi.com/v1/votes"

	payload := map[string]interface{}{
		"image_id": imageID,
		"sub_id":   "my-user-1234", // Hardcoded sub_id
		"value":    value,
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating POST request:", err)
		done <- true
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	log.Println("Sending POST request to Votes API...")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making POST request:", err)
		done <- true
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading POST response body:", err)
		done <- true
		return
	}
	log.Println("Votes API response body:", string(body))

	var voteResp VoteResponse
	err = json.Unmarshal(body, &voteResp)
	if err != nil || voteResp.Message != "SUCCESS" {
		log.Println("Failed to submit vote:", err)
		done <- true
		return
	}

	ch <- voteResp
	done <- true
}

func (c *CatController) Get() {
	log.Println("Starting API request...")

	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")
	log.Println("API Key fetched:", apiKey)

	ch := make(chan CatImage)
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Println("Making API call to The Cat API...")

		timeout := time.After(15 * time.Second)

		go fetchCatImages(apiKey, ch, done)

		select {
		case <-done:
			log.Println("API call completed within time.")
		case <-timeout:
			log.Println("API call timed out.")
		}
		wg.Wait()

		close(ch)
	}()

	log.Println("API call completed, waiting for channel response...")

	catImage := <-ch
	log.Println("Data received from channel.")

	if catImage.URL == "" {
		log.Println("No data received from API")
		c.Data["Error"] = "Failed to fetch image"
		c.TplName = "home.tpl"
		return
	}

	log.Println("Fetched image ID:", catImage.ID)
	log.Println("Fetched image URL:", catImage.URL)

	c.Data["ImageID"] = catImage.ID
	c.Data["ImageURL"] = catImage.URL
	c.TplName = "home.tpl"
}

func (c *CatController) Post() {
	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")
	imageID := c.GetString("image_id")
	action := c.GetString("action")

	if action == "fav" {
		favCh := make(chan FavResponse)
		done := make(chan bool)
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			log.Println("Starting Favourites API request...")

			go makeFavRequest(apiKey, imageID, favCh, done)

			select {
			case favResp := <-favCh:
				log.Println("Favourites API response:", favResp)
				if favResp.Message == "SUCCESS" {
					log.Println("Image successfully added to favourites with ID:", favResp.ID)
				}
			case <-done:
				log.Println("Favourites API request completed.")
			}
		}()

		wg.Wait()
		close(favCh)
	} else if action == "like" || action == "dislike" {
		value := 1
		if action == "dislike" {
			value = -1
		}

		voteCh := make(chan VoteResponse)
		done := make(chan bool)
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			log.Println("Starting Vote API request...")

			go makeVoteRequest(apiKey, imageID, value, voteCh, done)

			select {
			case voteResp := <-voteCh:
				log.Println("Votes API response:", voteResp)
				if voteResp.Message == "SUCCESS" {
					log.Println("Vote successfully submitted with ID:", voteResp.ID)
				}
			case <-done:
				log.Println("Vote API request completed.")
			}
		}()

		wg.Wait()
		close(voteCh)
	}

	c.Redirect("/", http.StatusFound)
}
