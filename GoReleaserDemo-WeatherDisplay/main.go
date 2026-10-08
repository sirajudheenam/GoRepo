package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

// Struct for parsing weather data
type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func main() {
	cities := []City{
		{"London", 51.5074, -0.1278},
		{"New York", 40.7128, -74.0060},
		{"Tokyo", 35.6762, 139.6503},
		{"Paris", 48.8566, 2.3522},
		{"Berlin", 52.5200, 13.4050},
	}

	header := color.New(color.FgCyan, color.Bold)
	header.Println("\n🌤  Top 5 Cities Weather Report\n")

	for _, city := range cities {
		printCityWeather(city)
		time.Sleep(500 * time.Millisecond) // smooth delay between fetches
	}

	color.New(color.FgGreen, color.Bold).Println("✅ Done.\n")
}

func printCityWeather(city City) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		city.Latitude, city.Longitude,
	)

	resp, err := http.Get(url)
	if err != nil {
		color.Red("Error fetching weather for %s: %v\n", city.Name, err)
		return
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		color.Red("Error decoding response for %s: %v\n", city.Name, err)
		return
	}

	cityTitle := color.New(color.FgYellow, color.Bold)
	tempText := color.New(color.FgHiMagenta)
	windText := color.New(color.FgCyan)
	timeText := color.New(color.FgWhite)

	cityTitle.Printf("🌆 %s\n", city.Name)
	tempText.Printf("   🌡 Temperature: %.1f°C\n", weather.CurrentWeather.Temperature)
	windText.Printf("   💨 Wind Speed: %.1f km/h\n", weather.CurrentWeather.Windspeed)
	timeText.Printf("   🕒 Time: %s\n\n", formatTime(weather.CurrentWeather.Time))
}

func formatTime(t string) string {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t
	}
	return parsed.Local().Format("Mon Jan 2 15:04")
}
