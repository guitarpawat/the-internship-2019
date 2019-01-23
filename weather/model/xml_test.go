package model

import "testing"

var input = CurrentXML{
	City: CityXML{
		ID:   "2643741",
		Name: "City of London",
		Coord: CoordXML{
			Lon: "-0.09",
			Lat: "51.51",
		},
		Country: "GB",
		Sun: SunXML{
			Rise: "2015-06-30T03:46:57",
			Set:  "2015-06-30T20:21:12",
		},
	},
	Temperature: TemperatureXML{
		Value: "72.34",
		Min:   "66.2",
		Max:   "79.88",
		Unit:  "fahrenheit",
	},
	Humidity: HumidityXML{
		Value: "43",
		Unit:  "%",
	},
	Pressure: PressureXML{
		Value: "1020",
		Unit:  "hPa",
	},
	Wind: WindXML{
		Speed: SpeedXML{
			Value: "7.78",
			Name:  "Moderate breeze",
		},
		Direction: DirectionXML{
			Value: "140",
			Code:  "SE",
			Name:  "SouthEast",
		},
	},
	Clouds: CloudsXML{
		Value: "0",
		Name:  "clear sky",
	},
	Visibility: VisibilityXML{
		Value: "10000",
	},
	Precipitation: PrecipitationXML{
		Mode: "no",
	},
	Weather: WeatherXML{
		Number: "800",
		Value:  "Sky is Clear",
		Icon:   "01d",
	},
	LastUpdate: LastUpdateXML{
		Value: "2015-06-30T08:36:14",
	},
}

var expected = CurrentJSON{
	City: CityJSON{
		ID:   "2643741",
		Name: "City of London",
		Coord: CoordJSON{
			Lon: "-0.09",
			Lat: "51.51",
		},
		Country: "GB",
		Sun: SunJSON{
			Rise: "2015-06-30T03:46:57",
			Set:  "2015-06-30T20:21:12",
		},
	},
	Temperature: TemperatureJSON{
		Value: "72.34",
		Min:   "66.2",
		Max:   "79.88",
		Unit:  "fahrenheit",
	},
	Humidity: HumidityJSON{
		Value: "43",
		Unit:  "%",
	},
	Pressure: PressureJSON{
		Value: "1020",
		Unit:  "hPa",
	},
	Wind: WindJSON{
		Speed: SpeedJSON{
			Value: "7.78",
			Name:  "Moderate breeze",
		},
		Direction: DirectionJSON{
			Value: "140",
			Code:  "SE",
			Name:  "SouthEast",
		},
	},
	Clouds: CloudsJSON{
		Value: "0",
		Name:  "clear sky",
	},
	Visibility: VisibilityJSON{
		Value: "10000",
	},
	Precipitation: PrecipitationJSON{
		Mode: "no",
	},
	Weather: WeatherJSON{
		Number: "800",
		Value:  "Sky is Clear",
		Icon:   "01d",
	},
	LastUpdate: LastUpdateJSON{
		Value: "2015-06-30T08:36:14",
	},
}

func testHelper(t *testing.T, fieldName, expected, output string) {
	if output != expected {
		t.Errorf("expected field `%s` to be `%s`, but get `%s`", fieldName, expected, output)
	}
}

func TestCurrentXML_ToJSON(t *testing.T) {
	output := input.ToJSON()

	testHelper(t, "city>id", expected.City.ID, output.City.ID)
	testHelper(t, "city>name", expected.City.Name, output.City.Name)
	testHelper(t, "city>coord>lon", expected.City.Coord.Lon, output.City.Coord.Lon)
	testHelper(t, "city>coord>lat", expected.City.Coord.Lat, output.City.Coord.Lat)
	testHelper(t, "city>country", expected.City.Country, output.City.Country)
	testHelper(t, "city>sun>rise", expected.City.Sun.Rise, output.City.Sun.Rise)
	testHelper(t, "city>sun>set", expected.City.Sun.Set, output.City.Sun.Set)
	testHelper(t, "temperature>value", expected.Temperature.Value, output.Temperature.Value)
	testHelper(t, "temperature>min", expected.Temperature.Min, output.Temperature.Min)
	testHelper(t, "temperature>max", expected.Temperature.Max, output.Temperature.Max)
	testHelper(t, "temperature>unit", expected.Temperature.Unit, output.Temperature.Unit)
	testHelper(t, "humidity>value", expected.Humidity.Value, output.Humidity.Value)
	testHelper(t, "humidity>unit", expected.Humidity.Unit, output.Humidity.Unit)
	testHelper(t, "pressure>value", expected.Pressure.Value, output.Pressure.Value)
	testHelper(t, "pressure>unit", expected.Pressure.Unit, output.Pressure.Unit)
	testHelper(t, "wind>speed>value", expected.Wind.Speed.Value, output.Wind.Speed.Value)
	testHelper(t, "wind>speed>name", expected.Wind.Speed.Name, output.Wind.Speed.Name)
	testHelper(t, "wind>direction>value", expected.Wind.Direction.Value, output.Wind.Direction.Value)
	testHelper(t, "wind>direction>code", expected.Wind.Direction.Code, output.Wind.Direction.Code)
	testHelper(t, "wind>direction>name", expected.Wind.Direction.Name, output.Wind.Direction.Name)
	testHelper(t, "clouds>value", expected.Clouds.Value, output.Clouds.Value)
	testHelper(t, "clouds>name", expected.Clouds.Name, output.Clouds.Name)
	testHelper(t, "visibility>value", expected.Visibility.Value, output.Visibility.Value)
	testHelper(t, "precipitation>mode", expected.Precipitation.Mode, output.Precipitation.Mode)
	testHelper(t, "weather>number", expected.Weather.Number, output.Weather.Number)
	testHelper(t, "weather>value", expected.Weather.Value, output.Weather.Value)
	testHelper(t, "weather>icon", expected.Weather.Icon, output.Weather.Icon)
	testHelper(t, "lastupdate>value", expected.LastUpdate.Value, output.LastUpdate.Value)
}
