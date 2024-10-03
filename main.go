package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Struct to hold weather information
type Weather struct {
	Current []struct {
		FeelsLikeC  string `json:"FeelsLikeC"`
		Humidity    string `json:"humidity"`
		TempC       string `json:"temp_C"`
		UvIndex     string `json:"uvIndex"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
	} `json:"current_condition"`
}

// Define constants for API URL and default location
const (
	WeatherAPIURL   = "https://wttr.in/%v?format=j1"
	DefaultLocation = "Budapest"
)

func main() {
	// Set default location or get location from command line args
	location := DefaultLocation
	if len(os.Args) >= 2 {
		location = os.Args[1]
	}

	// Fetch weather data
	weather, err := fetchWeather(location)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	// Display the weather data
	displayWeather(location, weather)
}

// fetchWeather makes an HTTP request to fetch weather data for a given location
func fetchWeather(location string) (*Weather, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout for HTTP request
	}

	resp, err := client.Get(fmt.Sprintf(WeatherAPIURL, location))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the API returned a non-OK status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status: %v", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	// Check if the response contains any current weather data
	if len(weather.Current) == 0 {
		return nil, fmt.Errorf("no weather data found for location: %v", location)
	}

	return &weather, nil
}

// displayWeather prints the weather data in a formatted way
func displayWeather(location string, weather *Weather) {
	color.Yellow("Your Weather Location: %s", location)

	// Build weather description using strings.Builder
	var weatherDesc strings.Builder
	for _, wd := range weather.Current[0].WeatherDesc {
		weatherDesc.WriteString(wd.Value + " ")
	}

	// Format and display weather information
	message := fmt.Sprintf(
		"Temperature: %s°C\nHumidity:    %s%%\nFeels Like:  %s°C\nUV Index:    %s\nDescription: %s",
		weather.Current[0].TempC,
		weather.Current[0].Humidity,
		weather.Current[0].FeelsLikeC,
		weather.Current[0].UvIndex,
		strings.TrimSpace(weatherDesc.String()), // Clean up trailing spaces
	)
	color.Green(message)
}
