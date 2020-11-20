package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

const dateTimeFormat = "2006/01/02 15:04"

// A "static" variable to save collected data
var weatherData weatherdiff

const accuweatherKey = "xxx"
const weatherbitKey = "xxx"
const opeweathermapKey = "xxx"

/* *********************************************************************** */
/*                          Weather service methodes                       */
/* *********************************************************************** */

// GetWeatherData returns collected weather data
func GetWeatherData() weatherdiff {
	return weatherData
}

// AccuweatherReply collect theirs data into "weatherData"
// @city: 			a city name to search
// @countryCode: 	a countryCode to search
func AccuweatherReply(city string, countryCode string) {
	// First, a location must be sought. A unique ID, being received,
	// is needed to obtain the weather data in the next request.
	// locationRaw := WeatherJSONHandler("http://crm.asart.local:8010/accuweather_location.json")
	locationRaw := WeatherJSONHandler("http://dataservice.accuweather.com/locations/v1/cities/search?apikey=" + accuweatherKey + "&q=" + city + "," + countryCode)
	// An array to store a location
	var dataSet []accuweatherpos

	// Gets and unmarshalls data into an array of "accuweatherpos"
	jsonError := json.Unmarshal(locationRaw, &dataSet)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	// Formats an item to a string
	countryStr := fmt.Sprintf("%v", dataSet[0].Country.Code)
	// Formats an item to a string
	cityStr := fmt.Sprintf("%v", dataSet[0].CityName)
	// Formats an item to a string
	LatitudeStr := fmt.Sprintf("%v", dataSet[0].Coordinates.Latitude)
	// Formats an item to a float
	LatitudeFl, conversionError := strconv.ParseFloat(LatitudeStr, 32)
	// Checks if the conversion is right
	if conversionError != nil {
		log.Println(conversionError)
	}

	// Formats an item to a string
	LongitudeStr := fmt.Sprintf("%v", dataSet[0].Coordinates.Longitude)
	// Formats an item to a float
	LongitudeFl, conversionError := strconv.ParseFloat(LongitudeStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	// Formats an item to a string
	// It is a checkpoint value that indicates a unique ID of a city to be sought
	// Then it's value to be used the weather data request
	uid := fmt.Sprintf("%v", dataSet[0].UID)

	// Second, when a UID has obtained, the weather data could be requested
	// weatherRaw := WeatherJSONHandler("http://crm.asart.local:8010/accuweather_data.json?q=" + uid)
	weatherRaw := WeatherJSONHandler("http://dataservice.accuweather.com/currentconditions/v1/" + uid + "?apikey=" + accuweatherKey + "&details=true")
	// An array to store weather data
	var weatherSet []accuweatherdata

	// Gets and unmarshalls data into an array of "accuweatherdata"
	jsonError = json.Unmarshal(weatherRaw, &weatherSet)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	temperatureStr := fmt.Sprintf("%v", weatherSet[0].Temperature.Metric["Value"])
	temperatureFl, conversionError := strconv.ParseFloat(temperatureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	pressureStr := fmt.Sprintf("%v", weatherSet[0].Pressure.Metric["Value"])
	pressureFl, conversionError := strconv.ParseFloat(pressureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	humidityStr := fmt.Sprintf("%v", weatherSet[0].Humidity)
	humidityFl, conversionError := strconv.ParseFloat(humidityStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	DirectionStr := fmt.Sprintf("%v", weatherSet[0].Wind.Direction["Degrees"])
	DirectionFl, conversionError := strconv.ParseFloat(DirectionStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	SpeedStr := fmt.Sprintf("%v", weatherSet[0].Wind.Speed.Metric["Value"])
	SpeedFl, conversionError := strconv.ParseFloat(SpeedStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	weatherData = weatherdiff{
		Supplier:        "Accuweather",
		Country:         countryStr,
		City:            cityStr,
		MeasureDateTime: fmt.Sprintf("%v", time.Unix(weatherSet[0].MeasureDateTime, 0).Format(dateTimeFormat)),
		Data: weatherdiffdata{
			Temperature: float32(temperatureFl),
			Pressure:    float32(pressureFl),
			Humidity:    float32(humidityFl),
			Description: fmt.Sprintf("%v", weatherSet[0].Description),
			Coordinates: weatherdiffdatacoord{
				Latitude:  float32(LatitudeFl),
				Longitude: float32(LongitudeFl),
			},
			Wind: weatherdiffdatawind{
				Direction: float32(DirectionFl),
				Speed:     float32(math.Round(SpeedFl * 1000 / 3600)),
			},
		},
	}
}

// WeatherbitReply returns a map with weather data
func WeatherbitReply(city string, countryCode string) {
	// weatherRaw := WeatherJSONHandler("http://crm.asart.local:8010/weatherbit_data.json")
	weatherRaw := WeatherJSONHandler("http://api.weatherbit.io/v2.0/current?city=" + city + "&country=" + countryCode + "&key=" + weatherbitKey)
	weatherSet := weatherbit{}

	jsonError := json.Unmarshal(weatherRaw, &weatherSet)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	temperatureStr := fmt.Sprintf("%v", weatherSet.Data[0].Temperature)
	temperatureFl, conversionError := strconv.ParseFloat(temperatureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	pressureStr := fmt.Sprintf("%v", weatherSet.Data[0].Pressure)
	pressureFl, conversionError := strconv.ParseFloat(pressureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	humidityStr := fmt.Sprintf("%v", weatherSet.Data[0].Humidity)
	humidityFl, conversionError := strconv.ParseFloat(humidityStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	LatitudeStr := fmt.Sprintf("%v", weatherSet.Data[0].Latitude)
	LatitudeFl, conversionError := strconv.ParseFloat(LatitudeStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	LongitudeStr := fmt.Sprintf("%v", weatherSet.Data[0].Longitude)
	LongitudeFl, conversionError := strconv.ParseFloat(LongitudeStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	DirectionStr := fmt.Sprintf("%v", weatherSet.Data[0].WindDir)
	DirectionFl, conversionError := strconv.ParseFloat(DirectionStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	SpeedStr := fmt.Sprintf("%v", weatherSet.Data[0].WindSpeed)
	SpeedFl, conversionError := strconv.ParseFloat(SpeedStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	weatherData = weatherdiff{
		Supplier:        "Weatherbit",
		Country:         fmt.Sprintf("%v", weatherSet.Data[0].Country),
		City:            fmt.Sprintf("%v", weatherSet.Data[0].CityName),
		MeasureDateTime: fmt.Sprintf("%v", time.Unix(weatherSet.Data[0].MeasureDateTime, 0).Format(dateTimeFormat)),
		Data: weatherdiffdata{
			Temperature: float32(temperatureFl),
			Pressure:    float32(pressureFl),
			Humidity:    float32(humidityFl),
			Description: fmt.Sprintf("%v", weatherSet.Data[0].Description.Clouds),
			Coordinates: weatherdiffdatacoord{
				Latitude:  float32(LatitudeFl),
				Longitude: float32(LongitudeFl),
			},
			Wind: weatherdiffdatawind{
				Direction: float32(DirectionFl),
				Speed:     float32(SpeedFl),
			},
		},
	}

}

// OpenweathermapReply returns a map with weather data
func OpenweathermapReply(city string, countryCode string) {

	// weatherRaw := WeatherJSONHandler("http://crm.asart.local:8010/openweathermap_data.json")
	weatherRaw := WeatherJSONHandler("http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + countryCode + "&appid=" + opeweathermapKey)
	weatherSet := openweathermap{}

	jsonError := json.Unmarshal(weatherRaw, &weatherSet)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	temperatureStr := fmt.Sprintf("%v", weatherSet.Main.Temperature)
	temperatureFl, conversionError := strconv.ParseFloat(temperatureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	pressureStr := fmt.Sprintf("%v", weatherSet.Main.Pressure)
	pressureFl, conversionError := strconv.ParseFloat(pressureStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	humidityStr := fmt.Sprintf("%v", weatherSet.Main.Humidity)
	humidityFl, conversionError := strconv.ParseFloat(humidityStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	LatitudeStr := fmt.Sprintf("%v", weatherSet.Coordinates.Latitude)
	LatitudeFl, conversionError := strconv.ParseFloat(LatitudeStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	LongitudeStr := fmt.Sprintf("%v", weatherSet.Coordinates.Longitude)
	LongitudeFl, conversionError := strconv.ParseFloat(LongitudeStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	DirectionStr := fmt.Sprintf("%v", weatherSet.Wind.Direction)
	DirectionFl, conversionError := strconv.ParseFloat(DirectionStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	SpeedStr := fmt.Sprintf("%v", weatherSet.Wind.Speed)
	SpeedFl, conversionError := strconv.ParseFloat(SpeedStr, 32)
	if conversionError != nil {
		log.Println(conversionError)
	}

	weatherData = weatherdiff{
		Supplier:        "Openweathermap",
		Country:         fmt.Sprintf("%v", weatherSet.Sys.Country),
		City:            fmt.Sprintf("%v", weatherSet.CityName),
		MeasureDateTime: fmt.Sprintf("%v", time.Unix(weatherSet.MeasureDateTime, 0).Format(dateTimeFormat)),
		Data: weatherdiffdata{
			Temperature: float32(math.Round(temperatureFl - 273.15)),
			Pressure:    float32(pressureFl),
			Humidity:    float32(humidityFl),
			Description: fmt.Sprintf("%v", weatherSet.Description[0].Clouds),
			Coordinates: weatherdiffdatacoord{
				Latitude:  float32(LatitudeFl),
				Longitude: float32(LongitudeFl),
			},
			Wind: weatherdiffdatawind{
				Direction: float32(DirectionFl),
				Speed:     float32(SpeedFl),
			},
		},
	}

}

// WeatherJSONHandler return HTTP JSON content
var WeatherJSONHandler = func(url string) []byte {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	httpRequest, httpRequestError := http.NewRequest(http.MethodGet, url, nil)
	if httpRequestError != nil {
		log.Fatal(httpRequestError)
	}

	httpRequest.Header.Set("User-Agent", "Zoosman-WebApp_v000.1")

	httpResponse, httpResponseError := client.Do(httpRequest)
	if httpResponseError != nil {
		log.Fatal(httpResponseError)
	}

	if httpResponse.Body != nil {
		defer httpResponse.Body.Close()
	}

	httpContent, httpContentError := ioutil.ReadAll(httpResponse.Body)
	if httpContentError != nil {
		log.Fatal(httpContentError)
	}

	return httpContent
}
