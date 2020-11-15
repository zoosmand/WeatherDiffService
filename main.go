package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"strings"
)

const listenPort = "8080"

/* --------------------------------------------------------------- */

func main() {

	http.HandleFunc("/", WeatherServiceHandler)

	RemoteSyslog("Starting server at port "+listenPort, syslog.LOG_INFO)

	if runServerError := http.ListenAndServe(":"+listenPort, nil); runServerError != nil {
		RemoteSyslog(runServerError.Error(), syslog.LOG_ERR)
		fmt.Println(runServerError.Error())
		os.Exit(-1)
	}
}

// WeatherServiceHandler handlers ....
func WeatherServiceHandler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	/* Only GET allowed */
	if httpRequest.Method != "GET" {
		http.Error(httpResponse, "Only GET method is supported here.", http.StatusNotAcceptable)
		return
	}
	// Parse parameter 'country'
	requestedCountry, ok := httpRequest.URL.Query()["country"]
	if !ok || len(requestedCountry[0]) < 1 {
		log.Println("Url Param 'country' is missing")
		errorResponse := weatherresponserror{
			Status:      "Error",
			Description: "Url Param 'country' is missing.",
		}
		httpResponse.Header().Set("Zoosman-Result", "ERROR")
		WeatherServiceResponseHandler(httpResponse, errorResponse)
		return
	}
	// Parse parameter 'city'
	requestedCity, ok := httpRequest.URL.Query()["city"]
	if !ok || len(requestedCity[0]) < 1 {
		log.Println("Url Param 'city' is missing")
		errorResponse := weatherresponserror{
			Status:      "Error",
			Description: "Url Param 'city' is missing.",
		}
		httpResponse.Header().Set("Zoosman-Result", "ERROR")
		WeatherServiceResponseHandler(httpResponse, errorResponse)
		return
	}

	weatherSources := [3]interface{}{
		AccuweatherReply,
		WeatherbitReply,
		OpenweathermapReply,
	}

	var resp []weatherdiff

	for _, weatherSource := range weatherSources {
		weatherSource.(func(string, string))(strings.ToUpper(requestedCity[0]), strings.ToUpper(requestedCountry[0]))
		resp = append(resp, GetWeatherData())
	}

	httpResponse.Header().Set("Zoosman-Result", "OK")
	WeatherServiceResponseHandler(httpResponse, resp)
}

// WeatherServiceResponseHandler handlers ....
func WeatherServiceResponseHandler(httpResponse http.ResponseWriter, responseData interface{}) {
	httpResponse.Header().Set("Content-Type", "application/json")

	jsonResponseData, marshallError := json.Marshal(responseData)
	if marshallError != nil {
		log.Println(marshallError)
	}

	httpResponse.Write(jsonResponseData)
}
