package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoLocation struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempC      float64 `json:"temp_c"`
	FeelslikeC float64 `json:"feelslike_c"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type Forecastday struct {
	Hour []Hour `json:"hour"`
}

type Hour struct {
	Time         string  `json:"time"`
	TempC        float64 `json:"temp_c"`
	FeelslikeC   float64 `json:"feelslike_c"`
	ChanceOfRain int     `json:"chance_of_rain"`
	ChanceOfSnow int     `json:"chance_of_snow"`
}

var apiKey string = "170fda74e0df45a59c6225221231709"

func GetGeoLocation(l string) Weather {
	d := Weather{}
	// g := getPlublicIpDetails()

	weatherApiUrl := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, l)

	req := client("GET", weatherApiUrl)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	unmarshalErr := json.Unmarshal(body, &d)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return d
}

func client(m string, u string) *http.Request {

	r, err := http.NewRequest(m, u, nil)

	if err != nil {
		panic(err)
	}

	return r
}

func getPlublicIpDetails() GeoLocation {

	g := GeoLocation{}

	req := client("GET", "https://ipinfo.io/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	err = json.Unmarshal(body, &g)

	if err != nil {
		panic(err)
	}

	return g
}
