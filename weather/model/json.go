package model

type CurrentJSON struct {
	City          CityJSON          `json:"city"`
	Temperature   TemperatureJSON   `json:"temperature"`
	Humidity      HumidityJSON      `json:"humidity"`
	Pressure      PressureJSON      `json:"pressure"`
	Wind          WindJSON          `json:"wind"`
	Clouds        CloudsJSON        `json:"clouds"`
	Visibility    VisibilityJSON    `json:"visibility"`
	Precipitation PrecipitationJSON `json:"precipitation"`
	Weather       WeatherJSON       `json:"weather"`
	LastUpdate    LastUpdateJSON    `json:"lastupdate"`
}

type CityJSON struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Coord   CoordJSON `json:"coord"`
	Country string    `json:"country"`
	Sun     SunJSON   `json:"sun"`
}

type CoordJSON struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}

type SunJSON struct {
	Rise string `json:"rise"`
	Set  string `json:"set"`
}

type TemperatureJSON struct {
	Value string `json:"value"`
	Min   string `json:"min"`
	Max   string `json:"max"`
	Unit  string `json:"unit"`
}

type HumidityJSON struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type PressureJSON struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type WindJSON struct {
	Speed     SpeedJSON     `json:"speed"`
	Direction DirectionJSON `json:"direction"`
}

type SpeedJSON struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type DirectionJSON struct {
	Value string `json:"value"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

type CloudsJSON struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type VisibilityJSON struct {
	Value string `json:"value"`
}

type PrecipitationJSON struct {
	Mode string `json:"mode"`
}

type WeatherJSON struct {
	Number string `json:"number"`
	Value  string `json:"value"`
	Icon   string `json:"icon"`
}

type LastUpdateJSON struct {
	Value string `json:"value"`
}
