package main

import (
	"encoding/json"
	"fmt"
	"github.com/guitarpawat/the-internship-2019/weather/model"
)

var example = `{
  "city": {
    "id": "2643741",
    "name": "City of London",
    "coord": {
      "lon": "-0.09",
      "lat": "51.51"
    },
    "country": "GB",
    "sun": {
      "rise": "2015-06-30T03:46:57",
      "set": "2015-06-30T20:21:12"
    }
  },
  "temperature": {
    "value": "72.34",
    "min": "66.2",
    "max": "79.88",
    "unit": "fahrenheit"
  },
  "humidity": {
    "value": "43",
    "unit": "%"
  },
  "pressure": {
    "value": "1020",
    "unit": "hPa"
  },
  "wind": {
    "speed": {
      "value": "7.78",
      "name": "Moderate breeze"
    },
    "direction": {
      "value": "140",
      "code": "SE",
      "name": "SouthEast"
    }
  },
  "clouds": {
    "value": "0",
    "name": "clear sky"
  },
  "visibility": {
    "value": "10000"
  },
  "precipitation": {
    "mode": "no"
  },
  "weather": {
    "number": "800",
    "value": "Sky is Clear",
    "icon": "01d"
  },
  "lastupdate": {
    "value": "2015-06-30T08:36:14"
  }
}`

func main() {
	weather := model.CurrentJSON{}
	json.Unmarshal([]byte(example), &weather)
	fmt.Printf("%+v\n", weather)
	//weather := model.WeatherXML{City: model.CityXML{
	//	ID:   "1",
	//	Name: "Bangkok",
	//	//Lon:     "-1",
	//	//Lat:     "-2",
	//	Country: "TH",
	//}}
	//raw, err := xml.Marshal(weather)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(string(raw))
}
