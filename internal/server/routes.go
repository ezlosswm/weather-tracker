package server

import (
	"context"
	"log"
	"net/http"
	"weather-tracker/internal/domain"
	"weather-tracker/internal/repository"
	"weather-tracker/views"
	"weather-tracker/views/partials"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.handler)
	mux.HandleFunc("POST /find/{place}", s.fetch)
	mux.HandleFunc("POST /process", s.process)

	return mux
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	var location domain.Location
	// default values
	placeID := "belize-city"
	weather, err := repository.FetchWeather(placeID)
	if err != nil {
		log.Println(err)
	}

	views.Home(weather, location, nil).Render(context.Background(), w)
}

func (s *Server) process(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	location := r.FormValue("search")
	locations, err := repository.FetchLocations(location)
	if err != nil {
		log.Println(err)
	}

	partials.Results(locations).Render(context.Background(), w)
}

func (s *Server) fetch(w http.ResponseWriter, r *http.Request) {
	q := r.PathValue("place")

	weather, err := repository.FetchWeather(q)
	if err != nil {
		log.Println(err)
	}

	locations, err := repository.FetchLocations(q)
	if err != nil {
		log.Println(err)
	}

	partials.Weather(weather, locations[0]).Render(context.Background(), w)
}
