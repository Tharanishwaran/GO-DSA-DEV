package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

const apiURL = "https://www.metaweather.com/api/"

type Location struct {
    Woeid int `json:"woeid"`
}

type WeatherResponse struct {
    ConsolidatedWeather []struct {
        WeatherStateName string  `json:"weather_state_name"`
        TheTemp          float64 `json:"the_temp"`
    } `json:"consolidated_weather"`
    Title string `json:"title"`
}

func getLocation(city string) (*Location, error) {
    url := fmt.Sprintf("%s/location/search/?query=%s", apiURL, city)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to get location data: %s", resp.Status)
    }

    var locations []Location
    if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
        return nil, err
    }

    if len(locations) == 0 {
        return nil, fmt.Errorf("location not found")
    }

    return &locations[0], nil
}

func getWeather(woeid int) (*WeatherResponse, error) {
    url := fmt.Sprintf("%s/location/%d/", apiURL, woeid)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
    }

    var weatherResponse WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
        return nil, err
    }

    return &weatherResponse, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: weather-cli <city>")
        os.Exit(1)
    }

    city := os.Args[1]
    location, err := getLocation(city)
    if err != nil {
        log.Fatal(err)
    }

    weather, err := getWeather(location.Woeid)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Weather in %s:\n", weather.Title)
    for _, w := range weather.ConsolidatedWeather {
        fmt.Printf("Date: %s, Temperature: %.2f°C, Weather: %s\n", w.WeatherStateName, w.TheTemp, w.WeatherStateName)
    }
}
