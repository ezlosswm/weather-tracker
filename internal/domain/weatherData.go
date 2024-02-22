package domain

type Location struct {
	Name    string `json:"name"`
	PlaceID string `json:"place_id"`

	// options
	// metric: Celsius
	// us: Farenheight
	Unit    string `json:"unit"`
	Country string `json:"country"`
}

type Weather struct {
	Current Current `json:"current"`
	Daily   Daily   `json:"daily"`
}

type Current struct {
	Summary     string  `json:"summary"`
	Temperature float32 `json:"temperature"`
}

type Daily struct {
	Data []Data `json:"data"`
}

type Data struct {
	Day     string `json:"day"`
	Weather string `json:"weather"`
	Summary string `json:"summary"`
}

type AllDay struct {
	Temperature    float32 `json:"temperature"`
	TemperatureMin float32 `json:"temperature_min"`
	TemperatureMax float32 `json:"temperature_max"`
}
