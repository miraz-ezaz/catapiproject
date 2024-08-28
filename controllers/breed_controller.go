package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	beego.Controller
}

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BreedDetails struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Origin       string `json:"origin"`
	Description  string `json:"description"`
	WikipediaURL string `json:"wikipedia_url"`
}

type BreedImage struct {
	URL string `json:"url"`
}

// Fetches the list of all breeds
func fetchAllBreeds(apiKey string, ch chan<- []Breed) {
	defer close(ch)

	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://api.thecatapi.com/v1/breeds"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating GET request:", err)
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	var breeds []Breed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		log.Println("Error decoding JSON response:", err)
		return
	}

	ch <- breeds
}

// Fetches details of a single breed by its ID
func fetchBreedDetails(apiKey, breedID string, ch chan<- BreedDetails) {
	defer close(ch)

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/%s", breedID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating GET request:", err)
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	var breedDetails BreedDetails
	if err := json.NewDecoder(resp.Body).Decode(&breedDetails); err != nil {
		log.Println("Error decoding JSON response:", err)
		return
	}

	ch <- breedDetails
}

// Fetches images of a breed by its ID
func fetchBreedImages(apiKey, breedID string, ch chan<- []BreedImage) {
	defer close(ch)

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=10", breedID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating GET request:", err)
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	var breedImages []BreedImage
	if err := json.NewDecoder(resp.Body).Decode(&breedImages); err != nil {
		log.Println("Error decoding JSON response:", err)
		return
	}

	ch <- breedImages
}

// Get handles the initial page load
func (c *BreedsController) Get() {
	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")

	// Channels to fetch data concurrently
	breedsCh := make(chan []Breed)
	breedDetailsCh := make(chan BreedDetails)
	breedImagesCh := make(chan []BreedImage)

	var wg sync.WaitGroup
	wg.Add(3)

	// Fetch all breeds
	go func() {
		defer wg.Done()
		fetchAllBreeds(apiKey, breedsCh)
		wg.Wait()
	}()

	// Fetching breeds completed
	breeds := <-breedsCh
	defaultBreedID := breeds[0].ID

	// Fetch breed details and images for the default breed
	go func() {
		defer wg.Done()
		fetchBreedDetails(apiKey, defaultBreedID, breedDetailsCh)
	}()

	go func() {
		defer wg.Done()
		fetchBreedImages(apiKey, defaultBreedID, breedImagesCh)
		wg.Wait()
	}()

	breedDetails := <-breedDetailsCh
	breedImages := <-breedImagesCh

	// Pass data to the template
	c.Data["Breeds"] = breeds
	c.Data["BreedDetails"] = breedDetails
	c.Data["BreedImages"] = breedImages
	c.TplName = "breeds.tpl"
}

// Post handles AJAX requests for breed changes
func (c *BreedsController) Post() {
	apiKey, _ := beego.AppConfig.String("catapiproject.apikey")
	breedID := c.GetString("breed_id")

	// Channels to fetch data concurrently
	breedDetailsCh := make(chan BreedDetails)
	breedImagesCh := make(chan []BreedImage)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fetchBreedDetails(apiKey, breedID, breedDetailsCh)
	}()

	go func() {
		defer wg.Done()
		fetchBreedImages(apiKey, breedID, breedImagesCh)
		wg.Wait()
	}()

	breedDetails := <-breedDetailsCh
	breedImages := <-breedImagesCh

	// Prepare response data
	response := map[string]interface{}{
		"breedDetails": breedDetails,
		"breedImages":  breedImages,
	}

	c.Data["json"] = response
	c.ServeJSON()
}
