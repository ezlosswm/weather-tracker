package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"weather-tracker/internal/domain"
)

const (
	BASE_URL = "https://www.meteosource.com/api/v1/free/"
	LANGUAGE = "language=en&"
)

var (
	// API
	API_KEY = os.Getenv("API_KEY")
	KEY     = fmt.Sprintf("key=%s", API_KEY)
)

// FetchLocations formulates the URL from the input value then
// makes a GET request.
func FetchLocations(input string) ([]domain.Location, error) {
	locationURL := BASE_URL + fmt.Sprintf("find_places?text=%s&", input) + LANGUAGE + KEY

	req, err := http.NewRequest(http.MethodGet, locationURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	weatherResp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal response
	var locations []domain.Location
	if err := json.Unmarshal(weatherResp, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func FetchWeather(input string) (domain.Weather, error) {
	// ex := `https://www.meteosource.com/api/v1/free/point?place_id=belize-city&sections=current%2Cdaily&language=en&units=us&key=woxhy321n51o3k8zjt0d0kwsco7rqvyb30hohst3`
	var (
		weather domain.Weather
		query   = fmt.Sprintf("point?place_id=%s&", input)
		key     = fmt.Sprintf("key=%s", API_KEY)
		etc     = "sections=current%2Cdaily&language=en&units=us&"
	)

	URL := BASE_URL + query + etc + key
	req, _ := http.NewRequest(http.MethodGet, URL, nil)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return weather, err
	}

	weatherResp, err := io.ReadAll(response.Body)
	if err != nil {
		return weather, err
	}

	// unmarshal response
	if err := json.Unmarshal(weatherResp, &weather); err != nil {
		var w domain.Weather
		return w, err
	}

	return weather, nil
}
