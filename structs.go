package main

/* ----------------------------------------------------------------------- */
/*                       Response set of structures                        */
/* ----------------------------------------------------------------------- */
type weatherdiff struct {
	Supplier        string
	Country         string
	City            string
	MeasureDateTime string
	Data            weatherdiffdata
}

type weatherdiffdatacoord struct {
	Latitude  float32
	Longitude float32
}

type weatherdiffdatawind struct {
	Direction float32
	Speed     float32
}

type weatherdiffdata struct {
	Temperature float32
	Pressure    float32
	Humidity    float32
	Description string
	Coordinates weatherdiffdatacoord
	Wind        weatherdiffdatawind
}

type weatherresponserror struct {
	Status      string
	Description string
}

/* ----------------------------------------------------------------------- */
/*                     AccuWeather set of structures                       */
/* ----------------------------------------------------------------------- */
type accuweathercoord struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}

type accuweathercontry struct {
	Code string `json:"ID"`
}

type accuweatherpos struct {
	CityName    string            `json:"EnglishName"`
	Coordinates accuweathercoord  `json:"GeoPosition"`
	Country     accuweathercontry `json:"Country"`
	UID         string            `json:"Key"`
}

type accuweathermetric struct {
	Metric map[string]interface{} `json:"Metric"`
}

type accuweatherwind struct {
	Direction map[string]interface{} `json:"Direction"`
	Speed     accuweathermetric      `json:"Speed"`
}

type accuweatherdata struct {
	MeasureDateTime int64             `json:"EpochTime"`
	Temperature     accuweathermetric `json:"Temperature"`
	Pressure        accuweathermetric `json:"Pressure"`
	Wind            accuweatherwind   `json:"Wind"`
	Humidity        float32           `json:"RelativeHumidity"`
	Description     string            `json:"WeatherText"`
}

/* ----------------------------------------------------------------------- */
/*                      WeatherBit set of structures                       */
/* ----------------------------------------------------------------------- */
type weatherbitdescr struct {
	Clouds string `json:"description"`
}

type weatherbitdata struct {
	Temperature     float32         `json:"temp"`
	CityName        string          `json:"city_name"`
	Pressure        float32         `json:"pres"`
	WindDir         uint16          `json:"wind_dir"`
	WindSpeed       float32         `json:"wind_spd"`
	Latitude        float32         `json:"lat"`
	Longitude       float32         `json:"lon"`
	MeasureDateTime int64           `json:"ts"`
	Description     weatherbitdescr `json:"weather"`
	Country         string          `json:"country_code"`
	Humidity        float32         `json:"rh"`
}

type weatherbit struct {
	Data  []weatherbitdata `json:"data"`
	Count uint8            `json:"count"`
}

/* ----------------------------------------------------------------------- */
/*                    OpenWeatherMap set of structures                     */
/* ----------------------------------------------------------------------- */
type openweathermapcoord struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type openweathermapmain struct {
	Temperature float32 `json:"temp"`
	Pressure    uint16  `json:"pressure"`
	Humidity    uint8   `json:"humidity"`
}

type openweathermapweather struct {
	Clouds string `json:"description"`
}

type openweathermapsys struct {
	Country string `json:"country"`
}

type openweathermapwind struct {
	Direction uint16  `json:"deg"`
	Speed     float32 `json:"speed"`
}

type openweathermap struct {
	// Data  []openweathermapdata ``
	CityName        string                  `json:"name"`
	Sys             openweathermapsys       `json:"sys"`
	Coordinates     openweathermapcoord     `json:"coord"`
	Main            openweathermapmain      `json:"main"`
	Description     []openweathermapweather `json:"weather"`
	MeasureDateTime int64                   `json:"dt"`
	Wind            openweathermapwind      `json:"wind"`
}
