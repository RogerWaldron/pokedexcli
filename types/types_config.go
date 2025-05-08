package types

import "net/http"

type ApiConfig struct {
	ApiClient        http.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

type Client struct {
	HttpClient http.Client
}