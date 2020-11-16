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

// The TCP Port to listen incoming requests
const listenPort = "8080"

/* --------------------------------------------------------------- */

func main() {
	// Calls HTTP handler callback using "root" of "web-server"
	http.HandleFunc("/", WeatherServiceHandler)
	// Initiates transfer log data to a Syslog server. It is quite enough to call
	// RemoteSyslog() once so as to use it further. Though in case of changing
	// severity level is happened, followed info uses the same severity until
	// RemoteSyslog() wouldn't have called again with the new severity.
	RemoteSyslog("Starting server at port "+listenPort, syslog.LOG_INFO)

	// Runs HTTP server to listen a port saved in listenPort
	if runServerError := http.ListenAndServe(":"+listenPort, nil); runServerError != nil {
		// Write syslog an error end close the application
		RemoteSyslog(runServerError.Error(), syslog.LOG_ERR)
		fmt.Println(runServerError.Error())
		os.Exit(-1)
	}
}

// WeatherServiceHandler allows the GET method only. Requests which use other methods will be denied.
// To use WeatherDiffService, the request should be properly combined. The service expects two incoming
// parameters "country" and "city". "country" parameter should contain two symbols "GB" or "us"
// (case insensitive) "city" is the city name like "london" or "Moscow" (case insensitive).
// @httpResponse: HTTP response
// @httpRequest: a pointer to HTTP request
func WeatherServiceHandler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	// Only the GET method allowed
	if httpRequest.Method != "GET" {
		http.Error(httpResponse, "Only GET method is supported here.", http.StatusNotAcceptable)
		return
	}
	// Parse parameter 'country'
	requestedCountry, ok := httpRequest.URL.Query()["country"]
	if !ok || len(requestedCountry[0]) < 1 {
		// in case the parameter is missing, write a log message
		// then send the resulting error marshaled in a JSON.
		log.Println("Url Param 'country' is missing")
		errorResponse := weatherresponserror{
			Status:      "Error",
			Description: "Url Param 'country' is missing.",
		}
		// Send special header info. Usually is uses for early decision.
		httpResponse.Header().Set("Zoosman-Result", "ERROR")
		// Call HTTP response rutine to send error data
		WeatherServiceResponseHandler(httpResponse, errorResponse)
		return
	}
	// Parse parameter 'city'
	requestedCity, ok := httpRequest.URL.Query()["city"]
	if !ok || len(requestedCity[0]) < 1 {
		// in case the parameter is missing, write a log message
		// then send the resulting error marshaled in a JSON.
		log.Println("Url Param 'city' is missing")
		errorResponse := weatherresponserror{
			Status:      "Error",
			Description: "Url Param 'city' is missing.",
		}
		// Send special header info. Usually is uses for early decision.
		httpResponse.Header().Set("Zoosman-Result", "ERROR")
		// Call HTTP response rutine to send error data
		WeatherServiceResponseHandler(httpResponse, errorResponse)
		return
	}
	// Collect all weather suppliers methods together. It helps to shrink the code.
	weatherSources := [3]interface{}{
		AccuweatherReply,
		WeatherbitReply,
		OpenweathermapReply,
	}
	// State the result data as an array of "weatherdiff" struct
	var resp []weatherdiff
	// Collect data into resp array
	for _, weatherSource := range weatherSources {
		// ToDo. I'd better cut country to two symbols and check it somewhere before calling a supplier
		// ToDo.  City shoud be checked somewhere, otherwise is have to call an exception
		weatherSource.(func(string, string))(strings.ToUpper(requestedCity[0]), strings.ToUpper(requestedCountry[0]))
		resp = append(resp, GetWeatherData())
	}
	// Send special header info. Usually is uses for early decision.
	httpResponse.Header().Set("Zoosman-Result", "OK")
	// Call HTTP response rutine to send error data
	WeatherServiceResponseHandler(httpResponse, resp)
}

// WeatherServiceResponseHandler responses directly carring marshalled data
// @httpResponse: a response writer
// @responseData: data to be sent
func WeatherServiceResponseHandler(httpResponse http.ResponseWriter, responseData interface{}) {
	// Sends a header indicating that JSON data are being sent
	httpResponse.Header().Set("Content-Type", "application/json")
	// Wraps response data into a byte buffer
	jsonResponseData, marshallError := json.Marshal(responseData)
	// In case when something wrong is happend, it writes a log message
	if marshallError != nil {
		log.Println(marshallError)
	}
	// Sends the buffer
	httpResponse.Write(jsonResponseData)
}
