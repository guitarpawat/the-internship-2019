package model

import "encoding/xml"

type CurrentXML struct {
	XMLName       xml.Name         `xml:"current"`
	City          CityXML          `xml:"city"`
	Temperature   TemperatureXML   `xml:"temperature"`
	Humidity      HumidityXML      `xml:"humidity"`
	Pressure      PressureXML      `xml:"pressure"`
	Wind          WindXML          `xml:"wind"`
	Clouds        CloudsXML        `xml:"clouds"`
	Visibility    VisibilityXML    `xml:"visibility"`
	Precipitation PrecipitationXML `xml:"precipitation"`
	Weather       WeatherXML       `xml:"weather"`
	LastUpdate    LastUpdateXML    `xml:"lastupdate"`
}

type CityXML struct {
	XMLName xml.Name `xml:"city"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Coord   CoordXML `xml:"coord"`
	Country string   `xml:"country"`
	Sun     SunXML   `xml:"sun"`
}

type CoordXML struct {
	XMLName xml.Name `xml:"coord"`
	Lon     string   `xml:"lon,attr"`
	Lat     string   `xml:"lat,attr"`
}

type SunXML struct {
	XMLName xml.Name `xml:"sun"`
	Rise    string   `xml:"rise,attr"`
	Set     string   `xml:"set,attr"`
}

type TemperatureXML struct {
	XMLName xml.Name `xml:"temperature"`
	Value   string   `xml:"value,attr"`
	Min     string   `xml:"min,attr"`
	Max     string   `xml:"max,attr"`
	Unit    string   `xml:"unit,attr"`
}

type HumidityXML struct {
	XMLName xml.Name `xml:"humidity"`
	Value   string   `xml:"value,attr"`
	Unit    string   `xml:"unit,attr"`
}

type PressureXML struct {
	XMLName xml.Name `xml:"pressure"`
	Value   string   `xml:"value,attr"`
	Unit    string   `xml:"unit,attr"`
}

type WindXML struct {
	XMLName   xml.Name     `xml:"wind"`
	Speed     SpeedXML     `xml:"speed"`
	Direction DirectionXML `xml:"direction"`
}

type SpeedXML struct {
	XMLName xml.Name `xml:"speed"`
	Value   string   `xml:"value,attr"`
	Name    string   `xml:"name,attr"`
}

type DirectionXML struct {
	XMLName xml.Name `xml:"direction"`
	Value   string   `xml:"value,attr"`
	Code    string   `xml:"code,attr"`
	Name    string   `xml:"name,attr"`
}

type CloudsXML struct {
	XMLName xml.Name `xml:"clouds"`
	Value   string   `xml:"value,attr"`
	Name    string   `xml:"name,attr"`
}

type VisibilityXML struct {
	XMLName xml.Name `xml:"visibility"`
	Value   string   `xml:"value,attr"`
}

type PrecipitationXML struct {
	XMLName xml.Name `xml:"precipitation"`
	Mode    string   `xml:"mode,attr"`
}

type WeatherXML struct {
	XMLName xml.Name `xml:"weather"`
	Number  string   `xml:"number,attr"`
	Value   string   `xml:"value,attr"`
	Icon    string   `xml:"icon,attr"`
}

type LastUpdateXML struct {
	XMLName xml.Name `xml:"lastupdate"`
	Value   string   `xml:"value,attr"`
}

func (c CurrentXML) ToJSON() CurrentJSON {
	return CurrentJSON{
		City: CityJSON{
			ID:   c.City.ID,
			Name: c.City.Name,
			Coord: CoordJSON{
				Lon: c.City.Coord.Lon,
				Lat: c.City.Coord.Lat,
			},
			Country: c.City.Country,
			Sun: SunJSON{
				Rise: c.City.Sun.Rise,
				Set:  c.City.Sun.Set,
			},
		},
		Temperature: TemperatureJSON{
			Value: c.Temperature.Value,
			Min:   c.Temperature.Min,
			Max:   c.Temperature.Max,
			Unit:  c.Temperature.Unit,
		},
		Humidity: HumidityJSON{
			Value: c.Humidity.Value,
			Unit:  c.Humidity.Unit,
		},
		Pressure: PressureJSON{
			Value: c.Pressure.Value,
			Unit:  c.Pressure.Unit,
		},
		Wind: WindJSON{
			Speed: SpeedJSON{
				Value: c.Wind.Speed.Value,
				Name:  c.Wind.Speed.Name,
			},
			Direction: DirectionJSON{
				Value: c.Wind.Direction.Value,
				Code:  c.Wind.Direction.Code,
				Name:  c.Wind.Direction.Name,
			},
		},
		Clouds: CloudsJSON{
			Value: c.Clouds.Value,
			Name:  c.Clouds.Name,
		},
		Visibility: VisibilityJSON{
			Value: c.Visibility.Value,
		},
		Precipitation: PrecipitationJSON{
			Mode: c.Precipitation.Mode,
		},
		Weather: WeatherJSON{
			Number: c.Weather.Number,
			Value:  c.Weather.Value,
			Icon:   c.Weather.Icon,
		},
		LastUpdate: LastUpdateJSON{
			Value: c.LastUpdate.Value,
		},
	}
}
