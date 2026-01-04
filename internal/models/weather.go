package models

type Weather struct {
	Location    string `json:"location"`
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}

var WeatherData = map[string]Weather{
	"frauenfeld": {Location: "Frauenfeld", Condition: "cloudy", Temperature: "15°"},
	"miami":      {Location: "Miami", Condition: "sunny", Temperature: "29°"},
	"bangkok":    {Location: "Bangkok", Condition: "sunny", Temperature: "35°"},
}
